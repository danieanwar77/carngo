package repository

import (
	"carngo/structs"
	"database/sql"
)

func GetAllCasualtyType(db *sql.DB) (results []structs.CasualtyType, err error) {
	sql := "SELECT * FROM casualty_type"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var casualty structs.CasualtyType

		err = rows.Scan(&casualty.ID, &casualty.Name, &casualty.Description, &casualty.Level, &casualty.CreatedAt, &casualty.CreatedBy, &casualty.ModifiedAt, &casualty.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, casualty)
	}
	return
}

func GetCasualtyType(db *sql.DB, casualty structs.CasualtyType) (results []structs.CasualtyType, err error) {
	sql := "SELECT * FROM casualty_type WHERE id = $1"

	rows, err := db.Query(sql, casualty.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var casualty structs.CasualtyType

		err = rows.Scan(&casualty.ID, &casualty.Name, &casualty.Description, &casualty.Level, &casualty.CreatedAt, &casualty.CreatedBy, &casualty.ModifiedAt, &casualty.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, casualty)
	}
	return

}

func InsertCasualtyType(db *sql.DB, casualty structs.CasualtyType) (err error) {

	sql := "INSERT INTO casualty_type(name, description, level, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &casualty.Name, &casualty.Description, &casualty.Level, &casualty.CreatedBy)

	return errs.Err()
}

func UpdateCasualtyType(db *sql.DB, casualty structs.CasualtyType) (err error) {
	sql := "UPDATE casualty_type SET name = $1, description = $2, level = $3, modified_by = $4 WHERE id = $5"

	errs := db.QueryRow(sql, casualty.Name, &casualty.Description, &casualty.Level, casualty.ID, casualty.ID)

	return errs.Err()
}

func DeleteCasualtyType(db *sql.DB, casualty structs.CasualtyType) (err error) {
	sql := "DELETE FROM casualty_type WHERE id = $1"

	errs := db.QueryRow(sql, casualty.ID)
	return errs.Err()
}
