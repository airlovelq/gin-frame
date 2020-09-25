package secret

import (
	"net/http"
	"reflect"
	"scoremanager/errorcode"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

const (
	SecretKey = "ScoreManager"
)

func GenerateToken(userID string, userType int) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(10)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = userID
	claims["user_type"] = userType
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stoken, err := token.SignedString([]byte(SecretKey))
	return stoken, err
}

func ParseToken(r *http.Request) (map[string]interface{}, error) {
	token_content, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	vt := make(map[string]interface{})
	if err != nil {
		return vt, err
	}
	v := reflect.ValueOf(token_content.Claims)

	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			//vt[k.Interface().(string)] = fmt.Sprintf("%v", value.Interface())
			// switch t := value.Interface().(type) {
			// case string:
			// 	vt[k.Interface().(string)] = t
			// }
			vt[k.Interface().(string)] = value.Interface()
		}
	}
	return vt, nil
}

func TokenServiceHandler(f func(c *gin.Context, tokenMap map[string]interface{})) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenMap, err := ParseToken(c.Request)
		if err != nil {
			panic(errorcode.TokenError)
		}
		f(c, tokenMap)
	}
}

// func (f TokenServiceHandler)
