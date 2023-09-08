package meraj

type Response struct {
	Success bool          `json:"Success"`
	Result  []ResultModel `json:"Result"`
}

type ResultModel struct {
	FlightGroups []FlightGroupsModel `json:"FlightGroups"`
}

type FlightGroupsModel struct {
	Flights []FlightsModel `json:"Flights"`
}

type FlightsModel struct {
	Id     int          `json:"Id"`
	Prices []PriceModel `json:"Prices"`
}

type PriceModel struct {
	Amount float32 `json:"Amount"`
}
