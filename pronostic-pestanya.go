package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pronosticTab() *fyne.Container {
	image := app.obtenirGrafic()
	contenidorGrafic := container.NewVBox(image)
	app.PronosticGraficContainer = contenidorGrafic
	return contenidorGrafic
}

func (app *Config) obtenirGrafic() *canvas.Image {
	url := "https://my.meteoblue.com/images/meteogram?temperature_units=C&wind_units=kmh&precipitation_units=mm&darkmode=true&iso2=es&lat=41.5168&lon=1.901&asl=111&tz=Europe%2FMadrid&dpi=72&apikey=jhMJTOUVRNvs25m4&lang=en&location_name=Abrera&windspeed_units=kmh&sig=2496a6325c6725ea1e1adc17ac02cde7"

	var img *canvas.Image

	err := app.downloadImage(url, "pronostic.png")
	if err != nil {
		// Quan no
		img = canvas.NewImageFromResource(resourceNodisponiblePng)

	} else {
		// es pugii recuperar
		img = canvas.NewImageFromFile("pronostic.png")
	}

	img.SetMinSize(fyne.Size{
		Width:  770,
		Height: 410,
	})

	// img.FillMode = canvas.ImageFillOriginal

	return img
}

func (app *Config) downloadImage(url string, nomArxiu string) error {
	res, err := app.HTTPClient.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("Hi ha algun problema")
	}

	binari, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	imatge, _, err := image.Decode(bytes.NewReader(binari))
	if err != nil {
		return err
	}
	// creem arxiu buit
	arxiu, err := os.Create(fmt.Sprintf("./%s", nomArxiu))
	if err != nil {
		return err
	}

	// afegim la imatge decodificada en format png
	err = png.Encode(arxiu, imatge)
	if err != nil {
		return err
	}

	return nil
}
