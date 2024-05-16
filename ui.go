package main

func (app *Config) makeUI() {
	// Conectem amb API aemet i obtenim dades filtrades
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	//Incloure la info al contenidor

	//Incloure el contenidor a la finestra principal
}
