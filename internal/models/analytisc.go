package models

import (
	"time"

	"github.com/google/uuid"
)

type Analytics struct {
	RoomId        uuid.UUID
	RoomName      string
	TotalBookings int
	TotalHours    float64
	BusiestDay    *time.Time
}
