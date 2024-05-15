package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobkc/ext"
	"github.com/gobkc/fit/conf"
	"github.com/gobkc/fit/driver"
	"github.com/gobkc/jwt"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	*gin.Engine
	d driver.Driver
	c *conf.Conf
}

func NewServer() *Server {
	s := &Server{}

	s.c = conf.GetConf()
	s.d = driver.NewDriver()

	s.Engine = gin.Default()
	s.LoadRouters()

	s.Run(s.c.RestAddr)

	return s
}

func (s *Server) JSON(c *gin.Context, data any) {
	if c.Request == nil {
		return
	}
	mb, _ := ext.MarshalGzipJson(data)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Content-Encoding", "gzip")
	c.Writer.Header().Set("Vary", "Accept-Encoding")
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%v", len(mb)))
	if authorization := c.GetHeader(`Authorization`); authorization != `` {
		tokenHeaders := strings.Split(authorization, ` `)
		token := ``
		if len(tokenHeaders) > 1 {
			token = strings.TrimSpace(tokenHeaders[1])
		}
		j := jwt.NewJwt(func() (secret, alg, typ string, expired time.Duration) {
			secret = s.c.JwtSalt
			return
		})
		if err := j.Refresh(&token); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{`error`: 1, `msg`: `StatusForbidden`, `more`: err.Error()})
			return
		}
		c.Writer.Header().Set("Authorization", `Bearer `+token)
	}
	c.Writer.Write(mb)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Status()
}

func (s *Server) JSONWithHTTPCode(c *gin.Context, httpCode int, data any) {
	s.JSON(c, data)
	c.Writer.WriteHeader(httpCode)
	c.Writer.Status()
}
