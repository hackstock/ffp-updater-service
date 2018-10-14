package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

// FlightRecord represent a FFP user's flight record
// to be used for calculating points to be awared
type FlightRecord struct {
	ID                   int    `json:"id"`
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

// FRResponse is a struct for holding API response
// from endpoint for fetching unprocessed flight records
type FRResponse struct {
	Status int             `json:"status"`
	Data   []*FlightRecord `json:"data"`
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

// Run fetches unprocessed flight records from API host
// and applies FFP rewards to each of the records
func (p *RewardsProcess) Run() error {
	res, err := p.getUnprocessedFlightRecords()
	if err != nil {
		return fmt.Errorf("failed fetching unprocessed flight records : %v", err)
	}

	p.logger.Info("fetched unprocessed flight records",
		zap.Int("count", len(res.Data)))

	for _, rec := range res.Data {
		err = p.applyReward(rec)
		if err != nil {
			p.logger.Warn("failed applying rewards",
				zap.Error(err))
		}

	}

	return nil
}

func (p *RewardsProcess) getUnprocessedFlightRecords() (*FRResponse, error) {
	url := fmt.Sprintf("%s/ffp/flightrecords/unprocessed", p.host)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		frRes := new(FRResponse)
		err = json.NewDecoder(res.Body).Decode(frRes)
		if err != nil {
			return nil, err
		}
		return frRes, nil
	}

	_, _ = io.Copy(ioutil.Discard, res.Body)
	return nil, errors.New(res.Status)
}

func (p *RewardsProcess) applyReward(fr *FlightRecord) error {
	url := fmt.Sprintf("%s/ffp/apply", p.host)
	payload, err := json.Marshal(fr)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := p.client.Do(req)
	if err != nil {
		return err
	}

	_, _ = io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	return nil
}
