package repos

import (
	"github.com/jmoiron/sqlx"
)

// FlightRecord represents a row read from the excel file exported from Raddix
type FlightRecord struct {
	AyTax                string `db:"ay_tax"`
	BaseFare             string `db:"base_fare"`
	BookDate             string `db:"book_date"`
	ChargeBalance        string `db:"charge_balance"`
	ConfirmationNum      string `db:"confirmation_num"`
	DepartureDate        string `db:"departure_date"`
	FareClass            string `db:"fare_class"`
	FirstName            string `db:"first_name"`
	FlightNum            string `db:"flight_num"`
	FlightStatus         string `db:"flight_status"`
	FlightStatusGroup    string `db:"flight_status_group"`
	FreqFlyer            string `db:"freq_flyer"`
	FromAirport          string `db:"from_airport"`
	HasBoardingPass      string `db:"has_boarding_pass"`
	HasBoardingPassLabel string `db:"has_boarding_pass_label"`
	HasPayment           string `db:"has_payment"`
	HasPaymentLabel      string `db:"has_payment_label"`
	LastName             string `db:"last_name"`
	MarketingCarrierCode string `db:"marketing_carrier_code"`
	NonRev               string `db:"non_rev"`
	NonRevLabel          string `db:"non_rev_label"`
	PaymentAmount        string `db:"payment_amount"`
	ATax                 string `db:"a_tax"`
	RecordNum            string `db:"record_num"`
	ReChannel            string `db:"res_channel"`
	ResSegStatus         string `db:"res_seg_status"`
	ResSegStatusLabel    string `db:"res_seg_status_label"`
	SavedFBCode          string `db:"saved_fb_code"`
	SegmentBalance       string `db:"segment_balance"`
	SSRAmount            string `db:"ssr_amount"`
	ToAirport            string `db:"to_airport"`
	USTax                string `db:"us_tax"`
	XFTax                string `db:"xf_tax"`
	ZPTax                string `db:"zp_tax"`
	Remarks              string `db:"remarks"`
}

// FlightRecordsRepo implements methods for querying flight records in db
type FlightRecordsRepo struct {
	db *sqlx.DB
}

// NewFlightRecordsRepo returns a pointer to a flight records repo
func NewFlightRecordsRepo(db *sqlx.DB) *FlightRecordsRepo {
	return &FlightRecordsRepo{
		db: db,
	}
}

// GetUnprocessedFlightRecords returns a slice of pointers to
// flight records that are yet to be processed
func (repo *FlightRecordsRepo) GetUnprocessedFlightRecords() ([]*FlightRecord, error) {
	rows, err := repo.db.Queryx("SELECT * FROM flightrecords")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []*FlightRecord
	for rows.Next() {
		rec := new(FlightRecord)
		err = rows.StructScan(&rec)
		if err != nil {
			return nil, err
		}

		records = append(records, rec)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
