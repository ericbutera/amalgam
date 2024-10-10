package internal

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const SwaggerUri = "/swagger/index.html"

func Swagger() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerfiles.Handler)
}
