package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Server storing target destinations' information
type Server struct {
	Method        string
	Path          string
	ProxyScheme   string
	ProxyPass     string
	ProxyPassPath string
	APIVersion    string
	CustomConfigs CustomConfig
}

// ReverseProxy forwards requests from upper proxy or clients to the corresponding destination
func (s Server) ReverseProxy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		director := func(req *http.Request) {
			req.Host = s.ProxyPass
			req.URL.Host = s.ProxyPass
			req.URL.Scheme = s.ProxyScheme
			// Process the specific URL format, such as: http://xxx.api.user:id
			suffixPath := strings.Replace(ctx.Request.URL.Path, s.Path, "", 1)
			req.URL.Path = s.ProxyPassPath + suffixPath
		}

		proxy := httputil.NewSingleHostReverseProxy(&url.URL{})
		proxy.Director = director
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
