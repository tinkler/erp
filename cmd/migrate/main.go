package main

import (
	"flag"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/tinkler/mqttadmin/pkg/logger"
)

var (
	down  = flag.Bool("down", false, "migrate up")
	p     = flag.String("p", "erpv1", "migrate project name")
	dev   = flag.Bool("dev", true, "migrate dev")
	force = flag.Int("force", -1, "migrate force version")
)

// Migrate database
func main() {
	logger.ConsoleLevel = logger.LL_DEBUG
	flag.Parse()
	if *dev {
		os.Chdir("../../")
	}
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file:" + err.Error())
		return
	}
	var m *migrate.Migrate
	switch *p {
	case "erpv1":
		m, err = migrate.New(
			"file://./db/erpv1",
			os.Getenv("MIGRATE_DSN"),
		)
	default:
		logger.Error("Unknown project name")
		return
	}
	if err != nil {
		logger.Error(err)
		return
	}
	if *force > -1 {
		if err := m.Force(*force); err != nil {
			logger.Error(err)
		}
		return
	}
	if *down {
		if err := m.Down(); err != nil {
			logger.Error(err)
		}
	} else {
		if err := m.Up(); err != nil {
			logger.Error(err)
		}
	}
}
