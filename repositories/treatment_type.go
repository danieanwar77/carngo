package repository

import (
	"carngo/structs"
	"database/sql"
)

func GetAllTreatmentType(db *sql.DB) (results []structs.TreatmentType, err error) {
	sql := "SELECT * FROM treatment_type"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var treatment structs.TreatmentType

		err = rows.Scan(&treatment.ID, &treatment.Name, &treatment.Description, &treatment.CreatedAt, &treatment.CreatedBy, &treatment.ModifiedAt, &treatment.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, treatment)
	}
	return
}

func GetTreatmentType(db *sql.DB, treatment structs.TreatmentType) (results []structs.TreatmentType, err error) {
	sql := "SELECT * FROM treatment_type WHERE id = $1"

	rows, err := db.Query(sql, treatment.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var treatment structs.TreatmentType

		err = rows.Scan(&treatment.ID, &treatment.Name, &treatment.Description, &treatment.CreatedAt, &treatment.CreatedBy, &treatment.ModifiedAt, &treatment.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, treatment)
	}
	return

}

func InsertTreatmentType(db *sql.DB, treatment structs.TreatmentType) (err error) {

	sql := "INSERT INTO treatment_type(name, description, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &treatment.Name, &treatment.Description, &treatment.CreatedBy)

	return errs.Err()
}

func UpdateTreatmentType(db *sql.DB, treatment structs.TreatmentType) (err error) {
	sql := "UPDATE treatment_type SET name = $1, description = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, treatment.Name, treatment.Description, treatment.ID, treatment.ID)

	return errs.Err()
}

func DeleteTreatmentType(db *sql.DB, treatment structs.TreatmentType) (err error) {
	sql := "DELETE FROM treatment_type WHERE id = $1"

	errs := db.QueryRow(sql, treatment.ID)
	return errs.Err()
}
