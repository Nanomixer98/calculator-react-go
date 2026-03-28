package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "calculator-server/docs"
)

type Router struct {
	ctrl *CalculatorController
}

func NewRouter(c *CalculatorController) *Router {
	return &Router{
		ctrl: c,
	}
}

func (r *Router) RegisterRoutes(engine *gin.Engine) {
	// Swagger documentation route
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api")
	{
		api.POST("/add", r.ctrl.Add)
		api.POST("/subtract", r.ctrl.Subtract)
		api.POST("/multiply", r.ctrl.Multiply)
		api.POST("/divide", r.ctrl.Divide)
		api.POST("/negate", r.ctrl.Negate)
		api.POST("/percentage", r.ctrl.Percentage)
	}
}

