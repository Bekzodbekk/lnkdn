-- name: CreateComment :exec
INSERT INTO comments
(
    user_id,
    post_id,
    content
) VALUES (
    $1, $2, $3
);

-- name: UpdateComment :exec
UPDATE comments
SET content = $2
WHERE id = $1 AND deleted_at=0;

-- name: DeleteComment :exec
UPDATE comments
SET deleted_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: GetCommentByUserId :many
SELECT 
    id,
    user_id,
    post_id,
    content
FROM comments
WHERE user_id=$1 and deleted_at = 0;

-- name: GetCommentByPostId :many
SELECT 
    id,
    user_id,
    post_id,
    content
FROM comments
WHERE post_id=$1 and deleted_at = 0;


