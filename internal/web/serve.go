package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/delordemm1/devemm-go/pkg/vite"
	"github.com/delordemm1/devemm-go/static"
	"github.com/delordemm1/devemm-go/ui/views"
	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/cli"
)

func Serve(debug bool, addr, url string) {
	mixManager, inertiaManager, err := newMixAndInertiaManager(debug, url)
	if err != nil {
		cli.ErrorLog.Fatal(err)
	}
	a := &app{
		debug:          debug,
		viteManager:    mixManager,
		inertiaManager: inertiaManager,
		url:            url,
		infoLog:        cli.InfoLog,
		errorLog:       cli.ErrorLog,
		HomeHandler:    new(Home),
	}
	srv := &http.Server{
		Addr:     addr,
		ErrorLog: cli.ErrorLog,
		Handler:  a.routes(),
	}
	a.infoLog.Printf("Starting server on "+cli.Green("%s"), addr)
	log.Fatal(srv.ListenAndServe())
	// http.ListenAndServe(addr, a)

}
func newMixAndInertiaManager(debug bool, url string) (*vite.Manifest, *inertia.Inertia, error) {
	viteManager, _ := vite.New("", "./static", "build", "")

	var version string
	var err error

	if debug {
		version, err = viteManager.Hash("build")
		if err != nil {
			return nil, nil, err
		}
	} else {
		version, err = viteManager.HashFromFS("build", static.Static)
		if err != nil {
			return nil, nil, err
		}
	}
	inertiaManager := inertia.NewWithFS(url, "app.gohtml", version, views.Templates)
	inertiaManager.Share("title", "devemm-go")

	inertiaManager.ShareFunc("vite", func(path string) (interface{}, error) {
		viteEmbeds, _ := viteManager.LoadViteEmbed(path)
		if err != nil {
			return nil, err
		}
		return template.HTML(viteEmbeds), nil
	})

	return viteManager, inertiaManager, nil
}
