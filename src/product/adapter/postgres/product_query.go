package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/product/entity"
)

const (
	getProductByID = `
		SELECT
			product_id,
			price,
			name 
		FROM
			"products" 
		WHERE
			product_id = $1;
	`
	getProducts = `
		SELECT
			product_id,
			price,
			name 
		FROM
			"products"; 
	`
)

func (q *Queries) GetProductByID(ctx context.Context, productID int64) (*entity.DBProduct, error) {
	i := &entity.DBProduct{}
	err := q.db.QueryRowContext(ctx, getProductByID, productID).Scan(
		&i.ProductID,
		&i.Price,
		&i.Name,
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (q *Queries) GetProducts(ctx context.Context) (*[]entity.DBProduct, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []entity.DBProduct{}
	for rows.Next() {
		var i entity.DBProduct
		if err := rows.Scan(
			&i.ProductID,
			&i.Price,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return &items, nil
}
