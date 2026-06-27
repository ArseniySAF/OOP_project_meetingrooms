package api

import (
	"encoding/json"
	"log"
	"meeting-rooms/internal/models"
	"meeting-rooms/myerrors"
	"net/http"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (hh *HttpHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var requestRoom RoomCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&requestRoom); err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	input := models.InputRoom{
		Name:      requestRoom.Name,
		Capacity:  requestRoom.Capacity,
		Floor:     requestRoom.Floor,
		Equipment: requestRoom.Equipment,
	}

	room, err := hh.MeetingService.CreateRoom(input)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(room, "", "	")
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		log.Println(err)
	}

}

func (hh *HttpHandler) ListRooms(w http.ResponseWriter, r *http.Request, params ListRoomsParams) {
	filter := models.RoomFilter{
		Floor:       params.Floor,
		MinCapacity: params.MinCapacity,
		Limit:       params.Limit,
		Offset:      params.Offset,
	}

	rooms, err := hh.MeetingService.GetListRooms(filter)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(rooms, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}

func (hh *HttpHandler) GetRoom(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	room, err := hh.MeetingService.GetRoomById(id)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(room, "", "	")
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Println(err)
	}
}

func (hh *HttpHandler) DeleteRoom(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	if err := hh.MeetingService.DeleteRoomById(id); err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
