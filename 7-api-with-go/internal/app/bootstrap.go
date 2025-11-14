package app

import (
	"github.com/gin-gonic/gin"
	"mobin.dev/internal/domain/notes/module"
	notesV1 "mobin.dev/internal/domain/notes/presentation/v1"
)

func (a *App) StartServer() {

	router := gin.Default()

	notes := module.Init(a.pg)

	v1 := router.Group("/v1")
	{
		notesV1.RegisterNotesRouter(v1, notes.V1)
	}

	router.Run(":4000")

}
