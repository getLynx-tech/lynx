-- name: CreateVertex :one
INSERT INTO vertices (
    polygon_id,
    longitude,
    latitude
) VALUES (
    @polygon_id,
    @longitude,
    @latitude
)
RETURNING *;