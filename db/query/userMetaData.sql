-- name: CreateUserMetaData :one
INSERT INTO "userMetaData" (
    id,
    phone_number,
    date_of_birth
) VALUES ( $1, $2,$3 )
RETURNING *;


-- name: FetchUserMetaDataByID :one
SELECT * FROM "userMetaData" 
WHERE id = $1
LIMIT 1;
