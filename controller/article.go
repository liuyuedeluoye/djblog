package controller

import (
	"djblog/codehandler"
	"djblog/logic"
	"djblog/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//管理员添加文章

func AddArticle(c *gin.Context) {
	data := new(models.Article)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("文章参数获取失败")
		codehandler.ResponseError(c, codehandler.CodeSeverBusy)
		return
	}
	//logic处理data
	if err := logic.AddArticle(data); err != nil {
		zap.L().Error("文章添加失败", zap.Error(err))
		codehandler.ResponseError(c, codehandler.CodeSeverBusy)
		return
	}
	codehandler.ResponseError(c, codehandler.CodeSuccess)

}

//删除文章

func DeleteArticle(c *gin.Context) {
	//删除文章
	//1.获取文章id
	//logic层处理错误
	//数据库执行操作
	//可以在前端管理员展示页面展示所有文章并且标出id,只要输入id就可以删除,
	id := c.Param("id")
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		zap.L().Error("帖子参数处理失败", zap.Error(err))
		codehandler.ResponseError(c, codehandler.CodeSeverBusy)
		return
	}
	//逻辑处理,根据id锁定文章
	err = logic.DeleteArticle(postID)
	if err != nil {
		zap.L().Error("删除文章失败", zap.Error(err))
		codehandler.ResponseError(c, codehandler.CodeSeverBusy)
		return
	}
	//返回响应
	codehandler.ResponseError(c, codehandler.CodeSuccess)
}

// GetArticleList 显示文章列表
//一部分是管理员显示列表,不需要显示内容只要有文章标题即可
//一部分是用户界面显示文章

func GetArticleListUser(c *gin.Context) {
	//逻辑处理,根据id查询内容
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		limit  int64
		offset int64
		err    error
	)
	offset, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		offset = 0
	}
	limit, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		limit = 0
	}

	data, err := logic.GetPostList(offset, limit)

	if err != nil {
		zap.L().Error("获取帖子内容失败", zap.Error(err))
		codehandler.ResponseError(c, codehandler.CodeSeverBusy)
		return
	}

	//返回响应
	codehandler.ResponseSuccessWithMsg(c, data)
}
func GetArticleListRoot(c *gin.Context) {

}

// GetArticle 查询文章
func GetArticle(c *gin.Context) {
	//controller,获取输入ID,根据id查询帖子
	post_id := c.Param("id")
	postId, err := strconv.ParseInt(post_id, 10, 64)
	if err != nil {
		codehandler.ResponseError(c, codehandler.CodeSeverBusy)
	}

	data, err := logic.GetArticle(postId)

	if err != nil {
		zap.L().Error("获取帖子内容失败", zap.Error(err))
		codehandler.ResponseError(c, codehandler.CodeSeverBusy)
		return
	}

	//返回响应
	codehandler.ResponseSuccessWithMsg(c, data)
	// logic

}
