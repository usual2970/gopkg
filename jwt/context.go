package jwt

import (
	"context"
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

var ErrNotLogin = errors.New("not logged in")

type authKey struct{}

func WithContext(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, authKey{}, token)
}

func FromContext(ctx context.Context) (*jwt.Token, bool) {
	rs, ok := ctx.Value(authKey{}).(*jwt.Token)
	return rs, ok
}

func GetUserID(ctx context.Context) (int, error) {
	token, ok := FromContext(ctx)
	if !ok {
		return 0, ErrNotLogin
	}

	mapClaim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrNotLogin
	}

	id, ok := mapClaim["jti"]
	if !ok {
		return 0, ErrNotLogin
	}

	return strconv.Atoi(id.(string))
}

func GetAccessToken(ctx context.Context) (string, error) {
	token, ok := FromContext(ctx)
	if !ok {
		return "", ErrNotLogin
	}

	return token.Raw, nil
}
