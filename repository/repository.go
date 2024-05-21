package repository

import (
	"errors"
	"time"
)

var (
	errorUpdatejant = errors.New("Ha fallat l'actulaització")
	errorsBorrant   = errors.New("Ha fallat borrant")
)

type Repository interface {
	Migrate() error
	// Crear, Llegir un elment, llegir tots, Update, Borrar
	InsertRegistre(registre Registres) (*Registres, error)
	ObtenirTotsRegistres() ([]Registres, error)
	ObtenirRegistrePerID(id int64) (*Registres, error)
	ActualitzarRegistre(id int64, actualitzar Registres) error
	BorrarRegistre(id int64) error
}

type Registres struct {
	ID           int64     `json:"id"`
	Data         time.Time `json:"data_registre"`
	Precipitacio int       `json:"precipitacio"`
	TempMax      int       `json:"temp_max"`
	TempMin      int       `json:"temp_min"`
	Humitat      int       `json:"humitat"`
}
