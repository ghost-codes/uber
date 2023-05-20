-- name: CreateCabLocation :one 
INSERT INTO "cabLocation"(
  driver,
  cell_id,
  position,
  available
)VALUES ($1,$2,point($3,$4),$5)
RETURNING *;

-- name: UpdateCabLocation :one
UPDATE "cabLocation" 
SET position = point($2,$3)
WHERE driver = $1
RETURNING *;

-- name: GetCabLocation :one
SELECT * FROM "cabLocation" 
WHERE driver=$1;

-- name: GetCabLocationForSubscription :many
SELECT * FROM "cabLocation"
WHERE cell_id = $1;

-- name: ListCabLocation :many
SELECT * FROM "cabLocation"
WHERE   ST_DWithin(point($1 ,$2)::geography,
              ST_MakePoint(long,lat),8046.72);
