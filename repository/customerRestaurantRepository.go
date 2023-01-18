package repository

import (
	"database/sql"
	"final-project/model"
)

func VerifyCustomer(db *sql.DB, phoneNumber string, customer model.Customers) (result model.Customers, errs error) {

	sqlStatement := `SELECT id, name, phone_number FROM customer WHERE phone_number = $1`
	errs = db.QueryRow(sqlStatement, phoneNumber).Scan(&customer.ID, &customer.Name, &customer.PhoneNumber)
	result.ID = customer.ID
	result.Name = customer.Name
	result.PhoneNumber = customer.PhoneNumber
	return result, errs
}

func InsertCustomer(db *sql.DB, customer model.Customers) (err error) {
	sqlStatement := `INSERT INTO customer(id, name, phone_number,created_at, updated_at) VALUES($1,$2,$3,$4,$5)`
	errs := db.QueryRow(sqlStatement, customer.ID, customer.Name, customer.PhoneNumber, customer.CreatedAt, customer.UpdatedAt)

	return errs.Err()
}

func VerifyRestaurant(db *sql.DB, phoneNumber string, restaurant model.Restaurants) (result model.Restaurants, errs error) {
	sqlStatement := `SELECT id, name, phone_number FROM restaurant WHERE phone_number = $1`
	errs = db.QueryRow(sqlStatement, phoneNumber).Scan(&restaurant.ID, &restaurant.Name, &restaurant.PhoneNumber)
	result.ID = restaurant.ID
	result.Name = restaurant.Name
	result.PhoneNumber = restaurant.PhoneNumber
	return result, errs
}

func InsertRestaurant(db *sql.DB, restaurant model.Restaurants) (err error) {
	sqlStatement := `INSERT INTO restaurant(id, name, phone_number, created_at, updated_at) VALUES($1,$2,$3,$4,$5)`
	errs := db.QueryRow(sqlStatement, restaurant.ID, restaurant.Name, restaurant.PhoneNumber, restaurant.CreatedAt, restaurant.UpdatedAt)

	return errs.Err()
}
