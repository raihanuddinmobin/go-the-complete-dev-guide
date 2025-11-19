package app

import (
	"github.com/gin-gonic/gin"
	"mobin.dev/internal/domain/notes/module"
	notesV1 "mobin.dev/internal/domain/notes/presentation/v1"
	"mobin.dev/internal/infrastructure/logger"
	"mobin.dev/internal/middleware"
)

func (a *App) StartServer() {

	router := gin.Default()

	notes := module.Init(a.pg)

	router.Use(middleware.TraceMiddleware())
	router.Use(middleware.ContentTypeMiddleware())

	v1 := router.Group("/v1")
	{
		notesV1.RegisterNotesRouter(v1, notes.V1)
	}

	_ = logger.Init()
	router.Run(":4000")

}
