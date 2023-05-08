package main

import (
	"embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"strconv"
	"bytes"
	"image"
)

//go:embed ab.png 
var f embed.FS

func setter(num int,lb *widget.Label ){
	lb.Text= strconv.Itoa(num)
	lb.Refresh()
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Counter")

	count:=0

	value:= widget.NewLabel(strconv.Itoa(count))

	imgfile,_:=f.ReadFile("ab.png")

	imgdata,_,_:= image.Decode(bytes.NewReader(imgfile))

	content := container.New(layout.NewHBoxLayout(), layout.NewSpacer(),value,layout.NewSpacer())
	img := canvas.NewImageFromImage(imgdata)


	img.FillMode = canvas.ImageFillOriginal

	centered := container.NewCenter(img)
	counter:=widget.NewButton("Count",func(){
		count+=1
		setter(count,value)
	})
	reset:=widget.NewButton("Reset",func(){
		count=0;
		setter(count,value)
	})
	buttons:=container.New(layout.NewHBoxLayout(),counter,layout.NewSpacer(),reset)
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered,buttons))
	myWindow.Resize(fyne.NewSize(100,100))
	myWindow.ShowAndRun()
}
