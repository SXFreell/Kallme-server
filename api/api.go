package api

import (
	"kallme/config"
	"kallme/model"
	"kallme/service"
	"kallme/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var reqBody model.LoginReq

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1011,
			"message": "fail to bind json, error: " + err.Error(),
		})
		return
	}

	// 检查用户名是否存在
	if !service.CheckUsernameExist(reqBody.Username) {
		c.JSON(http.StatusOK, gin.H{
			"code":    2011,
			"message": "username not exist",
		})
		return
	}

	// 检查密码是否匹配
	if !service.CheckPasswordMatch(reqBody.Username, reqBody.Password) {
		c.JSON(http.StatusOK, gin.H{
			"code":    2012,
			"message": "password not match",
		})
		return
	}

	// 生成token
	token, err := service.GenerateToken(reqBody.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    2013,
			"message": "fail to generate token",
		})
		return
	}

	// 返回token
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"token": token,
		},
	})

}

func Register(c *gin.Context) {
	var reqBody model.LoginReq

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1021,
			"message": "fail to bind json, error: " + err.Error(),
		})
		return
	}

	// 校验用户名是否合规
	if !service.CheckUsernameValid(reqBody.Username) {
		c.JSON(http.StatusOK, gin.H{
			"code":    2021,
			"message": "username not valid",
		})
		return
	}

	// 检查用户名是否存在
	if service.CheckUsernameExist(reqBody.Username) {
		c.JSON(http.StatusOK, gin.H{
			"code":    2022,
			"message": "username already",
		})
		return
	}

	var user model.User
	user.UUID = utils.GenerateShortUUID()
	user.Username = reqBody.Username
	user.Password = reqBody.Password

	if !service.InsertUser(user) {
		c.JSON(http.StatusOK, gin.H{
			"code":    2023,
			"message": "fail to insert user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}

func LinkWS(c *gin.Context) {
	token := c.Query("token")

	config.Log.Info("token: ", token)

	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1031,
			"message": "token is empty",
		})
	}

	service.HandleWebSocket(c, token)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}

func SendMessage(c *gin.Context) {
	var reqBody model.SendMessageReq

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1041,
			"message": "fail to bind json, error: " + err.Error(),
		})
		return
	}

	token := reqBody.Token
	msg := reqBody.Msg

	if err := service.SendMessage(token, msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1042,
			"message": "fail to send message, error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"code":    1051,
	// 		"message": "invalid token",
	// 	})
	// }

}
