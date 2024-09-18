package utils

import "context"

var userCtxKey = &contextKey{"user"}

type User struct {
	ID       string
	Username string
	Role     string
}

type contextKey struct {
	name string
}

func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}
