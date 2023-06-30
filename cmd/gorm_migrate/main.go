package main

import (
	"github.com/tinkler/erp/pkg/model/accounting_system"
	"github.com/tinkler/erp/pkg/model/company"
	"github.com/tinkler/erp/pkg/model/user"
	"github.com/tinkler/erp/test"
	"github.com/tinkler/mqttadmin/pkg/db"
)

func main() {
	test.LoadEnv()
	db := db.DB()
	db.AutoMigrate(
		&user.User{},
		&company.Company{},
		&accounting_system.AccountingSystem{},
	)

}
