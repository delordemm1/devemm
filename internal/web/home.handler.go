package web

import "net/http"

type Home struct {
}

func (a *app) homeHandler(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(a.inertiaManager.WithViewData(r.Context(), "title", "devemm-go"))

	a.inertiaManager.Render(w, r, "home", nil)
	return
	// w.Write([]byte("Hello World"))
}
