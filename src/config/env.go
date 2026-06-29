package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Env interface {
	LoadEnv()
	GetEnv(key string) string
}

func LoadEnv() {
	envPath, err := findEnvFile()

	if err != nil {
		log.Fatal("ENV gagal")
	}

	err = godotenv.Load(envPath)

	if err != nil {
		log.Fatal("ENV gagal")
	}

	log.Println("Succes")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func findEnvFile() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	for {
		envPath := filepath.Join(dir, ".env")

		if _, err := os.Stat(envPath); err == nil {
			return envPath, nil
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			return "", os.ErrNotExist
		}

		dir = parent
	}
}
