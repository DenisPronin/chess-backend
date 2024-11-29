package json_errors

import (
	"encoding/json"
	"net/http"
)

func CatchError(writer http.ResponseWriter, code int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	response := ErrorResponse{
		Code:    code,
		Message: message,
	}

	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "Failed to write response", http.StatusInternalServerError)
	}
}
