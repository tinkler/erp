package test

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Join(filepath.Dir(currentFile), "../")
}

func LoadEnv() {
	tr, _ := os.Getwd()
	rel, _ := filepath.Rel(tr, basepath)
	os.Chdir(rel)
	godotenv.Load(".env")
}

func Context() context.Context {
	return context.Background()
}

func GetBasePath() string {
	return basepath
}
