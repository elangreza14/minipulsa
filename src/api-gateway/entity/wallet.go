package entity

type (
	HTTPReqPostUseWallet struct {
		Amount int64 `json:"amount" db:"amount" validate:"required,gt=0"`
	}
)
