package repository

import (
	"database/sql"
	"final-project/model"
)

func VerifyCustomer(db *sql.DB, phoneNumber string, customer model.Customers) (errs error) {

	sqlStatement := `SELECT * FROM customer WHERE phone_number = $1`
	errs = db.QueryRow(sqlStatement, phoneNumber).Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.CreatedAt, &customer.UpdatedAt)
	return errs
}

func InsertCustomer(db *sql.DB, customer model.Customers) (err error) {
	sqlStatement := `INSERT INTO customer(id, name, phone_number,created_at, updated_at) VALUES($1,$2,$3,$4,$5)`
	errs := db.QueryRow(sqlStatement, customer.ID, customer.Name, customer.PhoneNumber, customer.CreatedAt, customer.UpdatedAt)

	return errs.Err()
}

func VerifyRestaurant(db *sql.DB, restaurant model.Restaurants) (err error) {
	sqlStatement := `SELECT * FROM restaurant WHERE phone_number = $1`
	errs := db.QueryRow(sqlStatement, restaurant.PhoneNumber)
	return errs.Err()
}

func InsertRestaurant(db *sql.DB, restaurant model.Restaurants) (err error) {
	sqlStatement := `INSERT INTO restaurant(name, phoneNumber,created_at, updated_at) VALUES($1,$2,$3,$4)`
	errs := db.QueryRow(sqlStatement, restaurant.Name, restaurant.PhoneNumber, restaurant.CreatedAt, restaurant.UpdatedAt)

	return errs.Err()
}
