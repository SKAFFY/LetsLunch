package query

import (
	"cmd/app/entities/meeting"
	"context"
	"database/sql"
	"fmt"
)

func Delete(ctx context.Context, meeting *meeting.Meeting, db *sql.DB) error {
	const query = `DELETE FROM meetings WHERE id = $1`
	_, err := db.Exec(query, meeting.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}
