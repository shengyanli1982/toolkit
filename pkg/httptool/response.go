package httptool

// HttpResponseItemsTotal 是响应对象的总数
// HttpResponseItemsTotal is the total count of response objects
type HttpResponseItemsTotal struct {
	// TotalCount 是响应对象的总数
	// TotalCount is the total count of response objects
	TotalCount int64 `json:"totalCount" yaml:"totalCount"`
}

// HttpResponseItemsID 是响应对象的 ID
// HttpResponseItemsID is the ID of the response object
type HttpResponseItemsID struct {
	// ID 是响应对象的 ID
	// ID is the ID of the response object
	ID int64 `json:"id" yaml:"id"`
}

// HttpQueryPaginated 是分页查询
// HttpQueryPaginated is the paginated query
type HttpQueryPaginated struct {
	// PageIndex 是当前页的索引
	// PageIndex is the index of the current page
	PageIndex int64 `json:"pageIndex" yaml:"pageIndex"`

	// PageSize 是每页显示的最大条目数
	// PageSize is the maximum number of items displayed per page
	PageSize int64 `json:"pageSize" yaml:"pageSize"`

	// Desc 表示是否倒序
	// Desc indicates whether it is in descending order
	Desc bool `json:"desc" yaml:"desc"`
}

// BaseHttpResponse 是基础响应对象
// BaseHttpResponse is the base response object
type BaseHttpResponse struct {
	// Code 是响应代码
	// Code is the response code
	Code int64 `json:"errorCode" yaml:"errorCode"`

	// ErrorMessage 是错误信息
	// ErrorMessage is the error message
	ErrorMessage string `json:"errorMessage,omitempty" yaml:"errorMessage,omitempty"`

	// ErrorDetail 是错误的详细信息
	// ErrorDetail is the detailed error message
	ErrorDetail interface{} `json:"errorDetail,omitempty" yaml:"errorDetail,omitempty"`

	// Data 是响应数据
	// Data is the response data
	Data interface{} `json:"data,omitempty" yaml:"data,omitempty"`
}

// HttpResponsePaginated 是一个包含分页信息和基础响应的结构体
// HttpResponsePaginated is a struct that includes pagination information and a base response
type HttpResponsePaginated struct {
	// HttpResponseItemsTotal 包含响应对象的总数
	// HttpResponseItemsTotal includes the total count of response objects
	HttpResponseItemsTotal

	// HttpQueryPaginated 包含分页查询信息
	// HttpQueryPaginated includes paginated query information
	HttpQueryPaginated

	// BaseHttpResponse 包含基础响应信息
	// BaseHttpResponse includes basic response information
	BaseHttpResponse
}
