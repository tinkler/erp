package erp

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/tinkler/erp/pkg/model/user"
	"github.com/tinkler/mqttadmin/pkg/status"
)

type contextKey string

const (
	loginUserKey contextKey = "loginUser"
)

type loginUser struct {
	user user.User
}

func (lu loginUser) ID() string {
	return lu.user.ID
}

func WrapAuth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := w.Header()
			id := header.Get("x-erp-uesr-id")
			roles := strings.Split(header.Get("x-erp-roles"), ",")
			if _, err := uuid.Parse(id); err != nil {
				status.HttpError(w, status.StatusBadGateway("x-erp-user-id illegal"))
				return
			}
			if len(roles) == 0 {
				status.HttpError(w, status.StatusBadGateway("missing x-erp-roles"))
				return
			}
			lu := loginUser{}
			lu.user.ID = header.Get("x-erp-uesr-id")
			lu.user.Roles = roles
			next.ServeHTTP(w, r.WithContext(WithLoginUser(r.Context(), lu)))
		})
	}
}

func WithLoginUser(ctx context.Context, lu loginUser) context.Context {
	return context.WithValue(ctx, loginUserKey, &lu)
}

func GetLoginUser(ctx context.Context) *loginUser {
	return ctx.Value(loginUserKey).(*loginUser)
}
