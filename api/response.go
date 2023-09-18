package api

import "encoding/json"

type Response struct {
	Code uint32 `json:"Code"`
	Msg  string `json:"Msg"`
	Data any    `json:"Data"`
}

func ResponseMsg(code uint32, msg string, data any) (string, error) {
	response := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(jsonResponse), nil
}

func ResponseSuccess(msg string) string {
	response := Response{
		Code: RSP_SUCCECC,
		Msg:  msg,
		Data: nil,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return ""
	}

	return string(jsonResponse)
}

func ResponseError(msg string) string {
	response := Response{
		Code: RSP_ERROR,
		Msg:  msg,
		Data: nil,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return ""
	}

	return string(jsonResponse)
}
