package module

import (
	"database/sql"

	"mobin.dev/internal/domain/notes/application"
	"mobin.dev/internal/domain/notes/infrastructure"
	noteV1 "mobin.dev/internal/domain/notes/presentation/http/v1"
	noteV2 "mobin.dev/internal/domain/notes/presentation/http/v2"
)

type NotesModule struct {
	V1 *noteV1.NotesHandler
	V2 *noteV2.NotesHandler
}

func Init(mysql *sql.DB) *NotesModule {
	repo := infrastructure.NewNoteRepo(mysql)
	service := application.NewNotesService(repo)

	// handler based on the version
	handlerV1 := noteV1.NewNoteHandler(service)
	handlerV2 := noteV2.NewNoteHandler(service)

	return &NotesModule{
		V1: handlerV1,
		V2: handlerV2,
	}
}
