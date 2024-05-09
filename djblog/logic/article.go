package logic

import (
	"djblog/dao"
	"djblog/models"
	"djblog/pkg/snowflake"
)

func AddArticle(data *models.Article) error {
	//根据雪花算法生成id,防止爆破文章id
	data.ArticleID = snowflake.GenID()

	//将数据保存到数据库
	if err := dao.AddArticle(data); err != nil {
		return err
	}

	return nil
}

func DeleteArticle(postID int64) (err error) {
	err = dao.DeleteArticle(postID)
	if err != nil {
		return err
	}
	return nil
}
func GetPostList(offset int64, limit int64) ([]*models.Article, error) {
	//获取数据库列表切片
	post, err := dao.GetPostList(offset, limit)
	//data := make([]*models.Article, 0, 2)
	return post, err
}
func GetArticle() {
	dao.GetArticle()
}
