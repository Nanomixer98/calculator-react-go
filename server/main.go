// @title           Calculator HTTP API
// @version         1.0
// @description     This is a calculator server implementing a Hexagonal Architecture.

// @host      localhost:8080
// @BasePath  /api

package main

import (
	"fmt"
	"log"

	"calculator-server/builder"
)

func main() {
	server := builder.BuildServer()

	port := ":8080"
	fmt.Printf("Calculator Gin server listening on http://localhost%s\n", port)

	if err := server.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
