/* SELECTION */

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;


/* INSERTION */

-- name: CreateUser :one
INSERT INTO users (
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

/* UPDATE */

-- name: UpdateUser :one
UPDATE users
SET username = $2,
    email = $3,
    password = $4
WHERE id = $1
RETURNING *;

/* DELETION */

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;





