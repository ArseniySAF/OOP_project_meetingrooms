package service

import (
	"meeting-rooms/internal/models"

	"github.com/google/uuid"
)

type Postgres interface {
	CreateRoom(input models.InputRoom) (*models.Room, error)
	GetListRooms(filter models.RoomFilter) ([]*models.Room, error)
	GetRoomById(id uuid.UUID) (*models.Room, error)
	DeleteRoomById(id uuid.UUID) error
	CreateBooking(input models.InputBooking) (*models.Booking, error)
	GetListBookings(filter models.BookingFilter) ([]*models.Booking, error)
	GetMyBookings(user_id uuid.UUID) ([]*models.Booking, error)
	DeleteMyBookingById(id uuid.UUID, user_id uuid.UUID) error
	GetAnalytics() ([]*models.Analytics, error)
}

type MeetingService struct {
	Store Postgres
}

func NewMeetingService(store Postgres) *MeetingService {
	return &MeetingService{Store: store}
}
