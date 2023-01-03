package server

import "encoding/json"

type Content struct {
	Message string `json:"message"`
}

type Response struct {
	Status int32   `json:"status"`
	Cont   Content `json:"content"`
}

func MakeResponse(status int32, message string) string {
	response := Response{
		Status: status,
		Cont:   Content{Message: message},
	}
	resByte, _ := json.Marshal(response)

	return string(resByte)
}
