package storage

import (
	"errors"
	"meeting-rooms/internal/models"

	"github.com/google/uuid"
)

func (ps *Postgres) CreateBooking(input models.InputBooking) (*models.Booking, error) {
	tx, errTx := ps.db.Begin()
	if errTx != nil {
		return nil, errTx
	}
	defer tx.Rollback()

	queryForUpdate := `
		SELECT is_active FROM rooms WHERE id = $1 FOR UPDATE;
	`

	var isActive bool

	rowForUpdate := tx.QueryRow(queryForUpdate, input.RoomId)

	if err := rowForUpdate.Scan(&isActive); err != nil {
		return nil, err
	}

	if !isActive {
		return nil, errors.New("Room is not active for booking")
	}

	queryForConflict := `
		SELECT EXISTS(
			SELECT 1
			FROM bookings
			WHERE room_id = $1
			  AND status = 'confirmed'
			  AND start_time < $3
			  AND end_time > $2
		);
	`

	var hasConflict bool

	rowConflict := tx.QueryRow(queryForConflict, input.RoomId, input.StartTime, input.EndTime)

	if err := rowConflict.Scan(&hasConflict); err != nil {
		return nil, err
	}

	if hasConflict {
		return nil, errors.New("Conflict")
	}

	queryInsertBookings := `
		INSERT INTO bookings (room_id, user_id, title, start_time, end_time)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, room_id, user_id, title, start_time, end_time, status, attendees_count, created_at, updated_at;
	`

	queryInsertBookingH := `
		INSERT INTO booking_history (booking_id, action, new_status, changed_by)
		VALUES ($1, 'created', 'confirmed', $2);
	`

	booking := &models.Booking{}

	row := tx.QueryRow(queryInsertBookings, input.RoomId, input.UserId, input.Title, input.StartTime, input.EndTime)
	if err := row.Scan(
		&booking.ID,
		&booking.RoomID,
		&booking.UserID,
		&booking.Title,
		&booking.StartTime,
		&booking.EndTime,
		&booking.Status,
		&booking.AttendeesCount,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	); err != nil {
		return nil, err
	}

	if _, err := tx.Exec(queryInsertBookingH, booking.ID, booking.UserID); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return booking, nil
}

func (ps *Postgres) GetListBookings(filter models.BookingFilter) ([]*models.Booking, error) {
	bookings := make([]*models.Booking, 0)

	query, args := buildBookingListQuery(filter)

	rows, err := ps.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		booking := &models.Booking{}

		if err := rows.Scan(
			&booking.ID,
			&booking.RoomID,
			&booking.UserID,
			&booking.Title,
			&booking.Description,
			&booking.StartTime,
			&booking.EndTime,
			&booking.Status,
			&booking.AttendeesCount,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		); err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (ps *Postgres) GetMyBookings(user_id uuid.UUID, filter models.MyBookingFilter) ([]*models.Booking, error) {
	myBookings := make([]*models.Booking, 0)

	query, args := buildMyBookingListQuery(filter)
	rows, err := ps.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		booking := &models.Booking{}

		if err := rows.Scan(
			&booking.ID,
			&booking.RoomID,
			&booking.UserID,
			&booking.Title,
			&booking.Description,
			&booking.StartTime,
			&booking.EndTime,
			&booking.Status,
			&booking.AttendeesCount,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		); err != nil {
			return nil, err
		}

		myBookings = append(myBookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return myBookings, nil
}

func (ps *Postgres) DeleteMyBookingById(id uuid.UUID, user_id uuid.UUID) error {
	tx, errTx := ps.db.Begin()
	if errTx != nil {
		return errTx
	}
	defer tx.Rollback()

	queryForUpdate := `
		UPDATE bookings SET status = 'cancelled', updated_at = NOW()
		WHERE id = $1 AND user_id = $2 AND status = 'confirmed';
	`

	if _, err := tx.Exec(queryForUpdate, id, user_id); err != nil {
		return err
	}

	queryInsertBookingH := `
		INSERT INTO booking_history (booking_id, action, old_status, new_status, changed_by, changed_at)
		VALUES ($1, 'changed_status', 'confirmed', 'canceled', $2, NOW());
	`

	if _, err := tx.Exec(queryInsertBookingH, id, user_id); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
