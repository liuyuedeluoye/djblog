package dao

import (
	"djblog/models"

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

func AddArticle(data *models.Article) error {

	sqlstr := "insert into article(article_id,title,content) values (?,?,?)"
	if _, err := db.Exec(sqlstr, data.ArticleID, data.Title, data.Content); err != nil {
		return err
	}
	return nil
}

func DeleteArticle(postID int64) (err error) {
	sqlstr := "DELETE FROM article WHERE id = ?"
	res, err := db.Exec(sqlstr, postID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("文章不存在")
	}
	return nil
}
func GetPostList(offset int64, limit int64) (data []*models.Article, err error) {
	data = make([]*models.Article, 0, 2)

	//查询帖子内容
	sqlstr := "select article_id, title, content, create_time from article limit ?,?"
	if err = db.Select(&data, sqlstr, (offset-1)*limit, limit); err != nil {
		zap.L().Error("mysql get post list failed", zap.Error(err))
		return
	}
	return
}
func GetArticle(postId int64) (post *models.Article, err error) {
	data := new(models.Article)
	println(postId)
	sqlstr := "select article_id, title, content, create_time from article where article_id = ?"
	if err = db.Get(data, sqlstr, postId); err != nil {
		zap.L().Error("mysql get single post failed", zap.Error(err))
		return data, err
	}
	return data, err

}
