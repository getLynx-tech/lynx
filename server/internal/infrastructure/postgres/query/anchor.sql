-- name: CreateAnchor :one
INSERT INTO "anchors" (
    anchor_id,
    "x",
    "y"
) VALUES (
$1,
$2,
$3
)
RETURNING *;

-- name: GetAllAnchors :many
SELECT *
FROM "anchors";

-- name: DeleteAllAnchors :exec
DELETE FROM "anchors";