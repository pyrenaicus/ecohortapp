package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App      fyne.App    // guardat de l'App, el canvas
	InfoLog  *log.Logger // log d'execució
	ErrorLog *log.Logger // log d'errors
}

var myApp Config

func main() {
	// Crear app amb fyne
	// id de l'app amb domini invers
	fyneApp := app.NewWithID("cat.cibernarium.ecohortapp")
	myApp.App = fyneApp

	// definirem els nostres logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile)

	// Conexió a la BBDD

	// Repositori de la BBDD

	// Configuració de la pantalla
	window := fyneApp.NewWindow("Eco Hort App")

	// Mostrar i executar l'app
	window.ShowAndRun()
}
