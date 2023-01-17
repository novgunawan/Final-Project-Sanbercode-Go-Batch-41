package model

type Orders struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"cust_id"`
	RestoID    int     `json:"resto_id"`
	Status     string  `json: "status"`
	Food       []Foods `json: "food"`
	CreatedAt  string  `json: "created_at"`
	UpdatedAt  string  `json: "updated_at"`
}
