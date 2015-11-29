package main

import (
	"flag"
	"net/http"

	af "github.com/elazarl/go-bindata-assetfs"
	ur "github.com/unrolled/render"
	"github.com/zenazn/goji"
)

//go:generate go-bindata 02_dist/...

type context struct {
	Title  string
	Result string
}

func main() {
	flag.Set("bind", ":3981")
	goji.DefaultMux.Get("/", index)
	goji.DefaultMux.Handle("/static/*", static)
	goji.Serve()
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := context{
		Title:  "This is a title",
		Result: "This is a result",
	}

	re := setLayout()
	re.HTML(w, http.StatusOK, "index", ctx)
}

func static(w http.ResponseWriter, r *http.Request) {
	http.FileServer(
		&af.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "02_dist/"})
}

func setLayout() *ur.Render {
	return ur.New(ur.Options{
		Asset:      Asset,
		AssetNames: AssetNames,
		Directory:  "02_dist/template",
		Layout:     "layout",
	})
}
