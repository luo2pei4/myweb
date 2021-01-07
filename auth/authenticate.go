package authenticate

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Mary")

// Claims 用于payload
type Claims struct {
	UserID string `json:"username"`
	jwt.StandardClaims
}

// DistributeToken 分发Token
func DistributeToken(response http.ResponseWriter, userID string) (returnCode int, err error) {

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	http.SetCookie(response, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return http.StatusOK, nil
}

// RedistributeToken 重新分发Token
func RedistributeToken(response http.ResponseWriter, userID string) (returnCode int, err error) {
	return DistributeToken(response, userID)
}

// AuthToken 验证Token
func AuthToken(request *http.Request) (claims *Claims, returnCode int) {

	// 获取认证
	cookie, err := request.Cookie("token")

	if err != nil {

		if err == http.ErrNoCookie {
			return nil, http.StatusUnauthorized
		}

		return nil, http.StatusBadRequest
	}

	tknstr := cookie.Value
	claims = &Claims{}

	token, err := jwt.ParseWithClaims(tknstr, claims, func(tkn *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {

		if err == jwt.ErrSignatureInvalid {
			return nil, http.StatusUnauthorized
		}

		return nil, http.StatusBadRequest
	}

	if !token.Valid {
		return nil, http.StatusUnauthorized
	}

	return claims, http.StatusOK
}
