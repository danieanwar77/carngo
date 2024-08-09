package repository

import (
	"carngo/structs"
	"database/sql"

)

func GetAllCasualty(db *sql.DB) (results []structs.Casualty, err error) {
	sql := "SELECT * FROM casualty"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var casualty structs.Casualty

		err = rows.Scan(&casualty.ID, &casualty.Car, &casualty.Type, &casualty.Date, &casualty.CreatedAt, &casualty.CreatedBy, &casualty.ModifiedAt, &casualty.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, casualty)
	}
	return
}

func GetCasualty(db *sql.DB, casualty structs.Casualty) (results []structs.Casualty, err error) {
	sql := "SELECT * FROM casualty WHERE id = $1"

	rows, err := db.Query(sql, casualty.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var casualty structs.Casualty

		err = rows.Scan(&casualty.ID, &casualty.Car, &casualty.Type, &casualty.Date, &casualty.CreatedAt, &casualty.CreatedBy, &casualty.ModifiedAt, &casualty.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, casualty)
	}
	return

}

func InsertCasualty(db *sql.DB, casualty structs.Casualty) (err error) {

	sql := "INSERT INTO casualty(car, type, date, created_by) VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sql, &casualty.Car, &casualty.Type, &casualty.Date, &casualty.CreatedBy)

	return errs.Err()
}

func UpdateCasualty(db *sql.DB, casualty structs.Casualty) (err error) {
	sql := "UPDATE casualty SET car = $1, type = $2, date = $3 , modified_by = $4 WHERE id = $5"

	errs := db.QueryRow(sql, casualty.Car, casualty.Type, casualty.Date, casualty.ID, casualty.ID)

	return errs.Err()
}

func DeleteCasualty(db *sql.DB, casualty structs.Casualty) (err error) {
	sql := "DELETE FROM casualty WHERE id = $1"

	errs := db.QueryRow(sql, casualty.ID)
	return errs.Err()
}
