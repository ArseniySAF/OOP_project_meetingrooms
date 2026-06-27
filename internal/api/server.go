package api

import (
	"meeting-rooms/internal/models"

	"github.com/google/uuid"
)

type MService interface {
	CreateRoom(input models.InputRoom) (*models.Room, error)
	GetListRooms(filter models.RoomFilter) ([]*models.Room, error)
	GetRoomById(id uuid.UUID) (*models.Room, error)
	DeleteRoomById(id uuid.UUID) error
	CreateBooking(input models.InputBooking) (*models.Booking, error)
	GetListBookings(filter models.BookingFilter) ([]*models.Booking, error)
	GetMyBookings(user_id uuid.UUID, filter models.MyBookingFilter) ([]*models.Booking, error)
	DeleteMyBookingById(id uuid.UUID, user_id uuid.UUID) error
	GetAnalytics() ([]*models.Analytics, error)
}

type HttpHandler struct {
	MeetingService MService
}

func NewServer(service MService) *HttpHandler {
	return &HttpHandler{MeetingService: service}
}

var _ ServerInterface = (*HttpHandler)(nil)
