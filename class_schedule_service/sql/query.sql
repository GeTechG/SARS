-- name: CreateClass :exec
INSERT INTO classes
(`date`, `order`, subject, teacher, `group`, class_subject)
VALUES(?, ?, ?, ?, ?, ?);

-- name: ReplaceClass :exec
REPLACE classes
(`date`, `order`, subject, teacher, `group`, class_subject)
VALUES(?, ?, ?, ?, ?, ?);

-- name: GetClasses :many
SELECT id, `date`, `order`, subject, teacher, `group`, class_subject
FROM classes;
