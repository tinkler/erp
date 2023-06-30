package pssm

import (
	"context"
	"net/http"

	"github.com/tinkler/mqttadmin/pkg/db"
	"github.com/tinkler/mqttadmin/pkg/status"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func WrapSchemaToDB() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			schemaName := w.Header().Get("x-erp-as")
			if schemaName == "" {
				status.HttpError(w, status.StatusBadGateway("x-erp-as illegal"))
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			r = r.WithContext(WithSchemaToDB(schemaName, r.Context()))
			next.ServeHTTP(w, r)
		})
	}
}

func DB(ctx context.Context) *gorm.DB {
	return db.GetDB(ctx)
}

func WithSchemaToDB(schemaName string, ctx context.Context) context.Context {
	dbInst := db.DB().WithContext(ctx)
	dbInst.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   schemaName + ".",
		SingularTable: true,
	}
	return db.WithValue(ctx, dbInst)
}
