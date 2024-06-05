package entity

type Catalog struct {
	ID int `json:"id"`
	//UserID    int       `json:"user_id"`
	//CarID     int       `json:"car_id"`
	Model string `json:"model"`
	Brand string `json:"brand"`
	Color string `json:"color"`
	Price int    `json:"price"`
	//StartDate time.Time `json:"start_date"`
	//EndDate   time.Time `json:"end_date"`
	//Status    string    `json:"status"`
}
