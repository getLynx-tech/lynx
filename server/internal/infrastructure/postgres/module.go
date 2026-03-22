package postgres

import (
	"database/sql"
	"github.com/getLynx-tech/lynx/internal/infrastructure/postgres/sqlc"
	"go.uber.org/fx"
)

func NewQueries(db *sql.DB) *sqlc.Queries {
	return sqlc.New(db)
}

var Module = fx.Module(
	"postgres",
	fx.Provide(
		Connect,
		NewQueries,
	),
)
