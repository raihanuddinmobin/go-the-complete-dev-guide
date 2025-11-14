package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"mobin.dev/pkg/constants"
)

type Config struct {
	// App
	Port   string
	AppEnv string

	// Postgres config
	PGHost     string
	PGUser     string
	PGPassword string
	PGDbName   string
	PGPort     string
}

var (
	Cnf      *Config
	isLoaded bool
	once     sync.Once
)

func Load() {
	once.Do(func() {
		err := godotenv.Load()
		isLoaded = false

		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// App
		Port := os.Getenv("PORT")
		AppEnv := os.Getenv("APP_ENV")

		// Postgres config
		PGHost := os.Getenv("PGHOST")
		PGUser := os.Getenv("PGUSER")
		PGPassword := os.Getenv("PGPASSWORD")
		PGDbName := os.Getenv("PGDBNAME")
		PGPort := os.Getenv("PGPORT")

		Cnf = &Config{
			Port:       Port,
			PGHost:     PGHost,
			PGUser:     PGUser,
			PGPassword: PGPassword,
			PGDbName:   PGDbName,
			PGPort:     PGPort,
			AppEnv:     AppEnv,
		}

		isLoaded = true
	})
}

func (c *Config) IsDevMode() bool {
	if !isLoaded {
		log.Print("⚠️ Config.IsDevMode() called before Load() — returning false")
		return false
	}
	return c.AppEnv == "" || c.AppEnv == constants.DEVELOPMENT
}

func (c *Config) IsStageMode() bool {
	if !isLoaded {
		log.Print("⚠️ Config.IsStageMode() called before Load() — returning false")
		return false
	}
	return c.AppEnv == constants.STAGING
}

func (c *Config) IsProdMode() bool {
	if !isLoaded {
		log.Print("⚠️ Config.IsProdMode() called before Load() — returning false")
		return false
	}
	return c.AppEnv == constants.PRODUCTION
}
