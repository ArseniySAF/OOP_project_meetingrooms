package storage

import (
 "database/sql"
 "meeting-rooms/internal/models"

 "github.com/google/uuid"
 "github.com/lib/pq"
)

func (ps *Postgres) CreateRoom(input models.InputRoom) (*models.Room, error) {
 query := `
  INSERT INTO rooms (name, capacity, floor, equipment)
  VALUES ($1, $2, $3, $4)
  RETURNING id, name, capacity, floor, equipment, is_active, created_at, updated_at;
 `

 room := &models.Room{}

 row := ps.db.QueryRow(query, input.Name, input.Capacity, input.Floor, pq.Array(input.Equipment))

 err := row.Scan(
  &room.ID,
  &room.Name,
  &room.Capacity,
  &room.Floor,
  pq.Array(&room.Equipment),
  &room.IsActive,
  &room.CreatedAt,
  &room.UpdatedAt,
 )

 if err == sql.ErrNoRows {
  return nil, sql.ErrNoRows
 } else if err != nil {
  return nil, err
 }
 return room, nil
}

func (ps *Postgres) GetListRooms(filter models.RoomFilter) ([]*models.Room, error) {
 rooms := make([]*models.Room, 0)

 query, args := buildRoomsListQuery(filter)

 rows, err := ps.db.Query(query, args...)
 if err != nil {
  return nil, err
 }
 defer rows.Close()

 for rows.Next() {
  room := &models.Room{}

  if err := rows.Scan(
   &room.ID,
   &room.Name,
   &room.Capacity,
   &room.Floor,
   pq.Array(&room.Equipment),
   &room.IsActive,
   &room.CreatedAt,
   &room.UpdatedAt,
  ); err != nil {
   return nil, err
  }

  rooms = append(rooms, room)
 }

 if err := rows.Err(); err != nil {
  return nil, err
 }

 return rooms, nil
}

func (ps *Postgres) GetRoomById(id uuid.UUID) (*models.Room, error) {
 query := `
  SELECT id, name, capacity, floor, equipment, is_active, created_at, updated_at
  FROM rooms
  WHERE id = $1;
 `

 room := &models.Room{}

 row := ps.db.QueryRow(query, id)

 err := row.Scan(
  &room.ID,
  &room.Name,
  &room.Capacity,
  &room.Floor,
  pq.Array(&room.Equipment),
  &room.IsActive,
  &room.CreatedAt,
  &room.UpdatedAt,
 )

 if err == sql.ErrNoRows {
  return nil, sql.ErrNoRows
 } else if err != nil {
  return nil, err
 }

 return room, nil
}

func (ps *Postgres) DeleteRoomById(id uuid.UUID) error {
 query := `
  DELETE FROM rooms WHERE id = $1;
 `
 _, err := ps.db.Exec(query, id)
 if err != nil {
  return err
 }

 return nil
}
