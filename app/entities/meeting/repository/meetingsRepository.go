package repository

import (
	"cmd/app/entities/meeting"
	"cmd/app/entities/meeting/query"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

//go:generate mockery --name=MeetingsRepository
type MeetingsRepository interface {
	FindByCriteria(ctx context.Context, criteria query.FindCriteria) ([]meeting.Meeting, error)
	FindByID(ctx context.Context, id uuid.UUID) (*meeting.Meeting, error)
	Create(ctx context.Context, meeting *meeting.Meeting) (*meeting.Meeting, error)
	Update(ctx context.Context, meeting *meeting.Meeting) (*meeting.Meeting, error)
	Delete(ctx context.Context, meeting *meeting.Meeting) error
}

type MeetingsDatabaseRepository struct {
	db *sql.DB
}

func NewMeetingsDatabaseRepository(providedConnection *sql.DB) *MeetingsDatabaseRepository {
	return &MeetingsDatabaseRepository{db: providedConnection}
}

func (repository *MeetingsDatabaseRepository) FindByID(ctx context.Context, id uuid.UUID) (*meeting.Meeting, error) {
	var currentMeeting meeting.Meeting
	row := query.FindByID(ctx, id, repository.db)
	if err := row.Scan(&currentMeeting.ID, &currentMeeting.GatheringPlaceId, &currentMeeting.InitiatorsId, &currentMeeting.StartTime, &currentMeeting.EndTime,
		&currentMeeting.UsersQuantity, &currentMeeting.State); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no such meeting: %w", err)
		}
		return nil, fmt.Errorf("cannot query the database: %w", err)
	}
	return &currentMeeting, nil
}

func (repository *MeetingsDatabaseRepository) FindByCriteria(ctx context.Context, criteria query.FindCriteria) ([]meeting.Meeting, error) {
	var meetings []meeting.Meeting
	rows, err := query.FindByCriteria(ctx, criteria, repository.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return nil, fmt.Errorf("no such meeting: %w", err)
		}
		return nil, fmt.Errorf("cannot query the database %w", err)
	}
	var currentMeeting meeting.Meeting
	for rows.Next() {
		if err = rows.Scan(&currentMeeting.ID, &currentMeeting.GatheringPlaceId, &currentMeeting.InitiatorsId,
			&currentMeeting.StartTime, &currentMeeting.EndTime, &currentMeeting.UsersQuantity, &currentMeeting.State); err != nil {
			return nil, fmt.Errorf("cannot query the database %w", err)
		}
		meetings = append(meetings, currentMeeting)
	}
	return meetings, nil
}

func (repository *MeetingsDatabaseRepository) Create(ctx context.Context, meeting *meeting.Meeting) (*meeting.Meeting, error) {
	err := query.Create(ctx, meeting, repository.db)
	if err != nil {
		return meeting, fmt.Errorf("meeting cannot be created: %v", err)
	}
	return meeting, nil
}

func (repository *MeetingsDatabaseRepository) Update(ctx context.Context, meeting *meeting.Meeting) (*meeting.Meeting, error) {
	var err = query.Update(ctx, meeting, repository.db)
	if err != nil {
		return meeting, fmt.Errorf("meeting cannot be updated: %v", err)
	}
	return meeting, nil
}

func (repository *MeetingsDatabaseRepository) Delete(ctx context.Context, meeting *meeting.Meeting) error {
	var err = query.Delete(ctx, meeting, repository.db)
	if err != nil {

		return fmt.Errorf("meeting cannot be deleted: %v", err)
	}
	return nil
}
