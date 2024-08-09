package repository

import (
	"carngo/structs"
	"database/sql"
)

func GetAllBrand(db *sql.DB) (results []structs.Brand, err error) {
	sql := "SELECT * FROM brand"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var brand structs.Brand

		err = rows.Scan(&brand.ID, &brand.Name, &brand.CreatedAt, &brand.CreatedBy, &brand.ModifiedAt, &brand.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, brand)
	}
	return
}

func GetBrand(db *sql.DB, brand structs.Brand) (results []structs.Brand, err error) {
	sql := "SELECT * FROM brand WHERE id = $1"

	rows, err := db.Query(sql, brand.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var brand structs.Brand

		err = rows.Scan(&brand.ID, &brand.Name, &brand.CreatedAt, &brand.CreatedBy, &brand.ModifiedAt, &brand.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, brand)
	}
	return

}

func InsertBrand(db *sql.DB, brand structs.Brand) (err error) {

	sql := "INSERT INTO brand(name, description, created_by) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &brand.Name, &brand.CreatedBy)

	return errs.Err()
}

func UpdateBrand(db *sql.DB, brand structs.Brand) (err error) {
	sql := "UPDATE brand SET name = $1, description = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, brand.Name, brand.ID, brand.ID)

	return errs.Err()
}

func DeleteBrand(db *sql.DB, brand structs.Brand) (err error) {
	sql := "DELETE FROM brand WHERE id = $1"

	errs := db.QueryRow(sql, brand.ID)
	return errs.Err()
}
