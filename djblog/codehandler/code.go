package codehandler

type Recode int64

const (
	CodeSuccess = 1000 + iota
	CodeSeverBusy
	CodeLogicFailed
	CodeNeedLogic
	CodeInvalidToken
)

var CodeMsg = map[Recode]string{
	CodeSuccess:      "SUCCESS",
	CodeSeverBusy:    "服务繁忙",
	CodeLogicFailed:  "账号或密码错误",
	CodeNeedLogic:    "请登录用户",
	CodeInvalidToken: "无效TOKEN",
}

func (recode Recode) Msg() string {
	msg := CodeMsg[recode]
	return msg
}
