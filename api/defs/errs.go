package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error Err
}

var (
	EeeorRequestBodyParseFailed = ErrorResponse{HttpSc:400, Error:Err{Error: "Request body is not correct", ErrorCode:"001"}}
	EeeorNotAuthUser = ErrorResponse{HttpSc:401, Error:Err{Error: "User authentication failed", ErrorCode:"002"}}
)
