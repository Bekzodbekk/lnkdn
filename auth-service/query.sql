-- name: SignIn :one
SELECT id, password FROM users
WHERE username = $1 AND deleted_at = 0;

-- name: GetPasswordWithId :one
SELECT password FROM users
WHERE id = $1;

-- name: CheckEmailExists :one
SELECT EXISTS (
    SELECT 1 FROM users
    WHERE email = $1 AND deleted_at = 0
);

-- name: UpdatePasswordWithEmail :exec
UPDATE users
SET password = $2
WHERE email = $1 and deleted_at = 0; 

-- name: UpdatePasswordWithId :exec
UPDATE users
SET password = $2
WHERE id = $1 and deleted_at = 0; 

-- name: CreateUser :exec
INSERT INTO users
(
    first_name,
    last_name,
    phone,
    email,
    username,
    password
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: UpdateUser :exec
UPDATE users
SET first_name = $2,
    last_name = $3,
    phone = $4,
    email = $5,
    username = $6
WHERE id = $1 AND deleted_at = 0;


-- name: DeleteUser :exec
UPDATE users
SET deleted_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: GetUserById :one
SELECT 
    id,
    first_name,
    last_name,
    phone,
    email,
    username
FROM users
WHERE id = $1 and deleted_at = 0;

-- name: GetUsers :many
SELECT 
    id,
    first_name,
    last_name,
    phone,
    email,
    username
FROM users WHERE deleted_at = 0;

