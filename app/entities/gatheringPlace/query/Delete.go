package query

import (
	"cmd/app/entities/gatheringPlace"
	"context"
	"database/sql"
	"fmt"
)

func Delete(ctx context.Context, place *gatheringPlace.GatheringPlace, db *sql.DB) error {
	const query = `DELETE FROM gathering_places WHERE id = $1`
	_, err := db.Exec(query, place.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}
