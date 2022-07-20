package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FeatchListResponse struct {
	Size      int `json:"size"`
	PageIndex int `json:"page_index"`
	Total     int `json:"total"`
	// Data      interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	errorResponse := ErrorResponse{}
	errorResponse.StatusCode = statusCode
	errorResponse.ErrorMessage = err.Error()
	if err != nil {
		JSON(w, statusCode, errorResponse)
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
