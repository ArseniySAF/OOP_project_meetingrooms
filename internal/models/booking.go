package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID             uuid.UUID
	RoomID         uuid.UUID
	UserID         uuid.UUID
	Title          string
	Description    *string
	StartTime      time.Time
	EndTime        time.Time
	Status         string
	AttendeesCount int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type InputBooking struct {
	RoomId         uuid.UUID
	UserId         uuid.UUID // нет в requestBody
	Title          string
	Description    *string
	StartTime      time.Time
	EndTime        time.Time
	AttendeesCount int
}

type BookingFilter struct {
	RoomId   *uuid.UUID
	DateFrom *time.Time
	DateTo   *time.Time
}
