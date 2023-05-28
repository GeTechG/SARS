-- name: CreateClass :exec
INSERT INTO classes (`date`, `order`, subject, teacher, `group`, class_subject)
VALUES(?, ?, ?, ?, ?, ?);

-- name: ReplaceClass :exec
REPLACE classes (`date`, `order`, subject, teacher, `group`, class_subject)
VALUES(?, ?, ?, ?, ?, ?);

-- name: GetClasses :many
SELECT * FROM classes;

-- name: GetClass :one
SELECT * FROM classes
WHERE id = ? LIMIT 1;

-- name: SetAttendances :exec
REPLACE attendance (class_id, user_uid, value)
VALUES(?, ?, ?);

-- name: GetAttendances :many
SELECT * FROM attendance
WHERE class_id = ?;