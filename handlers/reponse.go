package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	Error   = "Error"
	Message = "Message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Cookie string `json:"cookie"`
}

func newResponse(messageType string, message string, data interface{}) response {
	return response{
		MessageType: messageType,
		Message:  message,
		Data: data,
	}
}

func newResponseWithCookie(messageType string, message string, data interface{}, cookie string) response {
	return response{
		messageType,
		message,
		data,
		cookie,
	}
}

func responseJson(w http.ResponseWriter, statusCode int, resp response) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
