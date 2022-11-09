package searchengine

import (
	"fmt"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"github.com/mztlive/logger"
	"github.com/mztlive/utils/mapx"
	"go.uber.org/zap"
)

// MeiliSearchEngine MeiliSearch引擎
type MeiliSearchEngine struct {
	client *meilisearch.Client
}

// NewMeiliSearchEngine 实例化一个MeiliSearch引擎
func NewMeiliSearchEngine(dns, apiKey string) *MeiliSearchEngine {
	return &MeiliSearchEngine{
		client: meilisearch.NewClient(meilisearch.ClientConfig{
			Host:   dns,
			APIKey: apiKey,
		}),
	}
}

// Put 将多个文档写入搜索引擎
func (se *MeiliSearchEngine) Put(index string, data ...any) error {
	if index == "" {
		return fmt.Errorf("Put Data To SearchEngine Failed: index is empty")
	}

	if len(data) == 0 {
		return fmt.Errorf("Put Data To SearchEngine Failed: data is empty")
	}

	docIndex := se.client.Index(index)

	document := mapx.Structs2Maps(data...)
	task, err := docIndex.AddDocuments(document, "id")
	if err != nil {
		logger.Logger().Error(
			"Put Documents to MeiliSearch Error",
			zap.Error(err),
			zap.Any("documents", document),
			zap.String("index", index),
		)
		return fmt.Errorf("call addDocuments error: %w", err)
	}

	return se.checkTaskStatus(task.TaskUID, index, 0)
}

// 检查任务处理结果
func (se *MeiliSearchEngine) checkTaskStatus(taskUID int64, index string, repeatNum int) error {

	// 超过3次直接返回错误
	if repeatNum == 3 {
		return fmt.Errorf("check task status error")
	}

	docIndex := se.client.Index(index)

	task, err := docIndex.GetTask(taskUID)
	if err != nil {
		logger.Logger().Error(
			"Check Task Status Error",
			zap.Error(err),
			zap.Int64("taskUID", taskUID),
			zap.String("index", index),
		)
		return fmt.Errorf("get task error: %w", err)
	}

	// 重新处理
	if task.Status == meilisearch.TaskStatusProcessing {
		time.Sleep(time.Millisecond * 100)
		return se.checkTaskStatus(taskUID, index, repeatNum+1)
	}

	// 失败了直接返回
	if task.Status == meilisearch.TaskStatusFailed {
		logger.Logger().Error(
			"Check Task Status Error",
			zap.Int64("taskUID", taskUID),
			zap.String("index", index),
			zap.String("error", task.Error.Message),
			zap.String("errorCode", task.Error.Code),
		)

		return fmt.Errorf("task execute failed :%s", task.Error.Message)
	}

	return nil
}
