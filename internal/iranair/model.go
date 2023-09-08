package iranair

type Response struct {
	Success bool        `json:"success"`
	Result  ResultModel `json:"result"`
}

type ResultModel struct {
	Journeies []JourneySegDTOs `json:"journeySegDTOs"`
	SalesDate string           `json:"salesDate"`
}

type JourneySegDTOs struct {
	DTOs []SegDTOs `json:"journeySegDTOs"`
}

type SegDTOs struct {
	SeatAvailability bool `json:"seatAvailability"`
	SeatFull         bool `json:"seatFull"`
}
