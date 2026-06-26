package models

import (
	"time"

	"github.com/google/uuid"
)

type BookingHistory struct {
	ID        int64
	BookingID uuid.UUID
	Action    string
	OldStatus *string
	NewStatus *string
	ChangedBy *uuid.UUID
	ChangedAt time.Time
}
