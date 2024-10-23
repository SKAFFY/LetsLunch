package query

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid/v5"
)

func FindByID(ctx context.Context, id uuid.UUID, db *sql.DB) *sql.Row {
	const query = `SELECT * FROM meetings WHERE id = $1`
	row := db.QueryRow(query, id.String())
	return row
}
