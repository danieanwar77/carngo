package repository

import (
	"carngo/structs"
	"database/sql"

	"github.com/lib/pq"
)

func GetAllCar(db *sql.DB) (results []structs.Car, err error) {
	sql := "SELECT * FROM car"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var car structs.Car

		err = rows.Scan(&car.ID, &car.CarType, &car.Brand, &car.FlatNumber, &car.Transmission, &car.Colour, &car.Year, &car.Condition, &car.RegistrationStatus, &car.Seller, pq.Array(&car.Casualty), pq.Array(&car.Treatment), &car.Status, &car.CreatedAt, &car.CreatedBy, &car.ModifiedAt, &car.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, car)
	}
	return
}

func GetCar(db *sql.DB, car structs.Car) (results []structs.Car, err error) {
	sql := "SELECT * FROM car WHERE id = $1"

	rows, err := db.Query(sql, car.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var car structs.Car

		err = rows.Scan(&car.ID, &car.CarType, &car.Brand, &car.FlatNumber, &car.Transmission, &car.Colour, &car.Year, &car.Condition, &car.RegistrationStatus, &car.Seller, pq.Array(&car.Casualty), pq.Array(&car.Treatment), &car.Status, &car.CreatedAt, &car.CreatedBy, &car.ModifiedAt, &car.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, car)
	}
	return

}

func InsertCar(db *sql.DB, car structs.Car) (err error) {

	sql := "INSERT INTO car(car_type, brand, flat_number, transmission, colour, year, condition, registration_status, seller, casualty, treatment, status, created_by) VALUES ($1, $2, $3, '', $4, $5)"

	errs := db.QueryRow(sql, &car.CarType, &car.Brand, &car.FlatNumber, &car.Transmission, &car.Colour, &car.Year, &car.Condition, &car.RegistrationStatus, &car.Seller, pq.Array(&car.Casualty), pq.Array(&car.Treatment), &car.Status, &car.Seller)

	return errs.Err()
}

func UpdateCar(db *sql.DB, car structs.Car) (err error) {
	sql := "UPDATE car SET car_type = $1, brand = $2, flat_number = $3, transmission = $4, colour = $5, year = $6, condition = $7, registration_status = $8, seller = $9, casualty = $10, treatment = $11, status = $12, modified_by = $13 WHERE id = $6"

	errs := db.QueryRow(sql, &car.CarType, &car.Brand, &car.FlatNumber, &car.Transmission, &car.Colour, &car.Year, &car.Condition, &car.RegistrationStatus, &car.Seller, pq.Array(&car.Casualty), pq.Array(&car.Treatment), &car.Status, &car.ModifiedBy, &car.ID)

	return errs.Err()
}

func DeleteCar(db *sql.DB, car structs.Car) (err error) {
	sql := "DELETE FROM car WHERE id = $1"

	errs := db.QueryRow(sql, car.ID)
	return errs.Err()
}
