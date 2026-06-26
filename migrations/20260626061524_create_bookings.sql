-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bookings(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'confirmed',
    attendees_count INT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT bookings_time_order_check CHECK (end_time > start_time),
    CONSTRAINT bookings_status_check CHECK (status IN ('confirmed', 'cancelled', 'completed')),
    CONSTRAINT bookings_duration_check CHECK (end_time - start_time <= INTERVAL '8 hours')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bookings;
-- +goose StatementEnd
