package main

import (
	"flag"
	"net/http"

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

func setLayout() *ur.Render {
	return ur.New(ur.Options{
		Asset:      Asset,
		AssetNames: AssetNames,
		Directory:  "02_dist/template",
		Layout:     "layout",
	})
}
