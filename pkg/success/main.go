package pkg_success

import "fmt"

type ClientSuccess struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessGetData(data any) *ClientSuccess {
	return &ClientSuccess{
		Message: "Success",
		Data:    data,
	}
}

func SuccessDeleteData(id any) *ClientSuccess {
	return &ClientSuccess{
		Message: fmt.Sprintf("Success delete data with ID %v", id),
		Data:    id,
	}
}

func SuccessCreateData(data interface{}) *ClientSuccess {
	return &ClientSuccess{
		Message: "Success",
		Data:    data,
	}
}
