package middleware

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gobkc/fit/driver"
//	"github.com/gobkc/jwt"
//	"net/http"
//	"strings"
//	"time"
//)
//
//func JwtValidation(c *gin.Context) {
//	d := driver.NewDriver()
//	j := jwt.NewJwt(func() (secret, alg, typ string, expired time.Duration) {
//		secret = d.Conf().JwtSalt
//		return
//	})
//	tokenHeaders := strings.Split(c.GetHeader(`Authorization`), ` `)
//	token := ``
//	if len(tokenHeaders) > 1 {
//		token = strings.TrimSpace(tokenHeaders[1])
//	}
//	var userClaims jwt.UserClaims
//	if err := j.Verify(token, &userClaims); err != nil {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{`error`: 1, `msg`: `StatusUnauthorized`, `more`: err.Error()})
//		return
//	}
//	c.Next()
//}
