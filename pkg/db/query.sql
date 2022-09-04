-- name: GetEvent :one
SELECT * FROM event
WHERE id = $1 LIMIT 1;

-- name: ListEvents :many
SELECT * FROM event
ORDER BY name;

-- name: CreateEvent :one
INSERT INTO event (
  id, name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM event
WHERE id = $1;