package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID
	Name      string
	Capacity  int
	Floor     int
	Equipment []string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InputRoom struct {
	Name      string
	Capacity  int
	Floor     int
	Equipment []string
}

type RoomFilter struct {
	Floor       *int
	MinCapacity *int
	Limit       *int
	Offset      *int
}
