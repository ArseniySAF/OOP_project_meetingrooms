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

	// TODO: написать Middleware, чтобы извлекать userId(sub) из claims
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

	_ = filter

	// TODO: написать Middleware, чтобы извлекать userId(sub) из claims
}

func (hh *HttpHandler) CancelBooking(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO: написать Middleware, чтобы извлекать userId(sub) из claims

	// + какую то проверку на уровне http
}
