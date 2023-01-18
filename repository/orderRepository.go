package repository

import (
	"database/sql"
	"encoding/json"
	"final-project/model"
	"final-project/service"
)

func CreateOrder(db *sql.DB, order model.Orders) (status string, err error) {
	sql := `
	
	INSERT INTO orders(id, cust_id, resto_id, status, food, created_at, updated_at) VALUES ($1,$2,$3,$4,($5)::json,$6,$7)`
	order.Status = "accepted"
	jsonFood, _ := json.Marshal(order.Food)

	errs := db.QueryRow(sql, order.ID, order.CustomerID, order.RestoID, order.Status, jsonFood, order.CreatedAt, order.UpdatedAt)

	return order.Status, errs.Err()
}

func GetActiveOrderResto(db *sql.DB, restoID int, status string, order model.Orders) (result model.Orders, errs error) {
	var foodJSONString string
	sql := `SELECT id, cust_id, resto_id, status, food FROM orders WHERE resto_id = $1 AND status = $2`
	err := db.QueryRow(sql, restoID, status).Scan(&order.ID, &order.CustomerID, &order.RestoID, &order.Status, &foodJSONString)
	json.Unmarshal([]byte(foodJSONString), &order.Food)

	result.ID = order.ID
	result.CustomerID = order.CustomerID
	result.RestoID = order.RestoID
	result.Status = order.Status
	result.Food = order.Food
	return result, err

}

func GetActiveOrderCustomer(db *sql.DB, custID int, status string, order model.Orders) (result model.Orders, errs error) {
	var foodJSONString string
	sql := `SELECT id, cust_id, resto_id, status, food FROM orders WHERE cust_id = $1 AND status = $2`
	err := db.QueryRow(sql, custID, status).Scan(&order.ID, &order.CustomerID, &order.RestoID, &order.Status, &foodJSONString)
	json.Unmarshal([]byte(foodJSONString), &order.Food)

	result.ID = order.ID
	result.CustomerID = order.CustomerID
	result.RestoID = order.RestoID
	result.Status = order.Status
	result.Food = order.Food
	return result, err

}

func FinishOrder(db *sql.DB, order model.Orders, status *string, id int) int64 {
	sql := `UPDATE orders SET status = $1 WHERE id = $2`
	res, errs := db.Exec(sql, *status, id)
	// errs := db.QueryRow(sql, *status, id)
	service.CheckError(errs)

	count, err := res.RowsAffected()
	service.CheckError(err)

	return count

}
