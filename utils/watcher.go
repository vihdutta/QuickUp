package utils

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

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

					imageData := ImageDatas(event.Name)
					labelColorModel, labelDims := canvas.NewText(imageData[0], color.White), canvas.NewText(imageData[1], color.White)
					labelColorModel.Alignment = fyne.TextAlignCenter
					labelDims.Alignment = fyne.TextAlignCenter
					imgData := container.NewVBox(labelColorModel, labelDims)

					w.SetContent(container.NewBorder(Toolbar(w), imgData, nil, nil, img))
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

func ImageDatas(path string) [2]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(err)
	}

	return [2]string{fmt.Sprint("Color Model: ", &image.ColorModel),
		fmt.Sprint("Width: ", image.Width, "    |    ", "Height: ", image.Height)}
}
