package query

import (
	"cmd/app/entities/user"
	"context"
	"database/sql"
	"fmt"
)

func Create(ctx context.Context, user *user.User, db *sql.DB) error {
	const queryToUsers = `INSERT INTO users(id, username, display_name, current_meeting_id, rating, age, gender)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(queryToUsers, user.ID, user.Username, user.DisplayName, user.CurrentMeetingId, user.Rating, user.Age, user.Gender)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}

	const queryToHistory = `INSERT INTO meetings_history(user_id, meeting_id) +
		VALUES ($1, $2)`
	for _, meetingId := range user.MeetingHistory {
		_, err := db.Exec(queryToHistory, user.ID, meetingId)
		if err != nil {
			return fmt.Errorf("database query execution error: %w", err)
		}
	}
	return nil
}
