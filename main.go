package main

import (
	"github.com/StevenZack/frontpage"
	"github.com/gofaith/goctl-gui/views"
)

func main() {
	app := frontpage.New(views.Bytes_IndexHtml, 0)
	app.Open()
	app.Run()
}
