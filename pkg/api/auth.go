package api

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func secret() []byte {
	pass := os.Getenv("TODO_PASSWORD")
	h := sha256.Sum256([]byte(pass))
	return h[:]
}

func passHash() string {
	pass := os.Getenv("TODO_PASSWORD")
	return fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("TODO_PASSWORD") == "" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Authentification required", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неожиданный метод подписи")
			}
			return secret(), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Authentification required", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["hash"] != passHash() {
			http.Error(w, "Authentification required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func signInHandle(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error())
		return
	}

	if req.Password != os.Getenv("TODO_PASSWORD") {
		writeError(w, "Неверный пароль")
		return
	}

	claims := jwt.MapClaims{
		"hash": passHash(),
		"exp":  time.Now().Add(8 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secret())
	if err != nil {
		writeError(w, err.Error())
		return
	}

	writeJson(w, map[string]string{"token": signed})
}
