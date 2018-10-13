package pkg

import (
	"net/http"

	"go.uber.org/zap"
)

// FlightRecord ...
type FlightRecord struct {
	AyTax                string `json:"ayTax"`
	BaseFare             string `json:"baseFare"`
	BookDate             string `json:"bookDate"`
	ChargeBalance        string `json:"chargeBalance"`
	ConfirmationNum      string `json:"confirmationNum"`
	DepartureDate        string `json:"departureDate"`
	FareClass            string `json:"fareClassCode"`
	FirstName            string `json:"firstName"`
	FlightNum            string `json:"flightNum"`
	FlightStatus         string `json:"flightStatus"`
	FlightStatusGroup    string `json:"flightStatusGroup"`
	FreqFlyer            string `json:"freqFlyer"`
	FromAirport          string `json:"fromAirport"`
	HasBoardingPass      string `json:"hasBoardingPass"`
	HasBoardingPassLabel string `json:"hasBoardingPassLabel"`
	HasPayment           string `json:"hasPayment"`
	HasPaymentLabel      string `json:"hasPaymentLabel"`
	LastName             string `json:"lastName"`
	MarketingCarrierCode string `json:"marketingCarrierCode"`
	NonRev               string `json:"nonRev"`
	NonRevLabel          string `json:"nonRevLabel"`
	PaymentAmount        string `json:"paymentAmount"`
	QTax                 string `json:"qTax"`
	RecordNum            string `json:"recordNum"`
	ReChannel            string `json:"resChannel"`
	ResSegStatus         string `json:"resSegStatus"`
	ResSegStatusLabel    string `json:"resSegStatusLabel"`
	SavedFBCode          string `json:"savedFbCode"`
	SegmentBalance       string `json:"segmentBalance"`
	SSRAmount            string `json:"ssrAmount"`
	ToAirport            string `json:"toAirport"`
	USTax                string `json:"usTax"`
	XFTax                string `json:"xfTax"`
	ZPTax                string `json:"zpTax"`
	Remarks              string `json:"remarks"`
	IsProcessed          bool   `json:"isProcessed"`
}

// RewardsProcess ...
type RewardsProcess struct {
	host   string
	client *http.Client
	logger *zap.Logger
}

// NewRewardsProcess ...
func NewRewardsProcess(host string, client *http.Client, logger *zap.Logger) *RewardsProcess {
	return &RewardsProcess{
		host:   host,
		client: client,
		logger: logger,
	}
}

// Run ...
func (p *RewardsProcess) Run() {

}
