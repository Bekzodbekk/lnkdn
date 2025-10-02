-- name: CreateCourse :exec
INSERT INTO courses
(
    user_id,
    title,
    description,
    category,
    level,
    thumbnail_url
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: UpdateCourse :exec
UPDATE courses
SET
    title = $2,
    description = $3,
    category = $4,
    level = $5,
    thumbnail_url = $6
WHERE id = $1 and deleted_at = 0;

-- name: DeleteCourse :exec
UPDATE courses
SET deleted_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: GetCourses :many
SELECT
    id,
    user_id,
    title,
    description,
    category,
    level,
    thumbnail_url
FROM courses WHERE deleted_at = 0;

-- name: GetCourseById :one
SELECT
    id,
    user_id,
    title,
    description,
    category,
    level,
    thumbnail_url
FROM courses WHERE id = $1 and deleted_at = 0;


-- name: GetCourseByUserId :many
SELECT
    id,
    user_id,
    title,
    description,
    category,
    level,
    thumbnail_url
FROM courses WHERE user_id = $1 and deleted_at = 0;


-- name: CreateLesson :exec
INSERT INTO lessons(
    course_id,
    title,
    video_url,
    duration,
    order_number
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: UpdateLesson :exec
UPDATE lessons
SET 
    title = $2,
    video_url = $3,
    duration = $4,
    order_number = $5
WHERE id = $1 AND deleted_at = 0;

-- name: DeleteLesson :exec
UPDATE lessons
SET
    deleted_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: GetLessonById :one
SELECT
    id,
    course_id,
    title,
    video_url,
    duration,
    order_number
FROM lessons
WHERE id = $1 AND deleted_at = 0;

-- name: GetLessonByCourseId :many
SELECT
    id,
    course_id,
    title,
    video_url,
    duration,
    order_number
FROM lessons
WHERE course_id = $1 AND deleted_at = 0;
