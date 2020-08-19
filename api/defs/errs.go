package defs

type Err struct {
	Error     string `json:"rror"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{httpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrResponse{httpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
)
