package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenService struct {
}

var signKey = []byte("aga222n3223rgnaow2f352t24g2g4")

// 產生token
func (this TokenService) GenerateToken(userid string) (string, error) {
	// token 有效時間一天
	oneDayDuration, _ := time.ParseDuration("24h")
	exp := time.Now().Add(oneDayDuration).Unix()
	iat := time.Now().Unix()
	jti := uuid.New()
	iss := "larvata"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    iss,
		"userId": userid,
		"iat":    iat,
		"exp":    exp,
		"jti":    jti,
	})
	tokenString, err := token.SignedString(signKey)
	return tokenString, err
}

// 產生refresh token
func (this TokenService) GenerateRefreshToken(userid string) (string, error) {
	// token 有效時間一天
	oneDayDuration, _ := time.ParseDuration("720h")
	exp := time.Now().Add(oneDayDuration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userid,
		"exp":    exp,
	})
	tokenString, err := token.SignedString(signKey)
	return tokenString, err
}

// 驗證token
func (this TokenService) ValidateToken(uncheckToken string) (bool, error) {
	token, err := jwt.Parse(uncheckToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signKey, nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return true, nil
	} else {
		return false, err
	}
}

// 取得token中的資訊
func (this TokenService) getTokenInfo(authToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		isNotExpire := claims.VerifyExpiresAt(time.Now().Unix(), false)
		if isNotExpire {
			return claims, nil
		}
	} else {
		return nil, err
	}
	return nil, err
}
