package token

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTToken interface {
	// CreateToken create a token with a content
	CreateToken(content []byte, expired time.Duration) (string, error)

	// VerifyToken verify and return the content
	VerifyToken(tokenString string) ([]byte, error)
}

const fieldContent = "content"

type jwtToken struct {
	secretKey string
}

func NewJWTToken(secretKey string) JWTToken {
	return &jwtToken{secretKey: secretKey}
}

func (j jwtToken) CreateToken(content []byte, expired time.Duration) (string, error) {

	contentBase64 := base64.StdEncoding.EncodeToString(content)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(expired).Unix(),
		fieldContent: contentBase64,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (j jwtToken) VerifyToken(tokenString string) ([]byte, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claims is can not asserted")
	}

	decodeStringInBytes, err := base64.StdEncoding.DecodeString(claims[fieldContent].(string))
	if err != nil {
		return nil, err
	}

	return decodeStringInBytes, nil
}
