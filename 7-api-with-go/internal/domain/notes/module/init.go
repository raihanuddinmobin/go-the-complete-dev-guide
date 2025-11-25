package module

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"mobin.dev/internal/domain/notes/application"
	"mobin.dev/internal/domain/notes/infrastructure"
	v1 "mobin.dev/internal/domain/notes/presentation/v1"
	v2 "mobin.dev/internal/domain/notes/presentation/v1"
)

type noteModule struct {
	V1 *v1.NotesHandler
	V2 *v2.NotesHandler
}

func Init(pg *sql.DB, redisClient *redis.Client) *noteModule {
	cache := infrastructure.NewNotesCache(redisClient)
	repo := infrastructure.NewNotesRepository(pg, cache)
	service := application.NewNotesService(repo)

	// handler based on the version
	handlerV1 := v1.NewNotesHandler(service)
	handlerV2 := v2.NewNotesHandler(service)

	return &noteModule{
		V1: handlerV1,
		V2: handlerV2,
	}

}
