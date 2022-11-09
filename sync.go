package searchengine

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/mztlive/foundation/vo"
	"github.com/mztlive/repository/database/update"
)

// 将任意对象转为interface{}
func struct2Interface[T any](values []T) []interface{} {
	result := make([]interface{}, len(values))
	for i, value := range values {
		result[i] = interface{}(value)
	}
	return result
}

// Config 定义了同步数据到搜索引擎需要的内容
type Config[T any] struct {

	// Document 在搜索引擎的索引名称
	DocumentName string

	// FindItemsFn 获取未同步数据的方法
	FindItemsFn func(ctx context.Context, paginator vo.Paginator) ([]*T, error)

	// ItemTableName 数据的表名
	ItemTableName string

	// ItemIdentityDBColumnName 数据的唯一编号在DB里面的列名, 用于更新SyncStatus状态
	ItemIdentityDBColumnName string

	// Item结构体中Identity数据的属性名
	ItemIdentityAttributeName string
}

func Sync[T any](ctx context.Context, cfg Config[T], se Engine) error {
	var (
		items []*T
		err   error
	)

	if items, err = cfg.FindItemsFn(ctx, vo.Paginator{Page: 1, PageSize: 20}); err != nil {
		return fmt.Errorf("call findItem fn failed: %w", err)
	}

	// 如果没有数据，则直接返回
	if len(items) == 0 {
		return nil
	}

	if err = se.Put(cfg.DocumentName, struct2Interface(items)...); err != nil {
		return fmt.Errorf("put to SearchEngine Failed: %w", err)
	}

	identities := make([]string, len(items))

	// 因为T里面表示identity字段的名字可能都不一样，要通过反射获取值
	for i, item := range items {
		identity := ""
		v := reflect.ValueOf(item).Elem()
		field := v.FieldByName(cfg.ItemIdentityAttributeName)

		if cfg.ItemIdentityAttributeName == "ID" || cfg.ItemIdentityAttributeName == "Id" {
			identity = strconv.Itoa(int(field.Int()))
		} else {
			identity = field.String()
		}

		if identity == "<invalid Value>" {
			return errors.New("get Item Identity Error")
		}

		identities[i] = identity
	}

	// 更新为已同步
	err = SetSynced(ctx, cfg.ItemTableName, identities, cfg.ItemIdentityDBColumnName)

	if err != nil {
		return fmt.Errorf("update sync status failed: %w", err)
	}
	return nil
}

// SetNoSync 将数据设置为未同步
// 这个方法会操作数据库
func SetNoSync(ctx context.Context, tableName string, identities []string, identityKeyName string) error {
	return update.UpdateFieldByIdentities(ctx, update.UpdateSingleFieldPayload[int]{
		Table:           tableName,
		FieldName:       SyncStatusFieldName,
		FieldValue:      SyncStatusIsFalse,
		IdentityColName: identityKeyName,
		Identities:      identities,
	})
}

// SetNoSync 将数据设置为未同步
// 这个方法会操作数据库
func SetSynced(ctx context.Context, tableName string, identities []string, identityKeyName string) error {
	return update.UpdateFieldByIdentities(ctx, update.UpdateSingleFieldPayload[int]{
		Table:           tableName,
		FieldName:       SyncStatusFieldName,
		FieldValue:      SyncStatusIsTrue,
		IdentityColName: identityKeyName,
		Identities:      identities,
	})
}
