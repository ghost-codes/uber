// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: driver.sql

package db

import (
	"context"
)

const createDriver = `-- name: CreateDriver :one
INSERT INTO "driver"(
  name,
  contact,
  email,
  hashed_password,
  car_number,
  car_brand,
  car_color
) VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING id, name, email, hashed_password, contact, car_number, car_brand, car_color, profile_picture
`

type CreateDriverParams struct {
	Name           string `json:"name"`
	Contact        string `json:"contact"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CarNumber      string `json:"car_number"`
	CarBrand       string `json:"car_brand"`
	CarColor       string `json:"car_color"`
}

func (q *Queries) CreateDriver(ctx context.Context, arg CreateDriverParams) (Driver, error) {
	row := q.db.QueryRowContext(ctx, createDriver,
		arg.Name,
		arg.Contact,
		arg.Email,
		arg.HashedPassword,
		arg.CarNumber,
		arg.CarBrand,
		arg.CarColor,
	)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Contact,
		&i.CarNumber,
		&i.CarBrand,
		&i.CarColor,
		&i.ProfilePicture,
	)
	return i, err
}