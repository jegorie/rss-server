-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, owner_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE owner_id=$1;


