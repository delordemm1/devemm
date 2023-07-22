package web

import (
	"net/http"

	"github.com/delordemm1/devemm-go/static"
	"github.com/justinas/alice"
)

func (a *app) routes() http.Handler {
	baseMiddleware := alice.New(a.recoverPanic)
	webMiddleware := alice.New(
		a.inertiaManager.Middleware,
	)

	mux := http.NewServeMux()
	mux.Handle("/", webMiddleware.ThenFunc(a.homeHandler))

	var fileServer http.Handler

	if a.debug {
		fileServer = http.FileServer(http.Dir("./static/"))
	} else {
		staticFS := http.FS(static.Static)
		fileServer = http.FileServer(staticFS)
	}

	mux.Handle("/build/", fileServer)
	mux.Handle("/assets/", fileServer)
	mux.Handle("/favicon.ico", fileServer)

	return baseMiddleware.Then(mux)

	// var head string
	// head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	// if head == "" {
	// 	webMiddleware.ThenFunc(a.homeHandler).ServeHTTP(w, r)
	// 	// a.HomeHandler.ServeHTTP(w, r)
	// 	return
	// }
	// log.Println(head)
	// http.Error(w, "Not Found", http.StatusNotFound)
}
