package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// creem un canvas
	app := app.New()
	// creem finestra i el seu títol
	window := app.NewWindow("La meva App")
	// creem contingut
	window.SetContent(widget.NewLabel("Hola Classe"))
	window.ShowAndRun()
}
