package web

import (
	"log"

	"github.com/delordemm1/devemm-go/pkg/vite"
	"github.com/petaki/inertia-go"
)

type app struct {
	debug          bool
	url            string
	infoLog        *log.Logger
	errorLog       *log.Logger
	inertiaManager *inertia.Inertia
	viteManager    *vite.Manifest
	HomeHandler    *Home
}
