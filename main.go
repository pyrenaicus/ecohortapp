package main

import (
	"database/sql"
	"ecohortapp/repository"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/widget"
	_ "github.com/glebarez/go-sqlite"
)

// definim el tipus Config
type Config struct {
	App                      fyne.App    // guardat de l'App, el canvas
	InfoLog                  *log.Logger // log d'execució
	ErrorLog                 *log.Logger // log d'errors
	MainWindow               fyne.Window // Pantall principal de l'App
	ClimaDadesContainer      *fyne.Container
	PronosticGraficContainer *fyne.Container       // Grafi de la 1a pestanya
	DB                       repository.Repository // Punter de la DB
	HTTPClient               http.Client
}

// declarem myApp de tipus Config
var myApp Config

func main() {
	// Crear app amb fyne
	// id de l'app amb domini invers
	fyneApp := app.NewWithID("cat.cibernarium.ecohortapp")
	// assignem l'app a la clau App de myApp
	myApp.App = fyneApp

	// definirem els nostres logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile)

	// Conexió a la BBDD
	sqlDb, err := myApp.connectSQL()
	if err != nil {
		log.Panic(err)
	}
	// Repositori de la BBDD
	myApp.setupDB(sqlDb)

	// Configuració de la pantalla
	myApp.MainWindow = fyneApp.NewWindow("Eco Hort App")
	myApp.MainWindow.Resize(fyne.NewSize(800, 700)) // Definim tamany finestra pral
	myApp.MainWindow.SetFixedSize(true)             // Fixem el tamany
	myApp.MainWindow.SetMaster()                    // PIndiquem que es pantalla pral

	myApp.makeUI()
	// Mostrar i executar l'app
	myApp.MainWindow.ShowAndRun()
}

func (app *Config) connectSQL() (*sql.DB, error) {
	path := ""
	// Consultar la variable d'entorn
	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		// Creem sqlite si no existeix
		path = app.App.Storage().RootURI().Path() + "/sql.db"

		// Mostra la ruta al log
		app.InfoLog.Println("la bd esta a:", path)
	}

	// Crear la conexió
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// funció per configurar la bd amb el migrate
func (app *Config) setupDB(sqlDB *sql.DB) {
	repository.NewSQLiteRepository(sqlDB)
	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
}
