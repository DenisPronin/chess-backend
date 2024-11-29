package user_handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) Profile(writer http.ResponseWriter, request *http.Request) {
	userID, ok := request.Context().Value("userID").(float64)
	if !ok {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.UserRepo.GetUserByID(context.Background(), int(userID))
	if err != nil {
		http.Error(writer, "User not found", http.StatusNotFound)
		return
	}

	// Формируем JSON-ответ
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(user)

	if err != nil {
		http.Error(writer, "Failed to write response", http.StatusInternalServerError)
	}

}
