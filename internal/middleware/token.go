package middleware

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TypeAccess  = "access"
	TypeRefresh = "refresh"

	AccessTTL  = 15 * time.Minute
	RefreshTTL = 7 * 24 * time.Hour
)

type Claims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	secret []byte
}

func NewTokenManager(secret string) *TokenManager {
	return &TokenManager{secret: []byte(secret)}
}

func (m *TokenManager) Generate(userID uint, tokenType string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		Type: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(userID), 10),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
}

func (m *TokenManager) AccessToken(userID uint) (string, error) {
	return m.Generate(userID, TypeAccess, AccessTTL)
}

// ponytail: stateless refresh token (no revocation DB). If logout/revoke is
// required, persist a refresh-token table and check it on refresh.
func (m *TokenManager) RefreshToken(userID uint) (string, error) {
	return m.Generate(userID, TypeRefresh, RefreshTTL)
}

func (m *TokenManager) Validate(tokenString, tokenType string) (uint, error) {
	tok, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(*jwt.Token) (interface{}, error) {
		return m.secret, nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := tok.Claims.(*Claims)
	if !ok || !tok.Valid || claims.Type != tokenType {
		return 0, errors.New("invalid token")
	}
	id, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
