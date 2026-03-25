package http

import (
	"context"
	"errors"
	_ "github.com/getLynx-tech/lynx/cmd/docs"
	"github.com/getLynx-tech/lynx/internal/presentation/http/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(
	rootHandler *handler.RootHandler,
	deviceHandler *handler.DeviceHandler,
	scaleHandler *handler.ScaleHandler,
) *Server {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.GET("/", rootHandler.GetRoot)
	devices := engine.Group("/devices")
	{
		devices.POST("/position", deviceHandler.UpdatePosition)
	}

	scales := engine.Group("/scales")
	{
		scales.POST("", scaleHandler.CreateScale)
		scales.GET("", scaleHandler.GetScale)
	}

	return &Server{engine: engine}
}

func RegisterServer(lc fx.Lifecycle, s *Server) {
	server := &http.Server{
		Addr:    ":9090",
		Handler: s.engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("failed to start server: %v", err)
				}
			}()
			log.Printf("Server listening on port 9090")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Printf("Shutting down server...")
			return server.Shutdown(ctx)
		},
	})
}
