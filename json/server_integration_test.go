package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPOSTScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPOSTScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPOSTScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}