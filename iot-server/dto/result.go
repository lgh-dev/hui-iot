package dto

type ResultDTO struct {
	Code int         `json:"code"` //Succ:1  or err code :<1
	Data interface{} `json:"data"` //Result data
	Msg  string      `json:"msg"`  //Succ or Err msg
}

func BuildSucc(result *ResultDTO) *ResultDTO {
	result.Code = 1
	if result.Msg == "" {
		result.Msg = "succ"
	}
	return result
}

func BuildError(result *ResultDTO) *ResultDTO {
	result.Code = 0
	if result.Msg == "" {
		result.Msg = "error"
	}
	return result
}
