package repository

import (
	"database/sql"
	"carngo/structs"
)

func GetUserAuth(db *sql.DB, user structs.User) (results []structs.User, err error) {
	sql := "SELECT * FROM users WHERE username = $1 AND password = $2"

	rows, err := db.Query(sql, user.UserName, user.Password)
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

func ChangeUserToken(db *sql.DB, user structs.User) (err error) {
	sql := "UPDATE users SET token = $1 WHERE username = $2"

	errs := db.QueryRow(sql, &user.Token, &user.UserName)

	return errs.Err()
}

func GetVerifiyToken(db *sql.DB, user structs.User) (results []structs.User, err error) {
	sql := "SELECT * FROM users WHERE token = $1"

	rows, err := db.Query(sql, user.Token)
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

