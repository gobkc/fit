package rest

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobkc/do"
	"github.com/gobkc/fit/docs"
	"github.com/gobkc/fit/middleware"
	"github.com/gobkc/fit/static"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

func (s *Server) LoadRouters() {
	// set cors
	s.Engine.Use(middleware.WithCORS())

	// public routers
	s.public()

	// routers that require authorization
	s.routers()

	// static routers
	s.static()

	// docs
	s.Engine.Any(s.c.Version+`/docs/*any`, ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Base URL: swagger docs basic setting
	docs.SwaggerInfo.BasePath = "/" + s.c.Version
	docs.SwaggerInfo.Host = do.OneOr(os.Getenv("FIT-URL"), s.c.RestAddr)
}

// Routers that require authorization
func (s *Server) routers() {
	//r := s.Engine.Group(s.c.Version + `/a/`)
	//r.Use(middleware.WithCORS(), middleware.JwtValidation)
	//r.GET(`me`, s.Me)
}

// public routers (No authorization required)
func (s *Server) public() {
	r := s.Engine.Group(s.c.Version + `/p/`)
	r.GET(`version`, s.Version)
	r.GET(`health`, s.HealthCheck)
	r.POST(`new-note`, s.NewNote)
	r.GET(`list-cate`, s.ListCate)
	r.POST(`new-cate`, s.NewCate)
	r.GET(`/:cate/list-note`, s.ListNote)
}

// static routers (No authorization required)
func (s *Server) static() {
	s.Engine.Use(middleware.WithCORS()).GET(`/`, func(c *gin.Context) {
		b, err := static.GetWebByte(`index.html`)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		contentType := "text/html"
		c.Writer.Header().Set("content-type", contentType)
		c.Writer.Write(b)
	})
	staticMiddleware := func(c *gin.Context) {
		filePath := c.Param("filepath")
		fp := fmt.Sprintf(`assets%s.gz`, filePath)
		isGzip := true
		if filePath == `` {
			fp = `index.html`
			filePath = fp
			isGzip = false
		}
		if c.FullPath() == "/favicon.ico" {
			fp = "favicon.ico"
			filePath = fp
			isGzip = false
		}
		b, err := static.GetWebByte(fp)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				fp = fmt.Sprintf(`assets%s`, c.Param("filepath"))
				b, err = static.GetWebByte(fp)
				isGzip = false
			}
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
		}
		split := strings.Split(filePath, ".")
		var contentType string
		switch split[len(split)-1] {
		case "html", "htm", "xhtml":
			contentType = "text/html"
		case "css":
			contentType = "text/css"
		case "js":
			contentType = "application/javascript"
		case "svg":
			contentType = "image/svg+xml"
		case "gif", "png", "jpg", "jpeg", "bmp", "ico":
			contentType = "image/*"
		default:
			contentType = "text/plain"
		}
		c.Writer.Header().Set("Content-Type", contentType)
		if isGzip {
			c.Writer.Header().Set("Content-Encoding", "gzip")
			c.Writer.Header().Set("Vary", "Accept-Encoding")
			c.Writer.Header().Set("Content-Length", fmt.Sprintf("%v", len(b)))
		}
		c.Writer.Write(b)
	}
	s.Engine.Use(middleware.WithCORS()).GET(`assets/*filepath`, staticMiddleware)
	s.Engine.NoRoute(staticMiddleware)
}
