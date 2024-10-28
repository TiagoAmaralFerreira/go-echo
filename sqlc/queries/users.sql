-- name: ListUsers :many
SELECT * FROM users;

-- name: ListUser :many
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :many
INSERT INTO users (name, email, password) VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateUser :many
UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?;

-- name: DeleteUser :many
DELETE FROM users WHERE id = ?;