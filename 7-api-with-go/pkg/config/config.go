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

	// MongoDB Config
	MongoURI string

	// Mysql config
	MYSQL_Host     string
	MYSQL_User     string
	MYSQL_Password string
	MYSQL_DbName   string
	MYSQL_Port     string
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

		// Mongodb Config
		MongoURI := os.Getenv("MONGO_URI")

		// Mysql config
		MYSQL_Host := os.Getenv("MYSQL_HOST")
		MYSQL_User := os.Getenv("MYSQL_USER")
		MYSQL_Password := os.Getenv("MYSQL_PASSWORD")
		MYSQL_DbName := os.Getenv("MYSQL_DB_NAME")
		MYSQL_Port := os.Getenv("MYSQL_PORT")

		Cnf = &Config{
			// App
			Port:   Port,
			AppEnv: AppEnv,

			// Postgres
			PGHost:     PGHost,
			PGUser:     PGUser,
			PGPassword: PGPassword,
			PGDbName:   PGDbName,
			PGPort:     PGPort,

			// MongoDB
			MongoURI: MongoURI,

			// Mysql
			MYSQL_Host:     MYSQL_Host,
			MYSQL_User:     MYSQL_User,
			MYSQL_Password: MYSQL_Password,
			MYSQL_DbName:   MYSQL_DbName,
			MYSQL_Port:     MYSQL_Port,
		}

		isLoaded = true
	})
}

func Get() *Config {
	if !isLoaded {
		log.Print("⚠️ Config.Get() called before Load() — returning false")
	}

	return Cnf
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
