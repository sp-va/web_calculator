package middleware

import (
	"bytes"
	"calculator/internal/types"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
)

func isExpressionCorrect(expression string) bool {
	validExpression := regexp.MustCompile(`^[0-9+\-*/().\s]+$`)
	return validExpression.MatchString(expression)
}

func ValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(types.ErrorResponse{Error: "method not allowed"})
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(types.ErrorResponse{Error: "cant read request body"})
			return
		}

		log.Printf("Middleware received body: %s", string(body))

		var req types.Request
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(types.ErrorResponse{Error: "wrong request formatting"})
			return
		}

		if req.Expression == "" || !isExpressionCorrect(req.Expression) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(types.ErrorResponse{Error: "expression wrong"})
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next(w, r)
	}
}
