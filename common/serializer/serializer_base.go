package serializer

type ResponseBase struct {
	Code    int    `json:"status"`
	Message string `json:"detail"`
}

type ResponseData struct {
	ResponseBase
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string) ResponseBase {
	return ResponseBase{
		Code:    code,
		Message: msg,
	}
}

func NewResponseData(code int, msg string, data interface{}) ResponseData {
	return ResponseData{
		ResponseBase: NewResponse(code, msg),
		Data:         data,
	}
}