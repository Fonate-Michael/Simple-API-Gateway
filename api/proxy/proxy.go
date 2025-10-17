package proxy

import (
	"log"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string, prefix string) gin.HandlerFunc {

	url, err := url.Parse(target)

	if err != nil {
		log.Fatal(err)

	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	return func(context *gin.Context) {
		context.Request.URL.Path = strings.TrimPrefix(context.Request.URL.Path, prefix)
		proxy.ServeHTTP(context.Writer, context.Request)
	}

}
