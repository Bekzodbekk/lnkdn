-- name: CheckLike :one
SELECT EXISTS(
    SELECT 1 FROM likes
    WHERE user_id = $1 AND post_id = $2
);

-- name: RemoveLike :exec
DELETE FROM likes WHERE user_id=$1 AND post_id=$2;

-- name: AddLike :exec
INSERT INTO likes(user_id, post_id) VALUES ($1, $2);