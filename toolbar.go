package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) getToolbar(window fyne.Window) *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)
	return app.getToolbar()
}
