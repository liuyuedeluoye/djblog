package dao

import (
	"djblog/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func QueryUser(p *models.ParamLogic) error {
	var username string
	sqlstr := "select username from user where username = ?"
	if err := db.QueryRow(sqlstr, p.Username).Scan(&username); err != nil {
		zap.L().Error("QueryUser failed", zap.Error(err))
		return err
	}
	//fmt.Printf("username:%s\n", username)
	//fmt.Printf("username:%s\n", p.Username)
	return nil
}
func QueryPassword(p *models.ParamLogic) error {
	sqlstr := "select password from user where username = ?"
	var password string
	if err := db.Get(&password, sqlstr, p.Username); err != nil {
		return err
	}
	//对密码进行加密并且进行比较
	//fmt.Printf("password:%s\n", password)
	//fmt.Printf("password:%s\n", p.Password)
	if p.Password != password {
		zap.L().Error("QueryPassword failed")
		return errors.New("密码错误")
	}
	return nil
}

func SelectID(p *models.ParamLogic) error {
	sqlstr := "select user_id from user where username = ?"
	var userID int64
	if err := db.Get(&userID, sqlstr, p.Username); err != nil {
		zap.L().Error("查找user_id出错", zap.Error(err))
		return err
	}
	p.UserID = userID
	return nil
}
