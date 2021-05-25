package std

type (
	Code string

	ResponseBase struct {
		HttpCode int    `json:"status"`
		Message  string `json:"detail"`
	}

	ResponseData struct {
		HttpCode int         `json:"status"`
		Message  string      `json:"detail"`
		Data     interface{} `json:"data"`
	}
)
