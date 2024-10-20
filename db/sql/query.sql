-- name: GetUserByID :one
SELECT user_id, mail, name, code, hashed_password FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT user_id, mail, name, code, hashed_password FROM users
WHERE mail = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  user_id, mail, name, code, hashed_password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set mail = $2,
  name = $3,
  hashed_password = $4
WHERE user_id = $1;

-- name: UpdatePassword :exec
UPDATE users
  set hashed_password = $2
WHERE user_id = $1;

-- name: UpdateCode :exec
UPDATE users
  set code = $2
WHERE user_id = $1;

-- name: CreateChat :one
INSERT INTO chats (
  chat_id, user_id_1, user_id_2, content
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetChatByUserID :many
SELECT chat_id, user_id_1, user_id_2, content FROM chats
WHERE user_id_1 = $1 OR user_id_2 = $1;



