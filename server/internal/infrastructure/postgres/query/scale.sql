-- name: CreateScale :one
INSERT INTO "scales" (
    meters,
    pixels
) VALUES (
$1,
$2
)
RETURNING *;

-- name: GetScale :one
SELECT *
FROM "scales"
LIMIT 1;