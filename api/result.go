package api

import "encoding/json"

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RSuccess(data interface{}) Result {
	return Result{
		Code:    1,
		Message: "success",
		Data:    data,
	}
}

func RError(data interface{}) Result {
	return Result{
		Code:    0,
		Message: "error",
		Data:    data,
	}
}

func RErrorByte(data interface{}) []byte {
	b, _ := json.Marshal(RError(data))
	return b
}
