package controller

import (
	"djblog/codehandler"
	"djblog/logic"
	"djblog/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LogicHandler 管理员登录
func LogicHandler(c *gin.Context) {
	p := new(models.User)
	if err := c.ShouldBindJSON(&p); err != nil {
		//参数校验失败
		zap.L().Error("获取参数失败", zap.Error(err))
		codehandler.ResponseError(c, codehandler.CodeSeverBusy)
		return
	}
	//if err := logic.LogicHandler(&p); err != nil {
	//	zap.L().Error("logic层 failed", zap.Error(err))
	//}
	atoken, rtoken, err := logic.UserLogic(p)
	if err != nil {
		zap.L().Error("UserLogic failed", zap.Error(err))
		codehandler.ResponseError(c, codehandler.CodeLogicFailed)
		return
	}
	codehandler.ResponseWithToken(c, atoken, rtoken)
	return
}
