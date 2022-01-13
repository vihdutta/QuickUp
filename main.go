package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/vihdutta/fluxum/utils"
)

func main() {
	a := app.New()
	w := a.NewWindow("Auto Screenshot Opener")
	w.Resize(fyne.NewSize(800, 800))

	go utils.Watcher(w)
	w.ShowAndRun()
}
