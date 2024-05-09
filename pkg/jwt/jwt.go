package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TokenExpireDuration 过期时间365
const TokenExpireDuration = time.Hour * 24 * 365

var MySecret = []byte("秋天悄悄的过去")

type MyClaim struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义payload,生成token
func GenToken(userID int64, username string) (atoken string, rtoken string, err error) {
	//创建我们自己声明的payload数据
	p := new(MyClaim)
	p.Username = username
	p.UserID = userID
	p.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	p.Issuer = "djblog"

	//access token,使用secret签名并获得完整的token
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, p).SignedString(MySecret)

	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "djblog",
	}).SignedString(MySecret)

	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaim, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	//var token *jwt.Token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaim{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*MyClaim); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新access token
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	//refresh token无效直接返回
	_, err = jwt.Parse(rToken, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return
	}
	//从旧access中提取数据
	var claim *MyClaim
	claim = new(MyClaim)
	_, err = jwt.ParseWithClaims(aToken, claim, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	//*jwt.ValidationError 是 jwt 库（JSON Web Token 库）中定义的一种特殊错误类型
	var v *jwt.ValidationError
	_ = errors.As(err, &v)
	//这个错误表示JWT的过期时间（expiration time）已经过期
	if v.Errors == jwt.ValidationErrorExpired {
		newAToken, newRToken, err = GenToken(claim.UserID, claim.Username)
	}
	return
}
