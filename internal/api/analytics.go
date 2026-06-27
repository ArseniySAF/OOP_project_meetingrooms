package api

import (
	"encoding/json"
	"log"
	"meeting-rooms/myerrors"
	"net/http"
	"time"
)

func (hh *HttpHandler) GetRoomAnalytics(w http.ResponseWriter, r *http.Request) {
	analytics, err := hh.MeetingService.GetAnalytics()
	if err != nil {
		errDTO := myerrors.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(analytics, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}
