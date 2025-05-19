package main

import (
	"encoding/json"
	"net/http"
)

type responsejson struct {
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json: "message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := responsejson{
		Error:   false,
		Message: "Hi Request recieved",
	}
	output, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(output)

}
