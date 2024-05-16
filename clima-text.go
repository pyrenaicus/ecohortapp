package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (app *Config) getClimaText() (*canvas.Text, *canvas.Text, *canvas.Text, *canvas.Text) {

	var parte Diaria
	var precipitacio, tempMax, tempMin, humitat *canvas.Text

	prediccio, err := parte.GetPrediccions()

	// filtrar si la petició de l'api és correcte
	if err != nil {
		// codi per quan no funcioni
		gris := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
		precipitacio = canvas.NewText("Precipitació: No definit", gris)
		tempMax = canvas.NewText("Temp. Max.: No definit", gris)
		tempMin = canvas.NewText("Temp. Min.: No definit", gris)
		humitat = canvas.NewText("Humitat: No definit", gris)
	} else {
		// coid pier quan funcioni
		colorin := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		// Filtre si es menor a 50
		if prediccio.ProbPrecipitacio < 50 {
			colorin := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
		}

		// preparem els texts
		precipitacioTxt := fmt.Sprintf("Precipitació: %d%%", prediccio.ProbPrecipitacio)
		tempMaxTxt := fmt.Sprintf("temp. Max.: %dº", prediccio.TemperaturaMax)
		tempMinTxt := fmt.Sprintf("temp. Min.: %dº", prediccio.TemperaturaMin)
		humitatTxt := fmt.Sprintf("Humitat: %d%%", prediccio.HumitatRelativa)

		precipitacio = canvas.NewText(precipitacioTxt, colorin)
		tempMax = canvas.NewText(tempMaxTxt, nil)
		tempMin = canvas.NewText(tempMinTxt, nil)
		humitat = canvas.NewText(humitatTxt, colorin)
	}

	precipitacio.Alignment = fyne.TextAlignLeading
	tempMax.Alignment = fyne.TextAlignCenter
	tempMin.Alignment = fyne.TextAlignCenter
	humitat.Alignment = fyne.TextAlignTrailing

	return precipitacio, tempMax, tempMin, humitat
}
