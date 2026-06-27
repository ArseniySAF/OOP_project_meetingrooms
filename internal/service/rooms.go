
package service

import (
 "meeting-rooms/internal/models"
 "strings"

 "github.com/google/uuid"
)

func (ms *MeetingService) CreateRoom(input models.InputRoom) (*models.Room, error) {
 input.Name = strings.TrimSpace(input.Name)

 if err := validateForCreateRoom(input); err != nil {
  return nil, err
 }

 room, err := ms.Store.CreateRoom(input)
 if err != nil {
  return nil, err
 }

 return room, nil
}

func (ms *MeetingService) GetListRooms(filter models.RoomFilter) ([]*models.Room, error) {
 if err := validateListRoomsFilter(filter); err != nil {
  return nil, err
 }

 rooms, err := ms.Store.GetListRooms(filter)
 if err != nil {
  return nil, err
 }

 return rooms, nil
}

func (ms *MeetingService) GetRoomById(id uuid.UUID) (*models.Room, error) {
 room, err := ms.Store.GetRoomById(id)
 if err != nil {
  return nil, err
 }

 return room, nil
}

func (ms *MeetingService) DeleteRoomById(id uuid.UUID) error {
 if err := ms.Store.DeleteRoomById(id); err != nil {
  return err
 }

 return nil
}
