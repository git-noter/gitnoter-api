package httpservice

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vivekweb2013/batnoter/internal/applicationconfig"
)

// Start the http server
func Run(applicationconfig *applicationconfig.ApplicationConfig) error {
	gin.SetMode(gin.ReleaseMode)
	if applicationconfig.Config.HTTPServer.Debug {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	noteHandler := NewNoteHandler(applicationconfig.NoteService)
	v1 := router.Group("api/v1")
	v1.GET("/note/:id", noteHandler.GetNote)
	v1.POST("/note", noteHandler.CreateNote)
	v1.PUT("/note/:id", noteHandler.UpdateNote)
	v1.DELETE("/note/:id", noteHandler.DeleteNote)

	address := net.JoinHostPort(applicationconfig.Config.HTTPServer.Host, applicationconfig.Config.HTTPServer.Port)
	server := http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
