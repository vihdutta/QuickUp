package main

import (
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
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
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("C:\\Users\\Duttas\\Desktop\\LR Exports")
	if err != nil {
		log.Fatal(err)
	}

	gui()

	<-done
}

func gui() {
	a := app.New()
	w := a.NewWindow("Auto Screenshot Opener")
	w.Resize(fyne.NewSize(500, 500))

	btn := widget.NewButton("Open .jpg & .Png", func() {
		fileDialog := dialog.NewFileOpen(
			func(raw_img_data fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(raw_img_data)
				res := fyne.NewStaticResource(raw_img_data.URI().Name(), data)
				img := canvas.NewImageFromResource(res)

				w := fyne.CurrentApp().NewWindow(raw_img_data.URI().Name())
				w.SetContent(img)
				w.Resize(fyne.NewSize(500, 500))
				w.Show()
			}, w)
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg"}))
		fileDialog.Show()
	})

	w.SetContent(btn)
	w.ShowAndRun()
}
