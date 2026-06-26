package main

import (
	"fmt"
	"runtime"

	gippImage "github.com/dallasgere/gipp/internal/image"
	"github.com/gin-gonic/gin"
)

func main() {
	// Gin app
	r := gin.Default()

	numCPUs := runtime.NumCPU()
	fmt.Println("Num CPU:", numCPUs)

	// global version
	v1 := r.Group("api/v1")

	// register routes
	gippImage.ImageRoutes(v1)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("ERROR starting gipp server, shutting down")
	}
}
