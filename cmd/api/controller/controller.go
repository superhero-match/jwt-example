package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/jwt-example/cmd/api/service"
	"github.com/jwt-example/internal/config"
)

// Controller holds the Controller data.
type Controller struct {
	Service *service.Service
}

// NewController returns new controller.
func NewController(cfg *config.Config) (*Controller, error) {
	srv, err := service.NewService(cfg)
	if err != nil {
		return nil, err
	}

	return &Controller{
		Service: srv,
	}, nil
}

// RegisterRoutes registers all the superhero_suggestions API routes.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/jwt_example")

	// Middleware.
	// sr.Use(c.Authorize)

	// Routes.
	sr.POST("/login", ctl.Login)
	sr.POST("/todo", ctl.TokenAuthMiddleware(), ctl.CreateTodo)
	sr.POST("/logout", ctl.TokenAuthMiddleware(), ctl.Logout)
	sr.POST("/token/refresh", ctl.Refresh)

	return router
}