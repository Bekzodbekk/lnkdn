-- name: CreatePost :exec
INSERT INTO posts
(
    user_id,
    content
) VALUES (
    $1, $2
);

-- name: UpdatePost :exec
UPDATE posts
SET content = $2
WHERE id = $1 AND deleted_at = 0;

-- name: DeletePost :exec
UPDATE posts
SET deleted_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1 AND deleted_at = 0;

-- name: GetPostById :one
SELECT
    id,
    user_id,
    content
FROM posts
WHERE id = $1 AND deleted_at = 0;

-- name: GetAllPosts :many
SELECT
    id,
    user_id,
    content
FROM posts
WHERE deleted_at = 0;

-- name: GetPostByUserId :many
SELECT
    id,
    user_id,
    content
FROM posts
WHERE user_id = $1 AND deleted_at = 0;


