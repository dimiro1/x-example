package app

import (
	"io"
	"net/http"

	"github.com/dimiro1/x/xhttp"
	"github.com/dimiro1/x/xlog"
	"github.com/dimiro1/x/xtemplate"
)

func Index(compress xhttp.CompressMiddleware, logger xlog.OptionalLogger) xhttp.RouteMapping {
	return xhttp.RouteMapping{
		Route: &xhttp.Route{
			Path:   "/",
			Method: http.MethodGet,
			Middleware: []xhttp.Middleware{
				compress,
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

				template.ExecuteTemplate(w, "about.html", "Hello From About")
			}),
		},
	}
}
