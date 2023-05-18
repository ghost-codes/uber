-- name: CreateDriver :one
INSERT INTO "driver"(
  name,
  contact,
  email,
  hashed_password,
  car_number,
  car_brand,
  car_color
) VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;