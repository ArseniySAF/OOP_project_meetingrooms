package service

import (
	"errors"
	"meeting-rooms/internal/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

func validateForCreateRoom(input models.InputRoom) error {
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("Name cannot be empty")
	}

	if input.Capacity <= 0 {
		return errors.New("Сapacity must be greater than or equal to 0")
	}

	if input.Floor < 0 {
		return errors.New("Foor cannot be less than 0")
	}

	return nil
}

func validateListRoomsFilter(filter models.RoomFilter) error {
	if filter.Floor != nil && *filter.Floor < 0 {
		return errors.New("Foor cannot be less than 0")
	}

	if filter.MinCapacity != nil && *filter.MinCapacity < 1 {
		return errors.New("Сapacity must be greater than or equal to 0")
	}

	if filter.Limit != nil && *filter.Limit < 0 {
		return errors.New("Limit must be greater than or equal to 0")
	}

	if filter.Offset != nil && *filter.Offset < 0 {
		return errors.New("Offset must be greater than or equal to 0")
	}

	return nil
}

func validateFroCreateBooking(input models.InputBooking) error {
	if strings.TrimSpace(input.Title) == "" {
		return errors.New("Title cannot be empty")
	}

	if input.AttendeesCount < 1 {
		return errors.New("AttendeesCount must be greater than 0")
	}

	if input.StartTime.IsZero() {
		return errors.New("start time cannot be empty")
	}

	if input.EndTime.IsZero() {
		return errors.New("start time cannot be empty")
	}

	if !input.EndTime.After(input.StartTime) {
		return errors.New("end time must be after start time")
	}

	if input.StartTime.Before(time.Now()) {
		return errors.New("start time cannot be in the past")
	}

	return nil
}

func validateBookingFilter(filter models.BookingFilter) error {
	if filter.RoomId != nil && *filter.RoomId == uuid.Nil {
		return errors.New("room ID cannot be empty")
	}

	if filter.DateFrom != nil && filter.DateFrom.IsZero() {
		return errors.New("date from cannot be empty")
	}

	if filter.DateTo != nil && filter.DateTo.IsZero() {
		return errors.New("date to cannot be empty")
	}

	if filter.DateFrom != nil &&
		filter.DateTo != nil &&
		filter.DateFrom.After(*filter.DateTo) {
		return errors.New("date from cannot be after date to")
	}

	if filter.Limit != nil && *filter.Limit < 0 {
		return errors.New("Limit must be greater than or equal to 0")
	}

	if filter.Offset != nil && *filter.Offset < 0 {
		return errors.New("Offset must be greater than or equal to 0")
	}

	return nil
}

func validateMyBookingFilter(filter models.MyBookingFilter) error {
	if filter.Limit != nil && *filter.Limit < 0 {
		return errors.New("Limit must be greater than or equal to 0")
	}

	if filter.Offset != nil && *filter.Offset < 0 {
		return errors.New("Offset must be greater than or equal to 0")
	}

	return nil
}
