package service

import (
	"meeting-rooms/internal/models"

	"github.com/google/uuid"
)

func (ms *MeetingService) CreateBooking(input models.InputBooking) (*models.Booking, error) {
	if err := validateFroCreateBooking(input); err != nil {
		return nil, err
	}

	return nil, nil
}

func (ms *MeetingService) GetListBookings(filter models.BookingFilter) ([]*models.Booking, error) {
	if err := validateBookingFilter(filter); err != nil {
		return nil, err
	}

	bookings, err := ms.Store.GetListBookings(filter)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (ms *MeetingService) GetMyBookings(user_id uuid.UUID, filter models.MyBookingFilter) ([]*models.Booking, error) {
	if err := validateMyBookingFilter(filter); err != nil {
		return nil, err
	}

	bookings, err := ms.Store.GetMyBookings(user_id, filter)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (ms *MeetingService) DeleteMyBookingById(id uuid.UUID, user_id uuid.UUID) error {
	if err := ms.Store.DeleteMyBookingById(id, user_id); err != nil {
		return err
	}

	return nil
}
