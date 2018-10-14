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
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
		&FlightRecord{AyTax: "AyTax", BaseFare: "BaseFare", BookDate: "BookDate", ChargeBalance: "ChargeBalance", ConfirmationNum: "ConfirmationNum", DepartureDate: "DepartureDate", FareClass: "FareClass", FirstName: "FirstName", FlightNum: "FlightNum", FlightStatus: "FlightStatus", FlightStatusGroup: "FlightStatusGroup", FreqFlyer: "FreqFlyer", FromAirport: "FromAirport", HasBoardingPass: "HasBoardingPass", HasBoardingPassLabel: "HasBoardingPassLabel", HasPayment: "HasPayment", HasPaymentLabel: "HasPaymentLabel", LastName: "LastName", MarketingCarrierCode: "MarketingCarrierCode", NonRev: "NonRev", NonRevLabel: "NonRevLabel", PaymentAmount: "PaymentAmount", QTax: "QTax", RecordNum: "RecordNum", ReChannel: "ReChannel", ResSegStatus: "ResSegStatus", ResSegStatusLabel: "ResSegStatusLabel", SavedFBCode: "SavedFBCode", SegmentBalance: "SegmentBalance", SSRAmount: "SSRAmount", ToAirport: "ToAirport", USTax: "USTax", XFTax: "XFTax", ZPTax: "ZPTax", Remarks: "Remarks", IsProcessed: false},
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
