package app

import (
	"io"
	"net/http"

	"github.com/dimiro1/x/xhttp"
	"github.com/dimiro1/x/xlog"
	"github.com/dimiro1/x/xtemplate"
	"go.uber.org/fx"
)

type IndexMiddleware struct {
	fx.In

	Compress xhttp.Middleware `name:"x_compress_middleware"`
}

func Index(indexMiddleware IndexMiddleware, logger xlog.OptionalLogger) xhttp.RouteMapping {
	return xhttp.GET("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if xlog.IsProvided(logger) {
			logger.Logger.Println("serving /")
		}
		io.WriteString(w, "Index Page")
	}), xhttp.WithMiddleware(indexMiddleware.Compress))
}

func About(logger xlog.OptionalLogger, template xtemplate.Template) xhttp.RouteMapping {
	return xhttp.GET("/about", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if xlog.IsProvided(logger) {
			logger.Logger.Println("serving /about")
		}

		template.Template.ExecuteTemplate(w, "about.html", "Hello From About")
	}))
}
