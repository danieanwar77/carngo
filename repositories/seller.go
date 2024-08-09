package repository

import (
	"database/sql"
	"carngo/structs"
)

func GetAllSeller(db *sql.DB) (results []structs.Seller, err error) {
	sql := "SELECT * FROM seller"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var seller structs.Seller

		err = rows.Scan(&seller.ID, &seller.Email, &seller.Name, &seller.PhoneNumber, &seller.UserName, &seller.CreatedAt, &seller.CreatedBy, &seller.ModifiedAt, &seller.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, seller)
	}
	return
}

func GetSeller(db *sql.DB, seller structs.Seller) (results []structs.Seller, err error) {
	sql := "SELECT * FROM seller WHERE id = $1"

	rows, err := db.Query(sql, seller.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var seller structs.Seller

		err = rows.Scan(&seller.ID, &seller.Email, &seller.Name, &seller.PhoneNumber, &seller.UserName,  &seller.CreatedAt, &seller.CreatedBy, &seller.ModifiedAt, &seller.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, seller)
	}
	return

}

func InsertUserSeller(db *sql.DB, seller structs.Seller) (err error) {

	sql := `
	INSERT INTO users(username, password, level, token, created_by) VALUES ($1, $2, 3, '', (SELECT MAX(id)+1 FROM users))
	`

	errs := db.QueryRow(sql, &seller.UserName, &seller.UserName)

	return errs.Err()
}

func InsertSeller(db *sql.DB, seller structs.Seller) (err error) {

	sql := `
	INSERT INTO seller(email, name, phone_number, username, created_by) VALUES ($1, $2, $3, $4, (SELECT id FROM users WHERE username = $5))
	`

	errs := db.QueryRow(sql, &seller.Email, &seller.Name, &seller.PhoneNumber, &seller.UserName, &seller.UserName)

	return errs.Err()
}

func UpdateSeller(db *sql.DB, seller structs.Seller) (err error) {
	sql := "UPDATE seller SET email = $1, name = $2, phone_number = $3, modified_by = $4 WHERE id = $5"

	errs := db.QueryRow(sql, &seller.Email, &seller.Name, &seller.PhoneNumber, &seller.ID, &seller.ID)

	return errs.Err()
}

func DeleteSeller(db *sql.DB, seller structs.Seller) (err error) {
	sql := "DELETE FROM seller WHERE id = $1"

	errs := db.QueryRow(sql, seller.ID)
	return errs.Err()
}
