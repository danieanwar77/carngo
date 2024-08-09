package repository

import (
	"carngo/structs"
	"database/sql"

	"github.com/lib/pq"
)

func GetAllUserLevel(db *sql.DB) (results []structs.UserLevel, err error) {
	sql := "SELECT * FROM user_level"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var users structs.UserLevel

		err = rows.Scan(&users.ID, &users.Name, pq.Array(&users.Access), &users.CreatedAt, &users.CreatedBy, &users.ModifiedAt, &users.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, users)
	}
	return
}

func GetUserLevel(db *sql.DB, users structs.UserLevel) (results []structs.UserLevel, err error) {
	sql := "SELECT * FROM user_level WHERE id = $1"

	rows, err := db.Query(sql, users.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var users structs.UserLevel

		err = rows.Scan(&users.ID, &users.Name, pq.Array(&users.Access), &users.CreatedAt, &users.CreatedBy, &users.ModifiedAt, &users.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, users)
	}
	return

}

func InsertUserLevel(db *sql.DB, users structs.UserLevel) (err error) {

	sql := "INSERT INTO user_level(name, access, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &users.Name, pq.Array(&users.Access), &users.CreatedBy)

	return errs.Err()
}

func UpdateUserLevel(db *sql.DB, users structs.UserLevel) (err error) {
	sql := "UPDATE user_level SET name = $1, access = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, users.Name, pq.Array(&users.Access), users.ID, users.ID)

	return errs.Err()
}

func DeleteUserLevel(db *sql.DB, users structs.UserLevel) (err error) {
	sql := "DELETE FROM user_level WHERE id = $1"

	errs := db.QueryRow(sql, users.ID)
	return errs.Err()
}
