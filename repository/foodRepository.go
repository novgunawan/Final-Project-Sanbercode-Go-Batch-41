package repository

import (
	"database/sql"
	"final-project/model"
	"final-project/service"
)

func GetAllFood(db *sql.DB) (result []model.Foods, err error) {
	sqlStatement := `SELECT * from food`
	//Query function return rows
	rows, err := db.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		var food = model.Foods{}
		err = rows.Scan(&food.ID, &food.Name, &food.Price, &food.CreatedAt, &food.UpdatedAt)
		service.CheckError(err)
		result = append(result, food)
	}
	return
}

func InsertFood(db *sql.DB, food model.Foods) (err error) {
	sqlStatement := `INSERT INTO food(id, name, price, created_at, updated_at) VALUES($1,$2,$3,$4,$5)`
	errs := db.QueryRow(sqlStatement, food.ID, food.Name, food.Price, food.CreatedAt, food.UpdatedAt)
	return errs.Err()
}
