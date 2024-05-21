package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// Conectem amb API aemet i obtenim dades filtrades
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	//Incloure les dades obtenides al contenidor
	climaDadesContenidor := container.NewGridWithColumns(4,
		precipitacio,
		tempMax,
		tempMin,
		humitat,
	)
	//Incloure el contenidor a la finestra principal, assignant-lo a la clau ClimaDadesContainer
	app.ClimaDadesContainer = climaDadesContenidor

	// carrega del toolbar, cridant la nostra fn
	barraEines := app.getToolbar(app.MainWindow)

	// carrega grafic de la 1a pestanya
	contenidorGraficPestanya := app.pronosticTab()
	// pestanyes
	pestanyes := container.NewAppTabs(
		container.NewTabItemWithIcon("Pronòstic", theme.HomeIcon(), contenidorGraficPestanya),
		container.NewTabItemWithIcon("Diari Meteorològic", theme.InfoIcon(), canvas.NewText("cascascas", nil)),
	)
	// alineacio pestanyes superior
	pestanyes.SetTabLocation(container.TabLocationTop)

	contenidorFinal := container.NewVBox(climaDadesContenidor, barraEines, pestanyes)
	app.MainWindow.SetContent(contenidorFinal)

	// fn anonima controlada per una goroutine
	go func() {
		// execucio cada 30 segons
		for range time.Tick(time.Second * 30) {
			app.actualitzarClimaDadesContent()
		}
	}()
}

func (app *Config) actualitzarClimaDadesContent() {
	app.InfoLog.Println("Actualitzant dades")
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	app.ClimaDadesContainer.Objects = []fyne.CanvasObject{precipitacio, tempMax, tempMin, humitat}
	app.ClimaDadesContainer.Refresh()

	grafic := app.obtenirGrafic()
	app.PronosticGraficContainer.Objects = []fyne.CanvasObject{grafic}
	app.PronosticGraficContainer.Refresh()

}
