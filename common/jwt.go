package common

import (
	"Gin_Vue_Demo/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 登陆成功后发放token

func ReleaseToken(user model.User) (string, error) {
	// 设置 token 有效期 为 7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	// 里面就是 token密钥 的内容，解密后包括 claims 里面的所有信息
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "l mx",
			Subject:   "user token",
		},
	}

	// 切记使用 HS256 算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用上面的 jwtKey jwt密钥生成token
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

// 用来拿token里面的信息的

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	// 函数返回 token， 把 token 里面村粗的信息存到 claims 里面
	token, err := jwt.ParseWithClaims(tokenString, claims,
		// 函数
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	return token, claims, err
}
