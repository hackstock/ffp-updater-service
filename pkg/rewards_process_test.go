package pkg

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"go.uber.org/zap"
)

var (
	mockFlightRecords = []*FlightRecord{
		&FlightRecord{ID: 1},
		&FlightRecord{ID: 2},
		&FlightRecord{ID: 3},
		&FlightRecord{ID: 4},
		&FlightRecord{ID: 5},
		&FlightRecord{ID: 6},
		&FlightRecord{ID: 7},
		&FlightRecord{ID: 8},
		&FlightRecord{ID: 9},
		&FlightRecord{ID: 10},
	}
)

type MockResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func mockServer(shouldFail bool) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ffp/flightrecords/unprocessed", func(w http.ResponseWriter, r *http.Request) {
		res := MockResponse{}
		if shouldFail {
			res.Status = http.StatusInternalServerError
			res.Data = nil

			w.WriteHeader(http.StatusInternalServerError)
		} else {
			res.Status = http.StatusOK
			res.Data = mockFlightRecords

			w.WriteHeader(http.StatusOK)
		}

		resJSON, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.Write(resJSON)
	})

	mux.HandleFunc("/ffp/apply", func(w http.ResponseWriter, r *http.Request) {
		if shouldFail {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "something bad happened", http.StatusInternalServerError)
			return
		}

		var fr FlightRecord
		err := json.NewDecoder(r.Body).Decode(&fr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := struct {
			Status int          `json:"status"`
			Data   FlightRecord `json:"data"`
		}{
			Status: http.StatusOK,
			Data:   fr,
		}

		payload, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)

	})

	return mux
}

func TestGetUnprocessedFiles_ShouldPass(t *testing.T) {
	srv := httptest.NewServer(mockServer(false))
	defer srv.Close()

	cl := &http.Client{Timeout: 15 * time.Second}
	rp := NewRewardsProcess(srv.URL, cl, zap.NewNop())

	frRes, err := rp.getUnprocessedFlightRecords()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(frRes.Data) != len(mockFlightRecords) {
		t.Fatalf("expected %d records, got %d", len(mockFlightRecords), len(frRes.Data))
	}

	if !reflect.DeepEqual(frRes.Data, mockFlightRecords) {
		t.Fatalf("expected %v, got %v", mockFlightRecords, frRes.Data)
	}
}

func TestGetUnprocessedFiles_ShouldFail(t *testing.T) {
	srv := httptest.NewServer(mockServer(true))
	defer srv.Close()

	cl := &http.Client{Timeout: 15 * time.Second}
	rp := NewRewardsProcess(srv.URL, cl, zap.NewNop())

	frRes, err := rp.getUnprocessedFlightRecords()
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	if frRes != nil {
		t.Fatalf("expected nil, got %v", frRes)
	}
}

func TestApplyReward_ShouldPass(t *testing.T) {
	srv := httptest.NewServer(mockServer(false))
	defer srv.Close()

	cl := &http.Client{Timeout: 15 * time.Second}
	rp := NewRewardsProcess(srv.URL, cl, zap.NewNop())

	for _, fr := range mockFlightRecords {
		err := rp.applyReward(fr)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}
}

func TestApplyReward_ShouldFail(t *testing.T) {
	srv := httptest.NewServer(mockServer(true))
	defer srv.Close()

	cl := &http.Client{Timeout: 15 * time.Second}
	rp := NewRewardsProcess(srv.URL, cl, zap.NewNop())

	err := rp.applyReward(mockFlightRecords[0])
	if err == nil {
		t.Fatalf("expected error, got none")
	}

}
