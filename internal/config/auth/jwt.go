package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var TokenTime = map[string]time.Duration{
	"active_token":  time.Minute * 60,
	"refresh_token": time.Hour * 24 * 7,
}

type JWTAuth interface {
	GeneratedToken(id int32, tokenTime string) (string, error)
	ValidateActiveToken(token string) (*jwt.Token, error)
	ValidateRefreshToken(token string) (*jwt.Token, error)
}

type jwtStruct struct {
	secret string
	aud    string
	iss    string
}

func NewJWT(secret, aud, iss string) JWTAuth {
	return &jwtStruct{
		secret: secret,
		aud:    aud,
		iss:    iss,
	}
}

func (j *jwtStruct) GeneratedToken(id int32, tokenTime string) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(TokenTime[tokenTime]).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": j.iss,
		"aud": j.aud,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", fmt.Errorf("failed to generated token :%w", err)
	}

	return tokenString, nil
}

func (j *jwtStruct) ValidateActiveToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}

func (j *jwtStruct) ValidateRefreshToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	},
		jwt.WithoutClaimsValidation(),
	)
}
