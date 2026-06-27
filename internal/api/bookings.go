package api

import (
	"encoding/json"
	"log"
	"meeting-rooms/internal/models"
	"meeting-rooms/myerrors"
	"net/http"
	"time"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (hh *HttpHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var requestBooking BookingCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBooking); err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	userIdString := r.Header.Get("X-USER-ID")

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusUnauthorized)
		return
	}

	input := models.InputBooking{
		RoomId:         requestBooking.RoomId,
		UserId:         userId,
		Title:          requestBooking.Title,
		Description:    requestBooking.Description,
		StartTime:      requestBooking.StartTime,
		EndTime:        requestBooking.EndTime,
		AttendeesCount: *requestBooking.AttendeesCount,
	}

	booking, err := hh.MeetingService.CreateBooking(input)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(booking, "", "	")
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

func (hh *HttpHandler) ListBookings(w http.ResponseWriter, r *http.Request, params ListBookingsParams) {
	filter := models.BookingFilter{
		RoomId:   params.RoomId,
		DateFrom: params.DateFrom,
		DateTo:   params.DateTo,
		Limit:    params.Limit,
		Offset:   params.Offset,
	}

	bookings, err := hh.MeetingService.GetListBookings(filter)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(bookings, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}

func (hh *HttpHandler) GetMyBookings(w http.ResponseWriter, r *http.Request, params GetMyBookingsParams) {
	filter := models.MyBookingFilter{
		Limit:  params.Limit,
		Offset: params.Offset,
	}

	userIdString := r.Header.Get("X-USER-ID")

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusUnauthorized)
		return
	}

	myBookings, err := hh.MeetingService.GetMyBookings(userId, filter)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(myBookings, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}

func (hh *HttpHandler) CancelBooking(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userIdString := r.Header.Get("X-USER-ID")

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusUnauthorized)
		return
	}

	if err := hh.MeetingService.DeleteMyBookingById(id, userId); err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
