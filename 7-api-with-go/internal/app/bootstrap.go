package app

import (
	"github.com/gin-gonic/gin"
	noteModule "mobin.dev/internal/domain/notes/module"
	noteV1 "mobin.dev/internal/domain/notes/presentation/http/v1"
	noteV2 "mobin.dev/internal/domain/notes/presentation/http/v2"
)

func (a *App) RunServer() {
	router := gin.Default()

	v1 := router.Group("/v1")

	notes := noteModule.Init(a.mysql)

	{
		noteV1.RegisterNotesRouter(v1, notes.V1)
	}

	v2 := router.Group("/v2")
	{
		noteV2.RegisterNotesRouter(v2, notes.V2)
	}

	router.Run(":4000")
}
