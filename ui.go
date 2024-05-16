package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// Conectem amb API aemet i obtenim dades filtrades
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	//Incloure la info al contenidor
	climaDadesContenidor := container.NewGridWithColumns(4,
		precipitacio,
		tempMax,
		tempMin,
		humitat,
	)

	//Incloure el contenidor a la finestra principal
}
