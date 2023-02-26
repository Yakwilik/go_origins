package handler

import (
	"context"
	"net/http"
	"strings"
)

const (
	userCTX = "user"
)

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(AuthorizationHeader)
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			JSONResponse(w, http.StatusUnauthorized, interfaceMap{"message": "You are not authorized"})
			return
		}
		token := headerParts[1]
		user, err := h.services.ParseToken(token)
		if err != nil {
			JSONResponse(w, http.StatusUnauthorized, interfaceMap{"message": "Bad access token"})
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, userCTX, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	},
	)
}
