package repository

import (
	"carngo/structs"
	"database/sql"
)

func GetAllTransmission(db *sql.DB) (results []structs.Transmission, err error) {
	sql := "SELECT * FROM transmission"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var transmission structs.Transmission

		err = rows.Scan(&transmission.ID, &transmission.Name, &transmission.Description, &transmission.CreatedAt, &transmission.CreatedBy, &transmission.ModifiedAt, &transmission.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, transmission)
	}
	return
}

func GetTransmission(db *sql.DB, transmission structs.Transmission) (results []structs.Transmission, err error) {
	sql := "SELECT * FROM transmission WHERE id = $1"

	rows, err := db.Query(sql, transmission.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var transmission structs.Transmission

		err = rows.Scan(&transmission.ID, &transmission.Name, &transmission.Description, &transmission.CreatedAt, &transmission.CreatedBy, &transmission.ModifiedAt, &transmission.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, transmission)
	}
	return

}

func InsertTransmission(db *sql.DB, transmission structs.Transmission) (err error) {

	sql := "INSERT INTO transmission(name, description, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &transmission.Name, &transmission.Description, &transmission.CreatedBy)

	return errs.Err()
}

func UpdateTransmission(db *sql.DB, transmission structs.Transmission) (err error) {
	sql := "UPDATE transmission SET name = $1, description = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, transmission.Name, transmission.Description, transmission.ID, transmission.ID)

	return errs.Err()
}

func DeleteTransmission(db *sql.DB, transmission structs.Transmission) (err error) {
	sql := "DELETE FROM transmission WHERE id = $1"

	errs := db.QueryRow(sql, transmission.ID)
	return errs.Err()
}
