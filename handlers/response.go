package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// writeResponse is a helper method that allows to write and HTTP status & response
func writeResponse(w http.ResponseWriter, status int, data interface{}, err error) {
	resp := Response{
		Data: data,
	}
	if err != nil {
		resp.Error = fmt.Sprint(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != http.StatusOK {
		w.WriteHeader(status)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Fprintf(w, "error encoding resp %v:%s", resp, err)
	}
}
