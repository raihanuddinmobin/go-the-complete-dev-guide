package app

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type App struct {
	pg          *sql.DB
	redisClient *redis.Client
}

func NewApp(pg *sql.DB, redisClient *redis.Client) *App {
	return &App{
		pg,
		redisClient,
	}
}
