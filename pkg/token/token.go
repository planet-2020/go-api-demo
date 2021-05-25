/**
 * @Description: token
 * @author zhouhongpan
 * @date 2021/5/21 17:07
 */
package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID uint
	Mobile string
	jwt.StandardClaims
}

/**
 * @Description: 生成token
 * @param subject 主题
 * @param secret 密钥
 * @param expireSecond 过期时间，秒
 * @param userId 用户ID
 * @param mobile 用户手机号
 * @return string 返回token
 * @return error 返回错误
 * @author zhouhongpan
 * @date 2021-05-25 09:59:42
 */
func GenerateToken(subject string, secret string, expireSecond int64, userId uint, mobile string) (string, error) {
	claims := Claims{
		UserID:         userId,
		Mobile:         mobile,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expireSecond,
			Subject: subject,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err := tokenClaims.SignedString([]byte(secret))
	return token,err
}

/**
 * @Description: 解密token
 * @param secret 密钥
 * @param token token字符串
 * @return *Claims
 * @return error
 * @author zhouhongpan
 * @date 2021-05-25 09:59:02
 */
func ParseToken(secret string, token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return  nil, err
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, errors.New("token无效")
}