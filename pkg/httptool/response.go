package httptool

// HttpResponseItemsTotal 响应对象总数
// HttpResponseItemsTotal response total count of object
type HttpResponseItemsTotal struct {
	// 响应对象总数
	// response total count of object
	TotalCount int64 `json:"totalCount" yaml:"totalCount"`
}

// HttpResponseItemsID 响应对象 Id
// HttpResponseItemsID response id of object
type HttpResponseItemsID struct {
	// 响应对象 Id
	// response id of object
	ID int64 `json:"id" yaml:"id"`
}

// HttpQueryPaginated 分页查询
// HttpQueryPaginated paginated query
type HttpQueryPaginated struct {
	// 当前页数量
	// current page index
	PageIndex int64 `json:"pageIndex" yaml:"pageIndex"`

	// 每页显示最大条目数
	// max item count per page
	PageSize int64 `json:"pageSize" yaml:"pageSize"`

	// 是否倒序
	// is desc
	Desc bool `json:"desc" yaml:"desc"`
}

// BaseHttpResponse 基础响应对象
// BaseHttpResponse base response object
type BaseHttpResponse struct {
	// 响应代码
	// response code
	Code int64 `json:"errorCode" yaml:"errorCode"`

	// 错误信息
	// error message
	ErrorMessage string `json:"errorMessage,omitempty" yaml:"errorMessage,omitempty"`

	// 错误详细信息
	// error detail message
	ErrorDetail interface{} `json:"errorDetail,omitempty" yaml:"errorDetail,omitempty"`

	// 响应数据
	// response data
	Data interface{} `json:"data,omitempty" yaml:"data,omitempty"`
}

// HttpResponsePaginated 分页响应对象
// HttpResponsePaginated paginated response object
type HttpResponsePaginated struct {
	HttpResponseItemsTotal
	HttpQueryPaginated
	BaseHttpResponse
}
