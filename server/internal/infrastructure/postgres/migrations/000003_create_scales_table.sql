-- +goose Up
CREATE TABLE scales (
    id UUID PRIMARY KEY default gen_random_uuid(),
    meters DOUBLE PRECISION NOT NULL,
    pixels DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trigger_scale_updated_at
    BEFORE UPDATE ON scales
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS trigger_scale_updated_at ON scales;

DROP TABLE IF EXISTS scales;