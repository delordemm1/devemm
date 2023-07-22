package cmd

import "github.com/delordemm1/devemm-go/internal/web"

func WebServe() {
	web.Serve(true, ":4001", "http://localhost:4001")
}
