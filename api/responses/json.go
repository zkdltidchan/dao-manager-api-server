package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FeatchListResponse struct {
	Size      int         `json:"size"`
	PageIndex int         `json:"page_index"`
	Total     int         `json:"total"`
	Data      interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			// ErrorCode int    `json:"error_code"`
			Error string `json:"error"`
		}{
			// ErrorCode: statusCode,
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
