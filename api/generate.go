package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim-dewit/image-gen/generate"
	"image/png"
	"time"
)

func GenerateHandler(c *gin.Context) {
	var req generate.GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	img, err := generate.Render(&req)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.Writer.Header().Add("Content-Type", "image/png")

	es := time.Now()
	err = png.Encode(c.Writer, img)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	fmt.Printf("Encoded image in %s\n", time.Since(es))

	c.Status(200)
}
