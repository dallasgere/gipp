package image

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route Groups
func ImageRoutes(router *gin.RouterGroup) {
	imageGroup := router.Group("/image")
	{
		imageGroup.POST("/resize", resize)
	}
}

// resize endpoint
func resize(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not open form file"})
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not read bytes of form file"})
		return
	}

	result, err := ResizeImage(buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not resize the image"})
		return
	}

	contentType := http.DetectContentType(result)
	c.Data(http.StatusOK, contentType, result)
}
