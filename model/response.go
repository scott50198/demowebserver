package model

type Response struct {
	StatusCode int         `json:"statusCode"`
	Msg        string      `json:"msg"`
	Contents   interface{} `json:"contents"`
}
