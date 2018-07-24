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
	return xhttp.RouteMapping{
		Route: &xhttp.Route{
			Path:   "/",
			Method: http.MethodGet,
			Middleware: []xhttp.Middleware{
				indexMiddleware.Compress,
			},
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if xlog.IsProvided(logger) {
					logger.Logger.Println("serving /")
				}
				io.WriteString(w, "Index Page")
			}),
		},
	}
}

func About(logger xlog.OptionalLogger, template xtemplate.Template) xhttp.RouteMapping {
	return xhttp.RouteMapping{
		Route: &xhttp.Route{
			Path:   "/about",
			Method: http.MethodGet,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if xlog.IsProvided(logger) {
					logger.Logger.Println("serving /about")
				}

				template.Template.ExecuteTemplate(w, "about.html", "Hello From About")
			}),
		},
	}
}
