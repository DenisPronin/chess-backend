package user_handlers

import (
	"chess/internal/json_errors"
	"context"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) Profile(writer http.ResponseWriter, request *http.Request) {
	userID, ok := request.Context().Value("userID").(float64)
	if !ok {
		json_errors.CatchError(writer, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.UserRepo.GetUserByID(context.Background(), int(userID))
	if err != nil {
		json_errors.CatchError(writer, http.StatusNotFound, "User not found")
		return
	}

	// Формируем JSON-ответ
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(user)

	if err != nil {
		json_errors.CatchError(writer, http.StatusInternalServerError, "Error writing response")
	}
}
