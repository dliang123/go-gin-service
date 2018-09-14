package util

import "go-gin-demo/model"

type ResponseUtil struct{}

// 创建成功返回体
func (*ResponseUtil) CreateSuccessResponse(data interface{}) (response *model.Response) {
	response = &model.Response{ResCode: "000000", ResDesc: "success"}
	response.Data = data
	return response
}

func (*ResponseUtil) CreateFailResponse(err interface{}) (response *model.Response) {
	response = &model.Response{ResCode: "999999", ResDesc: "系统异常"}
	// TODO 处理异常
	return response
}