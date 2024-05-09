package logic

import (
	"djblog/dao"
	"djblog/models"
	"djblog/pkg/jwt"
)

func UserLogic(p *models.User) (string, string, error) {
	user := new(models.ParamLogic)
	user.Username = p.Username
	user.Password = p.Password
	err := dao.QueryUser(user)
	if err != nil {
		return "", "", err
	}
	err = dao.QueryPassword(user)
	if err != nil {
		return "", "", err
	}

	//生成token
	if err := dao.SelectID(user); err != nil {
		return "", "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}
