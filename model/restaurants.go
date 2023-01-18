package model

type Restaurants struct {
	ID          int    `json: "id"`
	Name        string `json: "name"`
	PhoneNumber string `json:"phoneNumber"`
	OrderID     int    `json: "orderID`
	CreatedAt   string `json: "created_at"`
	UpdatedAt   string `json: "updated_at"`
}
