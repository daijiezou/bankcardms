package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

var (
	MySecret = make([]byte, 20)
)

func randString(len int) []byte {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return bytes
}

func InitSecret() {
	MySecret = randString(20)
}

type MyClaims struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

func GenToken(userId int64, username string) (accessToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserId:   userId,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(14400)).Unix(), // 过期时间
			Issuer:    "gemini-userauth",                                         // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	accessTokenStr, err := token.SignedString(MySecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return accessTokenStr, nil
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
