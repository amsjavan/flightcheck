package atta

const (
	Failed = `
	{"result":[],"isSuccessful":true,"status":"OK","code":0}
	`

	Success = `
	{
		"result": [
		  {
			"origin": "IKA",
			"destination": "NJF",
			"originTitle": "تهران - امام خمینی (ره)",
			"destinationTitle": "نجف",
			"departureDateTime": "2023-08-26T06:00:00",
			"departureTime": "06:00",
			"departureJalaliDate": "1402/06/04",
			"arrivalDateTime": "2023-08-26T07:00:00",
			"arrivalTime": "07:00",
			"arrivalJalaliDate": "1402/06/04",
			"flightNo": 6625,
			"aircraftTypeCode": "M83",
			"class": "Y",
			"capacity": "C",
			"isSelectable": false,
			"isCapacityFull": true,
			"isSeatSelectable": false,
			"isDomestic": false,
			"capacityTitle": "ظرفیت تکمیل است!",
			"fare": {
			  "id": 0,
			  "route": "IKA-NJF",
			  "flightClass": "Y",
			  "adultFare": 29697000,
			  "adultFareDis": 29697000,
			  "childFare": 29697000,
			  "childFareDis": 29697000,
			  "infantFare": 2970000,
			  "infantFareDis": 2970000,
			  "adultTotalPrice": 40000000,
			  "adultTotalPriceDis": 40000000,
			  "childTotalPrice": 40000000,
			  "childTotalPriceDis": 40000000,
			  "infantTotalPrice": 3588000,
			  "infantTotalPriceDis": 3588000,
			  "hasBox": false,
			  "countBox": 0,
			  "discount": 0,
			  "childTaxes": "IR:350000.0,EN_Desc:IRAN INTERNATIONAL AIRPORT TAX,FA_Desc:عوارض فرودگاهي پرواز هاي بين المللي$V0:2673000.0,EN_Desc:VAT,FA_Desc:ماليات بر ارزش افزوده$I6:1646000.0,EN_Desc:PASSENGER SAFETY OVERSIGHT SERVICE,FA_Desc:عوارض امنيت مسافر$IQ:5634000.0,EN_Desc:IRAQ / DEPARTURE TAX,FA_Desc:عوارض فرودگاه عراق$",
			  "adultTaxes": "IR:350000.0,EN_Desc:IRAN INTERNATIONAL AIRPORT TAX,FA_Desc:عوارض فرودگاهي پرواز هاي بين المللي$V0:2673000.0,EN_Desc:VAT,FA_Desc:ماليات بر ارزش افزوده$I6:1646000.0,EN_Desc:PASSENGER SAFETY OVERSIGHT SERVICE,FA_Desc:عوارض امنيت مسافر$IQ:5634000.0,EN_Desc:IRAQ / DEPARTURE TAX,FA_Desc:عوارض فرودگاه عراق$",
			  "infantTaxes": "IR:350000.0,EN_Desc:IRAN INTERNATIONAL AIRPORT TAX,FA_Desc:عوارض فرودگاهي پرواز هاي بين المللي$V0:268000.0,EN_Desc:VAT,FA_Desc:ماليات بر ارزش افزوده$",
			  "crcnRule": [
				"از لحظه صدور تا 24 ساعت قبل از پرواز-25درصد جريمه",
				" از 24 ساعت قبل از پرواز و پس از آن-50درصد جريمه"
			  ],
			  "baggageAllowanceWeight": "25 KG",
			  "baggageAllowancePieces": "1 PS",
			  "kuAdult": 0,
			  "kuChild": 0,
			  "kuInfant": 0
			},
			"longDepartureDateTime": {
			  "year": 1402,
			  "month": 6,
			  "day": 4,
			  "weekday": "شنبه",
			  "monthName": "شهریور",
			  "weekOfYear": 24,
			  "dayOfWeek": 6,
			  "isLeapYear": false
			},
			"flightDuration": "یک ساعت",
			"discount": 0,
			"key": "",
			"baggagePrice": 1145000
		  }
		],
		"isSuccessful": true,
		"status": "OK",
		"code": 0
	  }
	`

	OldSuccess = `
	{
		"result": [
		  {
			"origin": "IKA",
			"destination": "NJF",
			"originTitle": "تهران - امام خمینی (ره)",
			"destinationTitle": "نجف",
			"departureDateTime": "2022-09-17T22:30:00",
			"departureTime": "22:30",
			"departureJalaliDate": "1401/06/26",
			"arrivalDateTime": "2022-09-17T22:15:00",
			"arrivalTime": "22:15",
			"arrivalJalaliDate": "1401/06/26",
			"flightNo": 6690,
			"aircraftTypeCode": "M83",
			"class": "E",
			"capacity": "A",
			"classPrice": 30000000,
			"isSelectable": true,
			"isCapacityFull": false,
			"isSeatSelectable": false,
			"isDomestic": false,
			"capacityTitle": "ظرفیت موجود",
			"fare": {
			  "id": 0,
			  "route": "IKA-NJF",
			  "flightClass": "E",
			  "adultFare": 22574000,
			  "adultFareDis": 22574000,
			  "childFare": 22574000,
			  "childFareDis": 22574000,
			  "infantFare": 2258000,
			  "infantFareDis": 2258000,
			  "adultTotalPrice": 30000000,
			  "adultTotalPriceDis": 30000000,
			  "childTotalPrice": 30000000,
			  "childTotalPriceDis": 30000000,
			  "infantTotalPrice": 2812000,
			  "infantTotalPriceDis": 2812000,
			  "hasBox": false,
			  "countBox": 0,
			  "discount": 0,
			  "childTaxes": "IR:350000.0,EN_Desc:IRAN INTERNATIONAL AIRPORT TAX,FA_Desc:عوارض فرودگاهي پرواز هاي بين المللي$V0:2032000.0,EN_Desc:VAT,FA_Desc:ماليات بر ارزش افزوده$I6:1059000.0,EN_Desc:PASSENGER SAFETY OVERSIGHT SERVICE,FA_Desc:عوارض امنيت مسافر$IQ:3985000.0,EN_Desc:IRAQ / DEPARTURE TAX,FA_Desc:عوارض فرودگاه عراق$",
			  "adultTaxes": "IR:350000.0,EN_Desc:IRAN INTERNATIONAL AIRPORT TAX,FA_Desc:عوارض فرودگاهي پرواز هاي بين المللي$V0:2032000.0,EN_Desc:VAT,FA_Desc:ماليات بر ارزش افزوده$I6:1059000.0,EN_Desc:PASSENGER SAFETY OVERSIGHT SERVICE,FA_Desc:عوارض امنيت مسافر$IQ:3985000.0,EN_Desc:IRAQ / DEPARTURE TAX,FA_Desc:عوارض فرودگاه عراق$",
			  "infantTaxes": "IR:350000.0,EN_Desc:IRAN INTERNATIONAL AIRPORT TAX,FA_Desc:عوارض فرودگاهي پرواز هاي بين المللي$V0:204000.0,EN_Desc:VAT,FA_Desc:ماليات بر ارزش افزوده$",
			  "crcnRule": [
				"در بازه 24 ساعت مانده به پرواز و پس از آن  - 50 درصدجريمه",
				"از لحظه صدور تا 24 ساعت قبل از پرواز  - * درصدجريمه"
			  ],
			  "kuAdult": 0,
			  "kuChild": 0,
			  "kuInfant": 0
			},
			"longDepartureDateTime": {
			  "year": 1401,
			  "month": 6,
			  "day": 26,
			  "weekday": "شنبه",
			  "monthName": "شهریور",
			  "weekOfYear": 27,
			  "dayOfWeek": 6,
			  "isLeapYear": false
			},
			"flightDuration": "",
			"discount": 0,
			"key": "",
			"baggagePrice": 1050000
		  }
		],
		"isSuccessful": true,
		"status": "OK",
		"code": 0
	  }
	`
)
