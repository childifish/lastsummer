package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey             string `json:"app_key"`
	AppSec             string `json:"app_sec"`
	jwt.StandardClaims        //里面有ExpiresAt:签发时间和Issuer:签发者
}

//md5加密，相比Base64具有不可逆性//网上抄的
func mdV(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//生成token
func GenerateToken(appKey, appSec string) (string, error) {
	now := time.Now()
	ExpiresTime := now.Add(720000 * time.Second)
	claims := Claims{
		AppKey: appKey,
		AppSec: mdV(appSec),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresTime.Unix(),
			Issuer:    "lastsummer",
		},
	}
	fmt.Println(claims)
	//jwt.NewWithClaims(转码方式,对象)经典sha256
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(tokenClaims)
	return tokenClaims.SignedString(GetJWTSec())
}

func GetJWTSec() []byte {
	return []byte("childifish")
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSec(), nil
	})
	if tokenClaims != nil {
		claim, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claim, nil
		}
	}
	return nil, err
}
