package utils

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/fsnotify/fsnotify"
)

func Watcher(w fyne.Window, dir string) {
	if dir == "" {
		w.SetContent(container.NewVBox(Toolbar(w)))
		return
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(dir)
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
					img := canvas.NewImageFromFile(event.Name)
					img.FillMode = canvas.ImageFillContain
					img.SetMinSize(fyne.NewSize(500, 500))
					img.ScaleMode = canvas.ImageScaleFastest

					label := canvas.NewText("hello", color.White)
					label.Alignment = fyne.TextAlignCenter

					w.SetContent(container.NewBorder(Toolbar(w), label, nil, nil, img))
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
