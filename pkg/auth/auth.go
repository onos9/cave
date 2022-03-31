package auth

import (
	"time"

	cfg "github.com/cave/config"
	"github.com/cave/pkg/models"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// RefreshClaims represents refresh token JWT claims
type RefreshClaims struct {
	RefreshTokenID string `json:"refreshTokenID"`
	ExternalID     string `json:"userID"`
	Role           string `json:"role"`
	jwt.StandardClaims
}

// AccessClaims represents access token JWT claims
type AccessClaims struct {
	AccessTokenID string `json:"accessTokenID"`
	ExternalID    string `json:"userID"`
	Role          string `json:"role"`
	jwt.StandardClaims
}

// IssueAccessToken generate access tokens used for auth
func IssueAccessToken(u models.User) (string, error) {
	expireTime := time.Now().Add(time.Hour) // 1 hour

	// Generate encoded token
	claims := AccessClaims{
		uuid.New().String(),
		u.ID.String(),
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(cfg.GetConfig().JWTAccessSecret))
}

// IssueRefreshToken generate refresh tokens used for auth
func IssueRefreshToken(u models.User) (string, error) {
	expireTime := time.Now().Add((24 * time.Hour) * 14) // 14 days

	// Generate encoded token
	claims := RefreshClaims{
		uuid.New().String(),
		u.ID.String(),
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(cfg.GetConfig().JWTRefreshSecret))
}
