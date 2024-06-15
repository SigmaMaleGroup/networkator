package tokens

import (
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateJWT returns a JWT token
func GenerateJWT(userID int64, email, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix()
	claims["userID"] = userID
	claims["email"] = email
	claims["role"] = role
	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		slog.Error("error signing token", err)
		return "Signing Error", err
	}

	return tokenString, nil
}
