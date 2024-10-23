package query

import (
	"cmd/app/entities/meeting"
	"context"
	"database/sql"
	"fmt"
)

func Update(ctx context.Context, meeting *meeting.Meeting, db *sql.DB) error {
	const query = `UPDATE meetings 
    SET id = $1, gathering_place_id = $2, initiators_id = $3, time_start = $4, 
        time_end = $5, max_participants = $6, state = $7 WHERE id = $1`
	_, err := db.Exec(query, meeting.ID, meeting.GatheringPlaceId, meeting.InitiatorsId,
		meeting.StartTime, meeting.EndTime, meeting.UsersQuantity, meeting.State)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}

/*func Update(ctx context.Context, meeting *meeting.Meeting, db *sql.DB) error {
	const query = `UPDATE meetings SET state = $1 WHERE id = $2`
	_, err := db.Exec(query, meeting.State, meeting.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}*/
