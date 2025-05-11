package routes

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("ed45w3ecrtdxs2t1nmftvby") // Substitua por uma chave secreta segura

// GenerateJWT gera um token JWT para autenticação
func GenerateJWT(cnpj string) (string, error) {
	claims := jwt.MapClaims{
		"cnpj": cnpj,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
