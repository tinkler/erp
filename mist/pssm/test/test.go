package test

import (
	"context"

	"github.com/tinkler/erp/mist/pssm"
	"github.com/tinkler/erp/test"
)

func LoadEnv() {
	test.LoadEnv()
}

func Context() context.Context {
	ctx := test.Context()
	return pssm.WithSchemaToDB("passmexample", ctx)
}
