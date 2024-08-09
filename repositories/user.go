package repository

import (
	"database/sql"
	"carngo/structs"
)

func GetAllUser(db *sql.DB) (results []structs.User, err error) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var users structs.User

		err = rows.Scan(&users.ID, &users.UserName, &users.Password, &users.Level, &users.Token, &users.CreatedAt, &users.CreatedBy, &users.ModifiedAt, &users.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, users)
	}
	return
}

func GetUser(db *sql.DB, users structs.User) (results []structs.User, err error) {
	sql := "SELECT * FROM users WHERE id = $1"

	rows, err := db.Query(sql, users.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var users structs.User

		err = rows.Scan(&users.ID, &users.UserName, &users.Password, &users.Level, & users.Token, &users.CreatedAt, &users.CreatedBy, &users.ModifiedAt, &users.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, users)
	}
	return

}

func InsertUser(db *sql.DB, users structs.User) (err error) {

	sql := "INSERT INTO users(username, password, level, token, created_at, created_by) VALUES ($1, $2, $3, '', $3, $4)"

	errs := db.QueryRow(sql, &users.UserName, &users.Password, &users.Level, &users.ID)

	return errs.Err()
}

func UpdateUser(db *sql.DB, users structs.User) (err error) {
	sql := "UPDATE users SET username = $1, password = $2, level = $3, modified_by = $4 WHERE id = $5"

	errs := db.QueryRow(sql, users.UserName, users.Password, users.Level, users.ID, users.ID)

	return errs.Err()
}

func DeleteUser(db *sql.DB, users structs.User) (err error) {
	sql := "DELETE FROM users WHERE id = $1"

	errs := db.QueryRow(sql, users.ID)
	return errs.Err()
}
