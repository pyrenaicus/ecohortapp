package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

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
	app.ClimaDadesContainer = climaDadesContenidor

	// carrega del toolbar
	barraEines := app.getToolbar(app.MainWindow)

	// pestanyes
	pestanyes := container.NewAppTabs(
		container.NewTabItemWithIcon("Pronòstic", theme.HomeIcon(), canvas.NewText("cascascas", nil)),
		container.NewTabItemWithIcon("Diari Meteorològic", theme.InfoIcon(), canvas.NewText("cascascas", nil)),
	)

	contenidorFinal := container.NewVBox(climaDadesContenidor, barraEines, pestanyes)
	app.MainWindow.SetContent(contenidorFinal)
}
