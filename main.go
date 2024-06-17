package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	m := newMilvusmetar()
	m.loadUI(app)
	app.Run()
}
