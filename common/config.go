/**
* @Author: ZhangZe
* @Date: 2022/4/9 18:47
 */

package common

// 定义数据的返回结构体

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) Result(code int, message string, data interface{}) Response {
	r.Code = code
	r.Message = message
	r.Data = data
	return r
}
