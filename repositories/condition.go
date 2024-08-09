package repository

import (
	"carngo/structs"
	"database/sql"
)

func GetAllCondition(db *sql.DB) (results []structs.Condition, err error) {
	sql := "SELECT * FROM condition"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var condition structs.Condition

		err = rows.Scan(&condition.ID, &condition.Name, &condition.Description, &condition.CreatedAt, &condition.CreatedBy, &condition.ModifiedAt, &condition.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, condition)
	}
	return
}

func GetCondition(db *sql.DB, condition structs.Condition) (results []structs.Condition, err error) {
	sql := "SELECT * FROM condition WHERE id = $1"

	rows, err := db.Query(sql, condition.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var condition structs.Condition

		err = rows.Scan(&condition.ID, &condition.Name, &condition.Description, &condition.CreatedAt, &condition.CreatedBy, &condition.ModifiedAt, &condition.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, condition)
	}
	return

}

func InsertCondition(db *sql.DB, condition structs.Condition) (err error) {

	sql := "INSERT INTO condition(name, description, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &condition.Name, &condition.Description, &condition.CreatedBy)

	return errs.Err()
}

func UpdateCondition(db *sql.DB, condition structs.Condition) (err error) {
	sql := "UPDATE condition SET name = $1, description = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, condition.Name, condition.Description, condition.ID, condition.ID)

	return errs.Err()
}

func DeleteCondition(db *sql.DB, condition structs.Condition) (err error) {
	sql := "DELETE FROM condition WHERE id = $1"

	errs := db.QueryRow(sql, condition.ID)
	return errs.Err()
}
