package atta

type Response struct {
	Success bool          `json:"isSuccessful"`
	Result  []ResultModel `json:"result"`
}

type ResultModel struct {
	Origin              string `json:"origin"`
	Destination         string `json:"destination"`
	OriginTitle         string `json:"originTitle"`
	DestinationTitle    string `json:"destinationTitle"`
	DepartureDateTime   string `json:"departureDateTime"`
	DepartureTime       string `json:"departureTime"`
	DepartureJalaliDate string `json:"departureJalaliDate"`
	ArrivalDateTime     string `json:"arrivalDateTime"`
	ArrivalTime         string `json:"arrivalTime"`
	ArrivalJalaliDate   string `json:"arrivalJalaliDate"`
	FlightNo            int    `json:"flightNo"`
	AircraftTypeCode    string `json:"aircraftTypeCode"`
	Class               string `json:"class"`
	Capacity            string `json:"capacity"`
	IsSelectable        bool   `json:"isSelectable"`
	IsCapacityFull      bool   `json:"isCapacityFull"`
	IsSeatSelectable    bool   `json:"isSeatSelectable"`
	IsDomestic          bool   `json:"isDomestic"`
	CapacityTitle       string `json:"capacityTitle"`
}
