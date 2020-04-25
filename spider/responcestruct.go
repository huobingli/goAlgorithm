package spider

import "net/http"

//数据响应的类型
type Response struct {
	// HTTP响应
	httpResp *http.Response
	//响应的深度
	depth uint32
}

//用于创建一个新的响应实例
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

//用于获取HTTP响应
func (resp *Response) HTTPResp() *http.Response {
	return resp.httpResp
}

//用于获取响应深度
func (resp *Response) Depth() uint32 {
	return resp.depth
}
