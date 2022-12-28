package jwt

import (
	"context"
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

var ErrNotLogin = errors.New("not logged in")

type authKey struct{}

func WithContext(ctx context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(ctx, authKey{}, claims)
}

func FromContext(ctx context.Context) (jwt.Claims, bool) {
	rs, ok := ctx.Value(authKey{}).(jwt.Claims)
	return rs, ok
}

func GetUserID(ctx context.Context) (int, error) {
	claim, ok := FromContext(ctx)
	if !ok {
		return 0, ErrNotLogin
	}
	mapClaim, ok := claim.(jwt.MapClaims)
	if !ok {
		return 0, ErrNotLogin
	}
	id, ok := mapClaim["jti"]
	if !ok {
		return 0, ErrNotLogin
	}

	return strconv.Atoi(id.(string))
}
