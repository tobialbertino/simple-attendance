package tokenize

import (
	"log"
	"simple-attendance/exception"
	"simple-attendance/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenKey string = config.GetKeyConfig("ACCESS_TOKEN_KEY")
var RefreshTokenKey string = config.GetKeyConfig("REFRESH_TOKEN_KEY")

type AccountClaims struct {
	jwt.RegisteredClaims
	ID        string
	ExpiresAt int64 `json:"exp,omitempty"`
}

// Create token
func GenerateAccessToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(AccessTokenKey))
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", exception.Wrap("Tokenize", 500, err)
	}

	return t, nil
}

func GenerateRefreshToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	rt, err := token.SignedString([]byte(RefreshTokenKey))
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", exception.Wrap("Tokenize", 500, err)
	}

	return rt, nil
}

func VerifyRefreshToken(auth string) (*jwt.Token, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(RefreshTokenKey), nil
	}
	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, exception.Wrap("error parsing token", 400, err)
	}
	if !token.Valid {
		return nil, exception.Wrap("invalid token general", 400, err)
	}
	return token, nil
}

func VerifyAccessToken(auth string) (*jwt.Token, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(AccessTokenKey), nil
	}
	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, exception.Wrap("error parsing token", 400, err)
	}
	if !token.Valid {
		return nil, exception.Wrap("invalid token general", 400, err)
	}
	return token, nil
}

// helper get ID from Token AccessToken
func GetIdUserFromToken(token string) (string, error) {
	// validasi dari token signature
	tokenDetail, err := VerifyAccessToken(token)
	if err != nil {
		return "", exception.NewClientError("Refresh token tidak valid", 400)
	}

	// Cast data to map[string]interface{} and cast data["name"] to string
	claims := tokenDetail.Claims.(jwt.MapClaims)
	dataID := claims["ID"].(string)

	return dataID, nil
}
