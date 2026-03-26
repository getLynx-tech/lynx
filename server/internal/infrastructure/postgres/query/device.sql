-- name: UpsertDevice :one
INSERT INTO devices (
    device_id,
    status,
    x,
    y
) VALUES (
    $1,
    $2,
    $3,
    $4
 )
ON CONFLICT (device_id)
    DO UPDATE SET
      status = EXCLUDED.status,
      x = EXCLUDED.x,
      y = EXCLUDED.y
RETURNING *;

-- name: GetAllDevices :many
SELECT * FROM devices;  