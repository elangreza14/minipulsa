package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/authentication/entity"
)

const (
	createUser = `
		INSERT INTO users (
  			email,
 			password,
 			token
		) VALUES (
  			$1, $2, $3
		) RETURNING user_id
	`

	updateUser = `
		UPDATE users SET 
  			email = $1,
 			password = $2,
 			token = $3
		WHERE
			user_id = $4
	`

	getUser = `
		SELECT
			user_id,
			email,
			password, 
			token 
		FROM
			users 
		WHERE
			user_id = $1
	`

	getByEmail = `
		SELECT
			user_id,
			email,
			password, 
			token 
		FROM
			users 
		WHERE
			email = $1
	`
)

// InsertUser query for inseting one records
func (q *Queries) InsertUser(ctx context.Context, req entity.ReqPostPutUser) (*int64, error) {
	var id int64

	if err := q.db.QueryRowContext(ctx, createUser, req.Email, req.Password, req.Token).Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (q *Queries) UpdateUser(ctx context.Context, req entity.ReqPostPutUser, userID int64) error {
	if _, err := q.db.ExecContext(ctx, updateUser, req.Email, req.Password, req.Token, userID); err != nil {
		return err
	}

	return nil
}

func (q *Queries) GetUser(ctx context.Context, userID int64) (*entity.DBUser, error) {
	var res entity.DBUser
	if err := q.db.QueryRowContext(ctx, getUser, userID).Scan(&res.UserID, &res.Email, &res.Password, &res.Token); err != nil {
		return nil, err
	}
	return &res, nil
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*entity.DBUser, error) {
	var res entity.DBUser
	if err := q.db.QueryRowContext(ctx, getByEmail, email).Scan(&res.UserID, &res.Email, &res.Password, &res.Token); err != nil {
		return nil, err
	}
	return &res, nil
}
