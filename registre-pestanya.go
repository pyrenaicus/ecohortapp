package main

import (
	"ecohortapp/repository"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) registresTab() *fyne.Container {
	app.RegistresTable = app.getRegistresTable()
	// registresContenidor := container.NewVBox(app.RegistresTable)
	registresContenidor := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.RegistresTable),
	)
	return registresContenidor

}
func (app *Config) getRegistresTable() *widget.Table {
	data := app.getRegistresSlice()
	app.Registres = data
	taula := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			contenidor := container.NewVBox(widget.NewLabel(""))
			return contenidor
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == (len(data[0])-1) && i.Row != 0 {
				// Situarem el botons d'opcions
				w := widget.NewButtonWithIcon("Borrar", theme.DeleteIcon(), func() {
					// Dialeg de confirmació
					dialog.ShowConfirm("Borrar segur?", "", func(deleted bool) {
						id, _ := strconv.Atoi(data[i.Row][0].(string))
						err := app.DB.BorrarRegistre(int64(id))
						if err != nil {
							app.ErrorLog.Println(err)
						}
						// Actualitzar o refrescar la taula davant de canvis com el borrat
						app.actualitzarRegistresTable()
					}, app.MainWindow)
				})
				// Destaquem el pop-up, similar z-index
				w.Importance = widget.HighImportance
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(data[i.Row][i.Col].(string)),
				}
			}
		})
	// amples de columnes
	colWidths := []float32{50, 200, 200, 200, 200, 200, 110}
	for i := 0; i < len(colWidths); i++ {
		taula.SetColumnWidth(i, colWidths[i])
	}
	return taula
}

func (app *Config) getRegistresSlice() [][]interface{} {
	var slice [][]interface{}
	registres, err := app.DB.ObtenirTotsRegistres()

	if err != nil {
		app.ErrorLog.Println(err)
	}

	//Afegim encapçalaments
	slice = append(slice, []interface{}{"ID", "Data", "Precipitació", "Temp max", "Temp min", "Humitat", "Opcions"})

	for _, v := range registres {
		// Cream interifice buida per fila actual
		var filaActual []interface{}
		// dades
		filaActual = append(filaActual, strconv.FormatInt(v.ID, 10))
		filaActual = append(filaActual, v.Data.Format("2006-01-02"))
		filaActual = append(filaActual, fmt.Sprintf("%d%%", v.Precipitacio))
		filaActual = append(filaActual, fmt.Sprintf("%d", v.TempMax))
		filaActual = append(filaActual, fmt.Sprintf("%d", v.TempMin))
		filaActual = append(filaActual, fmt.Sprintf("%d%%", v.Humitat))
		filaActual = append(filaActual, widget.NewButton("Borrar", func() {}))

		slice = append(slice, filaActual)
	}
	return slice
}
func (app *Config) registresBD() ([]repository.Registres, error) {
	registres, err := app.DB.ObtenirTotsRegistres()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}
	return registres, nil
}
