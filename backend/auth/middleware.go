package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/ent/user"
	"github.com/Faroukhamadi/likude/helpers"
	"github.com/Faroukhamadi/likude/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate   token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			userInstance := helpers.User{Username: username}
			ctx := context.Background()
			user, err := client.User.Query().Where(user.Username(username)).Only(ctx)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			userInstance.ID = strconv.Itoa(user.ID)
			// put it in context
			ctx = context.WithValue(r.Context(), userCtxKey, &userInstance)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *helpers.User {
	raw, _ := ctx.Value(userCtxKey).(*helpers.User)
	return raw
}
