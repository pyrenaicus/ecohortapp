package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

// Metode que poblara el struct amb la connexio a la bbdd
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

// Migrate
func (repo *SQLiteRepository) Migrate() error {
	sentencia := `create table if not exists registres(
		id integer primary key autoincrement,
		data_registre integer not null,
		precipitacio integer not null,
		temp_max integer not null,
		temp_min integer not null,
		humitat integer not null)`
	// executar sentencia sql
	_, err := repo.Conn.Exec(sentencia)
	return err
}

// Insertar registre
func (repo *SQLiteRepository) InsertRegistre(registre Registres) (*Registres, error) {
	sentencia := "insert into registres (data_registre, precipitacio, temp_max, temp_min, humitat) values (?,?,?,?,?)"

	res, err := repo.Conn.Exec(sentencia, registre.Data.Unix(), registre.Precipitacio, registre.TempMax, registre.TempMin, registre.Humitat)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	registre.ID = id

	return &registre, nil
}

// Obtenir tots registres
func (repo *SQLiteRepository) ObtenirTotsRegistres() ([]Registres, error) {
	sentencia := "select id, data_registre, precipitacio, temp_max, temp_min, humitat from registres order by data_registre"

	files, err := repo.Conn.Query(sentencia)
	if err != nil {
		return nil, err
	}
	// Tancant la connexio
	defer files.Close()

	// Crear slice de registres
	var conjunt []Registres

	for files.Next() {
		var fila Registres
		var unixTime int64

		err := files.Scan(
			&fila.ID,
			&unixTime,
			&fila.Precipitacio,
			&fila.TempMax,
			&fila.TempMin,
			&fila.Humitat,
		)
		if err != nil {
			return nil, err
		}
		fila.Data = time.Unix(unixTime, 0)
		conjunt = append(conjunt, fila)

	}
	return conjunt, nil

}

// Obtenir registre per Id
func (repo *SQLiteRepository) ObtenirRegistrePerID(id int64) (*Registres, error) {
	consulta := "select id, data_registre, precipitacio, temp_max, temp_min, humitat from registres where id = ?"

	row := repo.Conn.QueryRow(consulta, id)
	var fila Registres
	var unixTime int64

	err := row.Scan(
		&fila.ID,
		&unixTime,
		&fila.Precipitacio,
		&fila.TempMax,
		&fila.TempMin,
		&fila.Humitat,
	)
	if err != nil {
		return nil, err
	}
	fila.Data = time.Unix(unixTime, 0)
	return &fila, nil
}

// Actualitzar registre
func (repo *SQLiteRepository) ActualitzarRegistre(id int64, actualitzar Registres) error {
	if id == 0 {
		return errors.New("La informació és incorrecta")
	}
	sentencia := "update registres set data_registre = ?, precipitacio = ?, temp_max = ?, temp_min = ?, humitat = ? where id = ?"
	res, err := repo.Conn.Exec(sentencia, actualitzar.Data.Unix(), actualitzar.Precipitacio, actualitzar.TempMax, actualitzar.TempMin, actualitzar.Humitat, id)
	if err != nil {
		return err
	}
	filesAfectades, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if filesAfectades == 0 {
		return errorUpdatejant
	}
	return err
}

// Borrar registre
func (repo *SQLiteRepository) BorrarRegistre(id int64) error {
	res, err := repo.Conn.Exec("delete from registres where id = ?", id)
	if err != nil {
		return err
	}
	filesAfectades, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if filesAfectades == 0 {
		return errorsBorrant
	}
	return nil
}

// orm en go
// GorillaMux
// GinGonic
