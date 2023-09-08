package airtour

type Response struct {
	HasError bool        `json:"HasError"`
	Result   ResultModel `json:"ResultObject"`
}

type ResultModel struct {
	Successful   bool                `json:"Successful"`
	FlightOffers []FlightGroupsModel `json:"FlightOffers"`
}

type FlightGroupsModel struct {
	Date      string `json:"Date"`
	TotalFare int    `json:"TotalFare"`
}
