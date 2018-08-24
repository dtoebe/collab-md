package methodcheck

import (
	"log"
	"net/http"

	"github.com/dtoebe/collab-md/internal/middleware"
)

func NotImplemented(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func AllowedMethods(methods ...string) middleware.Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Method check:", r.Method)
			if ok := checkStr(r.Method, methods); !ok {
				log.Println("Bad Method:", r.Method)
				NotImplemented(w)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}

func checkStr(m string, a []string) bool {
	for _, v := range a {
		if m == v {
			return true
		}
	}

	return false
}
