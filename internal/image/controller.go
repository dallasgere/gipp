package image

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterImageRoutes Route Groups
func RegisterImageRoutes(router *gin.RouterGroup) {
	imageGroup := router.Group("/image")
	{
		imageGroup.POST("/resize", resize)
	}
}

// resize: resizes an image given height and width
func resize(c *gin.Context) {
	var req ResizeImageRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	file, err := req.File.Open()
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

	result, err := ResizeImage(buf, req.Width, req.Height)
	if err != nil {
		if errors.Is(err, ErrImageTypeNotAllowed) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not resize image"})
		return
	}

	contentType := http.DetectContentType(result)
	c.Data(http.StatusOK, contentType, result)
}
