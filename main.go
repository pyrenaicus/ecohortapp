package main

import (
	"log"

	"fyne.io/fyne/v2"
)

type Config struct {
	App      fyne.App    // guardat de l'App, el canvas
	InfoLog  *log.Logger // log d'execució
	ErrorLog *log.Logger // log d'errors
}

var myApp Config

func main() {
	// Crear app amb fyne

	// definirem els nostres logs

	// Conexió a la BBDD

	// Repositori de la BBDD

	// Configuració de la pantalla

	// Mostrar i executar l'app
}
