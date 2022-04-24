package entity

type (
	Product struct {
		ProductId int64  `json:"product_id,omitempty"`
		Price     int64  `json:"price,omitempty"`
		Name      string `json:"name,omitempty"`
	}

	HTTPResGetProducts struct {
		Products []Product `json:"products" db:"products" validate:"required"`
	}
)
