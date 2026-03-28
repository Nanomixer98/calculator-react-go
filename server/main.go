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
