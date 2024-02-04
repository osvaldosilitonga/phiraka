package web

type SuccessResp struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorResp struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message any    `json:"msg"`
}
