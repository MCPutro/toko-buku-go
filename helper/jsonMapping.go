package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func WriteToResponseBody2(w http.ResponseWriter, errParam error, data interface{}) {
	var webResponse Response

	if errParam != nil {
		webResponse = Response{
			Status:  "error",
			Message: errParam.Error(),
			Data:    nil,
		}
	} else {
		webResponse = Response{
			Status: "success",
			Data:   data,
		}
	}

	WriteToResponseBody(w, webResponse)
}
