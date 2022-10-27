package web

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed all:dist
	dist embed.FS
	//go:embed dist/index.html
	indexHtml embed.FS

	distDirFS     = echo.MustSubFS(dist, "dist")
	distIndexHtml = echo.MustSubFS(indexHtml, "dist")
)

func RegisterHandler(e *echo.Echo) {
	e.FileFS("/", "index.html", distIndexHtml)
	e.StaticFS("/", distDirFS)
}
