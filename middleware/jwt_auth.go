package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtMiddleware struct {
	Next http.Handler
}

func NewJwtMiddleware(Next http.Handler) *JwtMiddleware {
	return &JwtMiddleware{
		Next: Next,
	}
}

func (m *JwtMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	secret := []byte(os.Getenv("JWT_ACCESS"))
	tokenHeader := r.Header.Get("Authorization")
	splitToken := strings.TrimPrefix(tokenHeader, "Bearer ")
	token, err := jwt.Parse(splitToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			w.WriteHeader(401)
			return nil, errors.New("method invalid")
		}

		return secret, nil
	})
	if err != nil {
		w.WriteHeader(401)
		log.Print(err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		w.WriteHeader(401)
		log.Print(err)
		return
	}

	exp, ok := claims["exp"].(float64)
	if !ok || int64(exp) < time.Now().Unix() {
		w.WriteHeader(401)
		log.Print(err)
		return
	}
	m.Next.ServeHTTP(w, r)
}
