package codehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseErrorWithMsg 回显消息错误
func ResponseErrorWithMsg(c *gin.Context, code Recode, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Msg":  code.Msg(),
	})
}

// ResponseError 回显设置错误
func ResponseError(c *gin.Context, code Recode) {
	c.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Msh":  code.Msg(),
	})

}
func ResponseSuccessWithMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"Code": CodeSuccess,
		"Msg":  msg,
	})

}

func ResponseWithToken(c *gin.Context, atoken interface{}, rtoken interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"Code":         CodeSuccess,
		"Msh":          CodeMsg[CodeSuccess],
		"RefreshToken": rtoken,
		"AccessToken":  atoken,
	})

}
