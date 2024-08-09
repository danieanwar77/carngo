package repository

import (
	"carngo/structs"
	"database/sql"
)

func GetAllStatus(db *sql.DB) (results []structs.Status, err error) {
	sql := "SELECT * FROM status"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var status structs.Status

		err = rows.Scan(&status.ID, &status.Name, &status.Description, &status.CreatedAt, &status.CreatedBy, &status.ModifiedAt, &status.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, status)
	}
	return
}

func GetStatus(db *sql.DB, status structs.Status) (results []structs.Status, err error) {
	sql := "SELECT * FROM status WHERE id = $1"

	rows, err := db.Query(sql, status.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var status structs.Status

		err = rows.Scan(&status.ID, &status.Name, &status.Description, &status.CreatedAt, &status.CreatedBy, &status.ModifiedAt, &status.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, status)
	}
	return

}

func InsertStatus(db *sql.DB, status structs.Status) (err error) {

	sql := "INSERT INTO status(name, description, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &status.Name, &status.Description, &status.CreatedBy)

	return errs.Err()
}

func UpdateStatus(db *sql.DB, status structs.Status) (err error) {
	sql := "UPDATE status SET name = $1, description = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, status.Name, status.Description, status.ID, status.ID)

	return errs.Err()
}

func DeleteStatus(db *sql.DB, status structs.Status) (err error) {
	sql := "DELETE FROM status WHERE id = $1"

	errs := db.QueryRow(sql, status.ID)
	return errs.Err()
}
