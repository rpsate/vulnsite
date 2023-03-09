package jwt

import (
	"vulnsite/model"
	"vulnsite/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte(utils.JwtKey)
var issuer = utils.Issuer
var tokenExpire = time.Duration(utils.TokenExpire)

type MyClaims struct {
	model.AdminInfo
	jwt.StandardClaims
}

//从MyClaims中提取用户信息
func (claims *MyClaims) GetUserInfo() *model.AdminInfo {
	return &claims.AdminInfo
}

//生成token
func SetToken(id int, username string, name string, grade int) (string, error) {
	expireAt := time.Now().Add(tokenExpire * time.Hour)
	claims := MyClaims{
		model.AdminInfo{
			ID:       id,
			Username: username,
			Name:     name,
			Grade:    grade,
		},
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: expireAt.Unix(),
			Issuer: issuer,
		},
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := reqClaims.SignedString(jwtKey)
	return token, err
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}