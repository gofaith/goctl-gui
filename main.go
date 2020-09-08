package main

import (
	"log"
	"path/filepath"

	"github.com/StevenZack/openurl"

	"github.com/StevenZack/frontpage"
	"github.com/StevenZack/tools/fileToolkit"
	"github.com/gofaith/goctl-gui/views"
	"github.com/gofaith/goctl/api/jsgen"
)

func main() {
	app := frontpage.New(views.Bytes_IndexHtml, 0)
	app.Bind(genJs)
	app.OpenApp()
	app.Run()
}

func genJs(apiFile string) {
	dir := filepath.Join(fileToolkit.GetHomeDir(), "Downloads", "js")
	e := jsgen.JsGen(apiFile, dir)
	if e != nil {
		log.Println(e)
		return
	}
	openurl.Open(dir)
}
