package router

import (
	"gin-im/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	r.POST("/user", controller.UserController.Create)

	r.GET("/user", controller.UserController.List)

	r.DELETE("/user/:id", controller.UserController.Delete)
}
