package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status         int                 `json:"status"`
	Data           interface{}         `json:"data"`
	Message        string              `json:"message"`
	ContentType    string              `json:"contentType"`
	ResponseWriter http.ResponseWriter `json:"responseWriter"`
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:         http.StatusOK,
		ResponseWriter: w,
		ContentType:    "application/json",
	}
}

func (response *Response) Send() {
	response.ResponseWriter.Header().Set("Content-Type", response.ContentType)
	response.ResponseWriter.WriteHeader(response.Status)

	output, _ := json.Marshal(&response.Data)

	resp := string(output)

	if resp == "null" {
		fmt.Fprint(response.ResponseWriter, response.Message)
	} else {
		fmt.Fprint(response.ResponseWriter, resp)
	}
}

func (response *Response) NotFound() {
	response.Status = http.StatusNotFound
	response.Message = "Not Found"
}

func (response *Response) UnprocessableEntity() {
	response.Status = http.StatusUnprocessableEntity
	response.Message = "Unprocessable Entity"
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

func SendUnprocessable(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}
