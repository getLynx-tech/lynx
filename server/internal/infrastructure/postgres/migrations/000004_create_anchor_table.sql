-- +goose Up
CREATE TABLE anchors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    anchor_id TEXT NOT NULL UNIQUE,
    x DOUBLE PRECISION NOT NULL,
    y DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trigger_anchor_updated_at
    BEFORE UPDATE ON anchors
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS trigger_anchor_updated_at ON anchors;

DROP TABLE IF EXISTS anchors;