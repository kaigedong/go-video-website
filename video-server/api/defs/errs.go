package defs

type Err struct {
	Error     string `json:"rror"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HTTPSc int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HTTPSc: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{HTTPSc: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
)
