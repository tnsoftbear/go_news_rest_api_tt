package jwt

import (
	"errors"
	"fiber-reform-rest/internal/infra/config"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenPayload defines the payload for the token
type TokenPayload struct {
	ID uint
}

// JWTManager provides JWT token generation and verification
type JWTManager struct {
	cfg *config.Jwt
}

func NewJWTManager(cfg *config.Jwt) *JWTManager {
	return &JWTManager{
		cfg: cfg,
	}
}

// Generate generates the jwt token based on payload
func (j *JWTManager) Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration(j.cfg.Expiration)
	if err != nil {
		log.Fatalf("Invalid time duration: %v", err)
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})

	token, err := t.SignedString([]byte(j.cfg.Tokenkey))
	if err != nil {
		log.Fatal(err)
	}

	return token
}

func (j *JWTManager) parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(j.cfg.Tokenkey), nil
	})
}

// Verify verifies the jwt token against the secret
func (j *JWTManager) Verify(token string) (*TokenPayload, error) {
	parsed, err := j.parse(token)

	if err != nil {
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &TokenPayload{
		ID: uint(id),
	}, nil
}
