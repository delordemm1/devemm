package cmd

import (
	"github.com/delordemm1/devemm-go/internal/web"
)

func WebServe(debug bool, addr, url string) {
	web.Serve(debug, addr, url)
}
