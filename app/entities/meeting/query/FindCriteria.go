package query

import "github.com/gofrs/uuid/v5"

type FindCriteria struct {
	GatheringPlaceId uuid.NullUUID
	InitiatorID      uuid.NullUUID
}
