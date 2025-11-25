package services

import (
	"errors"
	"fmt"
	env "riskmanagement/lib/env"
	models "riskmanagement/models/user"

	// modelsLogin "riskmanagement/models/pgsuser"
	// "riskmanagement/lib"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/golang-package-library/logger"
)

// JWTAuthService service relating to authorization
type JWTAuthService struct {
	env    env.Env
	logger logger.Logger
}

// NewJWTAuthService creates a new auth service
func NewJWTAuthService(env env.Env, logger logger.Logger) JWTAuthService {
	return JWTAuthService{
		env:    env,
		logger: logger,
	}
}

// Authorize authorizes the generated token
func (s JWTAuthService) Authorize(tokenString string, pernr string) (bool, error, jwt.MapClaims) {
	claims := jwt.MapClaims{}

	// Parse token dengan claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})

	if err != nil {
		// Cek jenis error
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return false, errors.New("token malformed"), nil
			}
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return false, errors.New("token expired or not valid yet"), nil
			}
		}
		// Fallback error
		return false, fmt.Errorf("token parsing failed: %v", err), nil
	}

	// Debug log
	fmt.Println("token valid =>", token.Valid)
	fmt.Println("token raw =>", token.Raw)
	fmt.Println("token signature =>", string(token.Signature))

	// Ambil claim pernr
	pernrFromToken, ok := claims["pernr"].(string)
	if !ok || pernrFromToken == "" {
		return false, errors.New("pernr claim missing or not string"), nil
	}
	fmt.Println("pernrFromToken =>", pernrFromToken)

	// Validasi token + pernr
	if !token.Valid {
		return false, errors.New("token invalid"), nil
	}

	if pernrFromToken != pernr {
		return false, fmt.Errorf("pernr mismatch: token=%s, expected=%s", pernrFromToken, pernr), nil
	}

	return true, nil, claims
}

// CreateToken creates jwt auth token
func (s JWTAuthService) CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": *user.Email,
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Zap.Error("JWT validation failed: ", err)
	}

	return tokenString
}

// CreateTokenGlobal creates jwt auth token
func (s JWTAuthService) CreateTokenGlobal() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"appID":   "v2",
		"appName": "riskmanagement",
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Zap.Error("JWT validation failed: ", err)
	}

	return tokenString
}

// CreateTokenGlobal creates jwt auth token
func (s JWTAuthService) CreateTokenByPN(pernr string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pernr":   pernr,
		"appID":   "v2",
		"appName": "riskmanagement",
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Zap.Error("JWT validation failed: ", err)
	}

	return tokenString
}

func (s JWTAuthService) CreateRealisasiToken(pernr string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pernr":   pernr,
		"appID":   "v2",
		"appName": "realisasi-kredit",
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Zap.Error("JWT validation failed: ", err)
	}

	return tokenString
}

func (s JWTAuthService) CreateArlordsToken(pernr string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pernr":   pernr,
		"appID":   "v2",
		"appName": "arlords",
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Zap.Error("JWT validation failed: ", err)
	}

	return tokenString
}
