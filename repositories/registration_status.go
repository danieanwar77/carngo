package repository

import (
	"carngo/structs"
	"database/sql"
)

func GetAllRegistrationStatus(db *sql.DB) (results []structs.RegistrationStatus, err error) {
	sql := "SELECT * FROM registration_status"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var register structs.RegistrationStatus

		err = rows.Scan(&register.ID, &register.Name, &register.Description, &register.CreatedAt, &register.CreatedBy, &register.ModifiedAt, &register.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, register)
	}
	return
}

func GetRegistrationStatus(db *sql.DB, register structs.RegistrationStatus) (results []structs.RegistrationStatus, err error) {
	sql := "SELECT * FROM registration_status WHERE id = $1"

	rows, err := db.Query(sql, register.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var register structs.RegistrationStatus

		err = rows.Scan(&register.ID, &register.Name, &register.Description, &register.CreatedAt, &register.CreatedBy, &register.ModifiedAt, &register.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, register)
	}
	return

}

func InsertRegistrationStatus(db *sql.DB, register structs.RegistrationStatus) (err error) {

	sql := "INSERT INTO registration_status(name, description, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &register.Name, &register.Description, &register.CreatedBy)

	return errs.Err()
}

func UpdateRegistrationStatus(db *sql.DB, register structs.RegistrationStatus) (err error) {

	var modified = sql.NullInt64(register.ModifiedBy)

	sql := "UPDATE registration_status SET name = $1, description = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, register.Name, register.Description, modified, register.ID)

	return errs.Err()
}

func DeleteRegistrationStatus(db *sql.DB, register structs.RegistrationStatus) (err error) {
	sql := "DELETE FROM registration_status WHERE id = $1"

	errs := db.QueryRow(sql, register.ID)
	return errs.Err()
}
