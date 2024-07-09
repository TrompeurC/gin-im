package router

import (
	"gin-im/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @Tags test
// @Success 200 {string} Helloworld
// @Router /index [get]
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "sss",
	})
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/index", test)
	return r
}
