package searchengine

// Engine 搜索引擎的接口
type Engine interface {

	// 将多个文档写入搜索引擎
	//
	// params:
	//	index = "partner" //相当于表名
	Put(index string, data ...any) error
}
