package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"github.com/fsnotify/fsnotify"
)

var path = "C:\\Users\\Duttas\\Desktop\\LR Exports\\DSC_0005.JPG"

func main() {
	a := app.New()
	w := a.NewWindow("Auto Screenshot Opener")
	w.Resize(fyne.NewSize(500, 500))

	img := canvas.NewImageFromFile(path)

	go logic(img, w)
	w.ShowAndRun()
}

func logic(img *canvas.Image, w fyne.Window) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add("C:\\Users\\Duttas\\Desktop\\LR Exports")
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println(event.Name)
					path = event.Name
					img.File = path
					img.Refresh()
					w.SetContent(img)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	<-done
}
