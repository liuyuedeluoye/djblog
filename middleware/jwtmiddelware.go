package middleware

import (
	"djblog/codehandler"
	"djblog/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

const Ctxuserinfo = "userID"

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		//无token
		if authHeader == "" {
			//c.JSON(http.StatusOK, gin.H{
			//	"code": 2003,
			//	"msg":  "请求头中auth为空",
			//})
			codehandler.ResponseError(c, codehandler.CodeNeedLogic)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			//无效token
			//c.JSON(http.StatusOK, gin.H{
			//	"code": 2004,
			//	"msg":  "请求头中auth格式有误",
			//})
			codehandler.ResponseError(c, codehandler.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			//c.JSON(http.StatusOK, gin.H{
			//	"code": 2005,
			//	"msg":  "无效的Token",
			//})
			codehandler.ResponseError(c, codehandler.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userid信息保存到请求的上下文c上
		c.Set(Ctxuserinfo, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get(Ctxuserinfo)来获取当前请求的用户信息
	}
}
