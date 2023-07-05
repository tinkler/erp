package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/tinkler/mqttadmin/pkg/logger"
)

var (
	down  = flag.Bool("down", false, "migrate up")
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

	schema, err := copyAndReplaceFiles("./db")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", os.Getenv("MIGRATE_DSN"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + " AUTHORIZATION erp;")
	if err != nil {
		panic(err)
	}

	m, err := migrate.New(
		"file://./db/"+schema,
		os.Getenv("MIGRATE_DSN")+"&search_path="+schema,
	)
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

func copyAndReplaceFiles(path string) (schema string, err error) {
	tmpFolder, err := ioutil.TempDir(path, "pssm")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary folder: %v", err)
	}
	schema = filepath.Base(tmpFolder)

	files, err := filepath.Glob(filepath.Join(path, "*.sql"))
	if err != nil {
		return schema, fmt.Errorf("failed to get SQL files: %v", err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return schema, fmt.Errorf("failed to read file %s: %v", file, err)
		}

		newContent := strings.ReplaceAll(string(content), "passmexample", schema)

		destFile := filepath.Join(tmpFolder, filepath.Base(file))

		err = ioutil.WriteFile(destFile, []byte(newContent), 0644)
		if err != nil {
			return schema, fmt.Errorf("failed to write file %s: %v", destFile, err)
		}
	}

	fmt.Println("Files copied and replaced successfully!")

	return schema, nil
}
