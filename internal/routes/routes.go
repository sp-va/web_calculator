package routes

import (
	"calculator/internal/service"
	"calculator/internal/types"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func handleError(w http.ResponseWriter, err error) {
	errorMapping := map[string]struct {
		status int
		msg    string
	}{
		"expression invalid":  {http.StatusUnprocessableEntity, "expression invalid"},
		"invalid operator":    {http.StatusBadRequest, "invalid operator"},
		"division by zero":    {http.StatusUnprocessableEntity, "division by zero"},
		"not enough operands": {http.StatusUnprocessableEntity, "not enough operands"},
	}

	if mappedError, exists := errorMapping[err.Error()]; exists {
		w.WriteHeader(mappedError.status)
		json.NewEncoder(w).Encode(types.Response{Error: mappedError.msg})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(types.Response{Error: "server error"})
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req types.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.Response{Error: "json invalid"})
		return
	}

	expression := strings.TrimSpace(req.Expression)
	result, err := service.Calc(expression)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types.Response{Result: strconv.FormatFloat(result, 'f', -1, 64)})
}
