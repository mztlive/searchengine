package searchengine

import "github.com/morikuni/failure"

const (
	ReqBodyJsonMarshalError    failure.StringCode = "Request Body Json Marshal Error"
	RespBodyJsonUnmarshalError failure.StringCode = "Response Body Json Unmarshal Error"

	SendHttpError       failure.StringCode = "Send Http Error"
	HttpStatusNoSuccess failure.StringCode = "Http Status Not Success"
)
