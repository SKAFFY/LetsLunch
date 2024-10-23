package query

import (
	"cmd/app/entities/user"
	"context"
	"database/sql"
	"fmt"
)

func Delete(ctx context.Context, user *user.User, db *sql.DB) error {
	const queryToUsers = `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(queryToUsers, user.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}

	const queryToHistory = `DELETE FROM meetings_history WHERE user_id = $1`
	_, err = db.Exec(queryToHistory, user.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}
