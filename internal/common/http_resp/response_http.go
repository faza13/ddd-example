package http_resp

type Response struct {
	ErrorCode int                    `json:"error_code"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data"`
}

func SuccRequest(data map[string]interface{}) Response {
	return Response{Data: data, Message: "OK"}
}

func ErrRequest(errorCode int, err error) Response {
	return Response{ErrorCode: errorCode, Message: err.Error()}
}
