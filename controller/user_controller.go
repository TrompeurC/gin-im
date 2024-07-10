package controller

import (
	"gin-im/models"
	"gin-im/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var UserController = new(userController)

type userController struct {
}

// 创建用户
func (u userController) Create(c *gin.Context) {
	user := models.UserBasic{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	err = service.UserService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "创建用户成功！",
	})
}

// 获取用户列表
func (u userController) List(c *gin.Context) {
	userList, err := service.UserService.GetUserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "获取用户列表成功！",
		"data": userList,
	})
}

// 删除用户
func (u userController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请传入用户id",
		})
		return
	}
	userId, _ := strconv.Atoi(id)
	err := service.UserService.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除用户成功！",
	})
}
