package builder

import (
	"calculator-server/app"
	"calculator-server/controller"

	"github.com/gin-gonic/gin"
)

// corsMiddleware provides a basic CORS mechanism for the router
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func BuildServer() *gin.Engine {
	engine := gin.Default()
	engine.Use(corsMiddleware())

	calcApp := app.NewCalculatorApp()
	calcController := controller.NewCalculatorController(calcApp)

	router := controller.NewRouter(calcController)
	router.RegisterRoutes(engine)

	return engine
}
