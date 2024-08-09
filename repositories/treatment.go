package repository

import (
	"carngo/structs"
	"database/sql"

)

func GetAllTreatment(db *sql.DB) (results []structs.Treatment, err error) {
	sql := "SELECT * FROM treatment"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var treatment structs.Treatment

		err = rows.Scan(&treatment.ID, &treatment.Car, &treatment.Type, &treatment.Date, &treatment.CreatedAt, &treatment.CreatedBy, &treatment.ModifiedAt, &treatment.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, treatment)
	}
	return
}

func GetTreatment(db *sql.DB, treatment structs.Treatment) (results []structs.Treatment, err error) {
	sql := "SELECT * FROM treatment WHERE id = $1"

	rows, err := db.Query(sql, treatment.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var treatment structs.Treatment

		err = rows.Scan(&treatment.ID, &treatment.Car, &treatment.Type, &treatment.Date, &treatment.CreatedAt, &treatment.CreatedBy, &treatment.ModifiedAt, &treatment.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, treatment)
	}
	return

}

func InsertTreatment(db *sql.DB, treatment structs.Treatment) (err error) {

	sql := "INSERT INTO treatment(car, type, date, created_by) VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sql, &treatment.Car, &treatment.Type, &treatment.Date, &treatment.CreatedBy)

	return errs.Err()
}

func UpdateTreatment(db *sql.DB, treatment structs.Treatment) (err error) {
	sql := "UPDATE treatment SET car = $1, type = $2, date = $3 , modified_by = $4 WHERE id = $5"

	errs := db.QueryRow(sql, treatment.Car, treatment.Type, treatment.Date, treatment.ID, treatment.ID)

	return errs.Err()
}

func DeleteTreatment(db *sql.DB, treatment structs.Treatment) (err error) {
	sql := "DELETE FROM treatment WHERE id = $1"

	errs := db.QueryRow(sql, treatment.ID)
	return errs.Err()
}
