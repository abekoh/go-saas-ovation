-- name: CreateBacklogItem :exec
INSERT INTO backlog_items (id, type, summary, story_point)
VALUES ($1, $2, $3, $4);

-- name: GetOneBacklogItem :one
SELECT *
FROM backlog_items
WHERE id = $1 LIMIT 1;
