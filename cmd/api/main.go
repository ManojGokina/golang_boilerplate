package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	
	cfg := config.Load()

	// Initialize logger
	log := logger.New(cfg.LogLevel)

	// Initialize database
	db, err := database.NewMongoDB(cfg.MongoURL, cfg.MongoDB)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo, log)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService, log)

	// Setup router
	router := setupRouter(cfg, userHandler, log)

	// Start server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Info("Server starting on port ", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Info("Server exited")
}

func setupRouter(cfg *config.Config, userHandler *handler.UserHandler, log logger.Logger) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Middleware
	router.Use(middleware.Logger(log))
	router.Use(middleware.Recovery(log))
	router.Use(middleware.CORS())
	router.Use(middleware.RateLimit())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "timestamp": time.Now().Unix()})
	})

	// API routes
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
			users.GET("", userHandler.ListUsers)
		}
	}

	return router
}