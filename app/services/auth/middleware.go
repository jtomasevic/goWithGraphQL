package auth

import (
	"context"
	"fmt"
	"net/http"

	app "github.com/evax/app"
	jwt "github.com/evax/app/services/auth/jwt"
	u "github.com/evax/app/services/users/repository"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)

			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				fmt.Println("Invalid token")
				return
			}

			// create user and check if user exists in db
			user := u.User{Username: username}
			id, err := app.Application.UserService.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				fmt.Println("GetUserIdByUsername failed")
				return
			}
			user.ID = id
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *u.User {
	raw, _ := ctx.Value(userCtxKey).(*u.User)
	return raw
}
