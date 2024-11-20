package health

import (
	"log"
	"net/http"
)

func HandlerHealth(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("OK"))
	if err != nil {
		log.Fatalf("Error with health check")
	}
}
