package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/vihdutta/QuickUp/utils"
)

func main() {
	a := app.New()
	w := a.NewWindow("QuickUp")
	w.Resize(fyne.NewSize(800, 800))

	go utils.Toolbar(w)
	go utils.Watcher(w, "")
	w.ShowAndRun()
}
