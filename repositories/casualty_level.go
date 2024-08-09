package repository

import (
	"carngo/structs"
	"database/sql"
)

func GetAllCasualtyLevel(db *sql.DB) (results []structs.CasualtyLevel, err error) {
	sql := "SELECT * FROM casualty_level"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var casualty structs.CasualtyLevel

		err = rows.Scan(&casualty.ID, &casualty.Name, &casualty.Description, &casualty.CreatedAt, &casualty.CreatedBy, &casualty.ModifiedAt, &casualty.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, casualty)
	}
	return
}

func GetCasualtyLevel(db *sql.DB, casualty structs.CasualtyLevel) (results []structs.CasualtyLevel, err error) {
	sql := "SELECT * FROM casualty_level WHERE id = $1"

	rows, err := db.Query(sql, casualty.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var casualty structs.CasualtyLevel

		err = rows.Scan(&casualty.ID, &casualty.Name, &casualty.Description, &casualty.CreatedAt, &casualty.CreatedBy, &casualty.ModifiedAt, &casualty.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, casualty)
	}
	return

}

func InsertCasualtyLevel(db *sql.DB, casualty structs.CasualtyLevel) (err error) {

	sql := "INSERT INTO casualty_level(name, description, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &casualty.Name, &casualty.Description, &casualty.CreatedBy)

	return errs.Err()
}

func UpdateCasualtyLevel(db *sql.DB, casualty structs.CasualtyLevel) (err error) {
	sql := "UPDATE casualty_level SET name = $1, description = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, casualty.Name, casualty.Description, casualty.ID, casualty.ID)

	return errs.Err()
}

func DeleteCasualtyLevel(db *sql.DB, casualty structs.CasualtyLevel) (err error) {
	sql := "DELETE FROM casualty_level WHERE id = $1"

	errs := db.QueryRow(sql, casualty.ID)
	return errs.Err()
}
