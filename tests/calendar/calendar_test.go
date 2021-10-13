package calendar

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestGETEvent(t *testing.T) {
	events := StubCalendarStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &CalendarServer{&events}
	t.Run("Checks correct time work", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/events/1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got, err := strconv.Atoi(response.Body.String())

		if err != nil {
			t.Errorf("Response parsing error. Incorrect day received.")
		}

		want := time.Now().Day()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

type StubCalendarStore struct {
	scores map[string]int
}

func (s *StubCalendarStore) GetDate(name string) int {
	score := s.scores[name]
	return score
}
