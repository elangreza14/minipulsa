package entity

type (
	DBProduct struct {
		ProductID int64  `json:"product_id" db:"product_id"`
		Name      string `json:"name" db:"name"`
		Price     int64  `json:"price" db:"price"`
	}

	// ReqPRoductDetail struct {
	// 	ProductID int64 `json:"product_id" db:"product_id"`
	// }
)
