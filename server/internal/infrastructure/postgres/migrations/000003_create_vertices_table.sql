-- +goose Up
CREATE TABLE vertices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    polygon_id TEXT NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trigger_vertices_updated_at
    BEFORE UPDATE ON vertices
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS trigger_vertices_updated_at ON vertices;

DROP TABLE vertices;