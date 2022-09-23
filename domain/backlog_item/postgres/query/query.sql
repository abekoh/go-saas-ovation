-- name: GetOne :one
SELECT *
FROM backlog_items
WHERE id = $1 LIMIT 1;
