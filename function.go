package image_gen

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gin-gonic/gin"
	"github.com/tim-dewit/image-gen/api"
)

func init() {
	r := gin.Default()
	r.POST("/", api.GenerateHandler)

	functions.HTTP("generate", r.Handler().ServeHTTP)
}
