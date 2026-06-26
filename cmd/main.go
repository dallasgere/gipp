package main

import (
	"runtime"

	"log/slog"

	gippImage "github.com/dallasgere/gipp/internal/image"
	"github.com/gin-gonic/gin"
)

func main() {
	// Gin app
	r := gin.Default()

	numCPUs := runtime.NumCPU()
	slog.Info("Num CPU:", numCPUs)

	// global version
	v1 := r.Group("api/v1")

	// register routes
	gippImage.RegisterImageRoutes(v1)

	err := r.Run(":8080")
	if err != nil {
		slog.Error("Error starting gipp server", err.Error())
	}
}
