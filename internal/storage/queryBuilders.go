package storage

import (
	"fmt"
	"meeting-rooms/internal/models"
	"strings"
)

func buildRoomsListQuery(filter models.RoomFilter) (string, []any) {
	query := `
		SELECT id, name, capacity, floor, equipment, is_active, created_at, updated_at
		FROM rooms
	`

	conditions := make([]string, 0, 2)
	args := make([]any, 0, 2)

	if filter.Floor != nil {
		args = append(args, *filter.Floor)

		conditions = append(conditions, fmt.Sprintf("floor = $%d", len(args)))
	}

	if filter.MinCapacity != nil {
		args = append(args, *filter.MinCapacity)

		conditions = append(conditions, fmt.Sprintf("capacity >= $%d", len(args)))
	}

	if filter.Limit != nil {
		args = append(args, *filter.Limit)

		conditions = append(conditions, fmt.Sprintf("LIMIT $%d", len(args)))
	}

	if filter.Offset != nil {
		args = append(args, *filter.Offset)

		conditions = append(conditions, fmt.Sprintf("OFFSET $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += "\nWHERE " + strings.Join(conditions, "\n  AND ")
	}

	return query, args
}

func buildBookingListQuery(filter models.BookingFilter) (string, []any) {
	query := `
		SELECT
			id,
			room_id,
			user_id,
			title,
			description,
			start_time,
			end_time,
			status,
			attendees_count,
			created_at,
			updated_at
		FROM bookings
	`

	conditions := make([]string, 0, 3)
	args := make([]any, 0, 3)

	if filter.RoomId != nil {
		args = append(args, *filter.RoomId)

		conditions = append(
			conditions,
			fmt.Sprintf("room_id = $%d", len(args)),
		)
	}

	if filter.DateFrom != nil {
		args = append(args, *filter.DateFrom)

		conditions = append(
			conditions,
			fmt.Sprintf("end_time > $%d", len(args)),
		)
	}

	if filter.DateTo != nil {
		args = append(args, *filter.DateTo)

		conditions = append(
			conditions,
			fmt.Sprintf("start_time < $%d", len(args)),
		)
	}

	if filter.Limit != nil {
		args = append(args, *filter.Limit)

		conditions = append(conditions, fmt.Sprintf("LIMIT $%d", len(args)))
	}

	if filter.Offset != nil {
		args = append(args, *filter.Offset)

		conditions = append(conditions, fmt.Sprintf("OFFSET $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += "\nWHERE " + strings.Join(conditions, "\n  AND ")
	}

	query += "\nORDER BY start_time ASC"

	return query, args
}

func buildMyBookingListQuery(filter models.MyBookingFilter) (string, []any) {
	query := `
		SELECT id, room_id, user_id, title, description, start_time, end_time, status, attendees_count, created_at, updated_at
		FROM bookings
		WHERE user_id = $1;
	`

	conditions := make([]string, 0, 2)
	args := make([]any, 0, 2)

	if filter.Limit != nil {
		args = append(args, *filter.Limit)

		conditions = append(conditions, fmt.Sprintf("LIMIT $%d", len(args)))
	}

	if filter.Offset != nil {
		args = append(args, *filter.Offset)

		conditions = append(conditions, fmt.Sprintf("OFFSET $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += "\nWHERE " + strings.Join(conditions, "\n  AND ")
	}

	return query, args
}
