package graph

import (
	"context"
	"errors"
	"fmt"
	"graphql-server/db"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store db.KeyValueStore
}

func NewResolver(store db.KeyValueStore) *Resolver {
	return &Resolver{store}
}

var (
	// JWT用のシークレット
	jwtSecret = []byte("NUr5V86gjUHIGEA8GWEgN2qSpMdRV0s61prkNFVnk")
)

func (r *Resolver) createJWT(userName string) (string, error) {
	now := time.Now()
	expire := now.Add(365 * 24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        userName,
		Audience:  "user",
		Issuer:    "flutter-exam",
		ExpiresAt: expire.Unix(),
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (r *Resolver) validateJWT(ctx context.Context) (userName string, err error) {
	header, ok := ctx.Value("header").(http.Header)
	if !ok {
		return "", errors.New("could not get http header")
	}

	authz := header.Get("Authorization")
	parts := strings.Split(authz, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid `Authorization` header format")
	}

	claims := &jwt.StandardClaims{}

	_, err = jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	if err := claims.Valid(); err != nil {
		return "", err
	}

	return claims.Id, nil

}
