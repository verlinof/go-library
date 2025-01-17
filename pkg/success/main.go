package pkg_success

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

func SuccessCreateData(data interface{}) *ClientSuccess {
	return &ClientSuccess{
		Message: "Success",
		Data:    data,
	}
}
