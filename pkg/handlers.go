package pkg

import (
	"net/http"
)

// StatusHandler is the handler for checking the status
// of this service
func StatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
