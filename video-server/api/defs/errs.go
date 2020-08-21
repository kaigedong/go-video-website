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
	ErrorDBError                = ErrorResponse{HTTPSc: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrorResponse{HTTPSc: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
