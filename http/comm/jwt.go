package comm

import (
	"focus/http/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("##HI~focus@By-UnKown$233##") // 加密/解密 密钥

// LoginToken结构体
// 用于校验用户是否已更改密码等等
// TODO 以后可以验证异地登录或不同设备
type Claims struct {
	Uid      int
	Username string
	Password string
	jwt.StandardClaims
}

// 生成Token: LoginToken
// 有效期: 7天 (理论上每次登录都会重新创建token)
func GenerateLoginToken(account *model.Account) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 7 * time.Hour)
	issuer := "focus"
	claims := Claims{
		Uid:      account.Uid,
		Username: account.Username,
		Password: account.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
			Subject:   "login token",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	return token, err
}

// 解析Token: LoginToken
// 获取token的内容
func ParseLoginToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
