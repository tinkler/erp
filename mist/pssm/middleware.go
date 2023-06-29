package pssm

import (
	"context"
	"net/http"

	"github.com/tinkler/mqttadmin/pkg/db"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	dbKey keyString = "db"
)

type keyString string

func WrapSchemaToDB() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			schemaName := w.Header().Get("x-erp-as")
			if schemaName == "" {
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			r = r.WithContext(WithSchemaToDB(schemaName, r.Context()))
			next.ServeHTTP(w, r)
		})
	}
}

func DB(ctx context.Context) *gorm.DB {
	return ctx.Value(dbKey).(*gorm.DB)
}

func WithSchemaToDB(schemaName string, ctx context.Context) context.Context {
	db := db.DB().WithContext(ctx)
	db.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   schemaName + ".",
		SingularTable: true,
	}
	return context.WithValue(ctx, dbKey, db)
}
