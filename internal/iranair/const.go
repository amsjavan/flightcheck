package iranair

const (
	Failed = `
	{
		"success": true,
		"message": null,
		"errorCode": null,
		"result": {
		  "journeySegDTOs": [
			{
			  "originCode": "IKA",
			  "destinationCode": "NJF",
			  "direction": "OUTBOUND",
			  "openOND": false,
			  "journeySegDTOs": [
				{
				  "seatAvailability": false,
				  "seatFull": true,
				  "fareProductDetails": [],
				  "flightSegments": [
					{
					  "flightID": 101288,
					  "fltSegID": 103971,
					  "flightScheduleID": 41492,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5333",
					  "direction": null,
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A321-200",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1662876600000,
					  "arrivalDateUCT": 1662882300000,
					  "departureDateLocal": 1662892800000,
					  "arrivalDateLocal": 1662893100000,
					  "departureDateLocalStr": "11/09/2022 10:40:00",
					  "arrivalDateLocalStr": "11/09/2022 10:45:00",
					  "departureDateUTCStr": "11/09/2022 06:10:00",
					  "arrivalDateUTCStr": "11/09/2022 07:45:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": false,
					  "flightFull": true,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 0,
					  "legINFCount": 0,
					  "actualDirection": null,
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": false,
				  "seatFull": true,
				  "fareProductDetails": [],
				  "flightSegments": [
					{
					  "flightID": 83419,
					  "fltSegID": 85672,
					  "flightScheduleID": 45514,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5319",
					  "direction": null,
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A300B4-203",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1662862200000,
					  "arrivalDateUCT": 1662867900000,
					  "departureDateLocal": 1662878400000,
					  "arrivalDateLocal": 1662878700000,
					  "departureDateLocalStr": "11/09/2022 06:40:00",
					  "arrivalDateLocalStr": "11/09/2022 06:45:00",
					  "departureDateUTCStr": "11/09/2022 02:10:00",
					  "arrivalDateUTCStr": "11/09/2022 03:45:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": false,
					  "flightFull": true,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 0,
					  "legINFCount": 0,
					  "actualDirection": null,
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": false,
				  "seatFull": true,
				  "fareProductDetails": [],
				  "flightSegments": [
					{
					  "flightID": 106963,
					  "fltSegID": 109686,
					  "flightScheduleID": 47072,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5321",
					  "direction": null,
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A300B4-203",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1662880200000,
					  "arrivalDateUCT": 1662885900000,
					  "departureDateLocal": 1662896400000,
					  "arrivalDateLocal": 1662896700000,
					  "departureDateLocalStr": "11/09/2022 11:40:00",
					  "arrivalDateLocalStr": "11/09/2022 11:45:00",
					  "departureDateUTCStr": "11/09/2022 07:10:00",
					  "arrivalDateUTCStr": "11/09/2022 08:45:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": false,
					  "flightFull": true,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 0,
					  "legINFCount": 0,
					  "actualDirection": null,
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": false,
				  "seatFull": true,
				  "fareProductDetails": [],
				  "flightSegments": [
					{
					  "flightID": 100994,
					  "fltSegID": 103677,
					  "flightScheduleID": 46993,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5339",
					  "direction": null,
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A321-200",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1662899400000,
					  "arrivalDateUCT": 1662905100000,
					  "departureDateLocal": 1662915600000,
					  "arrivalDateLocal": 1662915900000,
					  "departureDateLocalStr": "11/09/2022 17:00:00",
					  "arrivalDateLocalStr": "11/09/2022 17:05:00",
					  "departureDateUTCStr": "11/09/2022 12:30:00",
					  "arrivalDateUTCStr": "11/09/2022 14:05:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": false,
					  "flightFull": true,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 0,
					  "legINFCount": 8,
					  "actualDirection": null,
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": false,
				  "seatFull": true,
				  "fareProductDetails": [],
				  "flightSegments": [
					{
					  "flightID": 100508,
					  "fltSegID": 103191,
					  "flightScheduleID": 46971,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5329",
					  "direction": null,
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A321-200",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1662858600000,
					  "arrivalDateUCT": 1662864300000,
					  "departureDateLocal": 1662874800000,
					  "arrivalDateLocal": 1662875100000,
					  "departureDateLocalStr": "11/09/2022 05:40:00",
					  "arrivalDateLocalStr": "11/09/2022 05:45:00",
					  "departureDateUTCStr": "11/09/2022 01:10:00",
					  "arrivalDateUTCStr": "11/09/2022 02:45:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": false,
					  "flightFull": true,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 0,
					  "legINFCount": 10,
					  "actualDirection": null,
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": false,
				  "seatFull": true,
				  "fareProductDetails": [],
				  "flightSegments": [
					{
					  "flightID": 83915,
					  "fltSegID": 86168,
					  "flightScheduleID": 46276,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5325",
					  "direction": null,
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A310-300",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1662892500000,
					  "arrivalDateUCT": 1662898200000,
					  "departureDateLocal": 1662908700000,
					  "arrivalDateLocal": 1662909000000,
					  "departureDateLocalStr": "11/09/2022 15:05:00",
					  "arrivalDateLocalStr": "11/09/2022 15:10:00",
					  "departureDateUTCStr": "11/09/2022 10:35:00",
					  "arrivalDateUTCStr": "11/09/2022 12:10:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": false,
					  "flightFull": true,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 0,
					  "legINFCount": 10,
					  "actualDirection": null,
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": false,
				  "seatFull": true,
				  "fareProductDetails": [],
				  "flightSegments": [
					{
					  "flightID": 106987,
					  "fltSegID": 109660,
					  "flightScheduleID": 47122,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5317",
					  "direction": null,
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A321-200",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1662917400000,
					  "arrivalDateUCT": 1662923100000,
					  "departureDateLocal": 1662933600000,
					  "arrivalDateLocal": 1662933900000,
					  "departureDateLocalStr": "11/09/2022 22:00:00",
					  "arrivalDateLocalStr": "11/09/2022 22:05:00",
					  "departureDateUTCStr": "11/09/2022 17:30:00",
					  "arrivalDateUTCStr": "11/09/2022 19:05:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": false,
					  "flightFull": true,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 0,
					  "legINFCount": 0,
					  "actualDirection": null,
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				}
			  ],
			  "daySummaryDTOMap": {},
			  "daySeatFullMap": {
				"0": true
			  },
			  "journeyAvailabilityMap": {
				"0": true
			  }
			}
		  ],
		  "priceQuote": null,
		  "salesDate": "05/09/2022 15:07:00:000"
		}
	  }
	`
	Success = `
	{
		"success": true,
		"message": null,
		"errorCode": null,
		"result": {
		  "journeySegDTOs": [
			{
			  "originCode": "IKA",
			  "destinationCode": "NJF",
			  "direction": "OUTBOUND",
			  "openOND": false,
			  "journeySegDTOs": [
				{
				  "seatAvailability": true,
				  "seatFull": false,
				  "fareProductDetails": [
					{
					  "fareProduct": "P1",
					  "fareIDs": [
						50814
					  ],
					  "availableSeatCount": 6,
					  "amount": "60000000",
					  "currencyCode": "IRR",
					  "selected": null,
					  "searchUuid": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0",
					  "fareProductPQ": {
						"paxTypes": [
						  "ADT"
						],
						"baseFares": [
						  {
							"segmentCode": "IKA-NJF",
							"marriedSegment": false,
							"bookingClass": "E",
							"fareClass": "EIRARBW",
							"fareRuleId": 6203,
							"fareId": 50814,
							"fareProduct": "P1",
							"salesType": "STANDARD",
							"bundled": false,
							"addOnFare": false,
							"addOnFareClass": null,
							"linkedFltSegIDs": [
							  103649
							],
							"paxPrice": {
							  "ADT": "23270000"
							},
							"vatAmount": {},
							"discountAmount": {},
							"freeVoucherAmount": null,
							"extraSeatAmount": {},
							"fareRuleSurchargeAmount": {},
							"segmentCodeWithMarriedStatus": "IKA-NJF"
						  }
						],
						"baseFaresTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "23270000"
						  },
						  "total": "46540000"
						},
						"taxes": [
						  {
							"name": "DEPARTURE_ARRIVAL TAX (IQ)",
							"chargeCode": "IQ",
							"chargeDescription": "DEPARTURE_ARRIVAL TAX",
							"paxPrice": {
							  "ADT": "3990000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "YQ Surcharge (YQ)",
							"chargeCode": "YQ",
							"chargeDescription": "YQ Surcharge",
							"paxPrice": {
							  "ADT": "1330000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "Airport Tax (IR)",
							"chargeCode": "IR",
							"chargeDescription": "Airport Tax",
							"paxPrice": {
							  "ADT": "350000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "Passenger Safety Oversight Services Fees (I6)",
							"chargeCode": "I6",
							"chargeDescription": "Passenger Safety Oversight Services Fees",
							"paxPrice": {
							  "ADT": "1060000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  }
						],
						"taxesTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "6730000"
						  },
						  "total": "13460000"
						},
						"charges": [],
						"chargesTotal": null,
						"perPaxTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "30000000"
						  },
						  "total": "60000000"
						},
						"total": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "60000000"
						  },
						  "total": "60000000"
						},
						"commissionTotal": null,
						"commissionOrTotal": null,
						"agentCommissionTotal": null,
						"discountTotal": null,
						"extraSeatAmountTotal": {},
						"currency": "IRR",
						"selectedCurrencyTotal": null,
						"ancillaryBundle": {
						  "103649": false
						},
						"pqLight": {
						  "fares": [
							{
							  "fareId": 50814,
							  "fareRuleID": 6203,
							  "fareSeq": 1,
							  "amount": {
								"ADT": "23270000"
							  },
							  "segmentSpan": [
								{
								  "fltSegID": 103649,
								  "cabinClassID": 2,
								  "rbdID": 20,
								  "rbdInvType": "STANDARD",
								  "invFltSegCCID": 177556,
								  "actualDirection": "OUTBOUND",
								  "flightID": 100966,
								  "segmentCode": "IKA-NJF",
								  "ownerAirlineID": 15,
								  "puGroupID": 5055634421933952000,
								  "flightType": "INT",
								  "flightOperationType": "STANDARD",
								  "overBookingOnDisruption": false,
								  "addonSegment": false
								}
							  ],
							  "segmentSpanCode": "IKA-NJF"
							}
						  ],
						  "fltSegs": {
							"103649": {
							  "fltSegID": 103649,
							  "cabinClassID": 2,
							  "rbdID": 20,
							  "rbdInvType": "STANDARD",
							  "invFltSegCCID": 177556,
							  "actualDirection": "OUTBOUND",
							  "flightID": 100966,
							  "segmentCode": "IKA-NJF",
							  "ownerAirlineID": 15,
							  "puGroupID": 5055634421933952000,
							  "flightType": "INT",
							  "flightOperationType": "STANDARD",
							  "overBookingOnDisruption": false,
							  "addonSegment": false
							}
						  },
						  "inventoryConsumption": {
							"adtConsumption": 2,
							"infConsumption": 0
						  },
						  "bookingCurrencyCode": "IRR",
						  "searchUID": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0"
						},
						"segmentFareAmount": {
						  "ADT": {
							"103649": "23270000"
						  }
						},
						"segmentChargeAmount": {},
						"segmentTaxAmount": {
						  "ADT": {
							"103649": "6730000"
						  }
						},
						"segmentFareChargeAmount": {
						  "ADT": {
							"103649": "30000000"
						  }
						},
						"openReturnBooking": false,
						"openReturnMaxDate": null,
						"openReturnMaxDateStr": null,
						"agentDiscountApplied": false,
						"addToDisruptionOverbookQueue": false,
						"searchUID": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0"
					  }
					}
				  ],
				  "flightSegments": [
					{
					  "flightID": 100966,
					  "fltSegID": 103649,
					  "flightScheduleID": 46999,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5339",
					  "direction": "OUTBOUND",
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A321-200",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1663417800000,
					  "arrivalDateUCT": 1663423500000,
					  "departureDateLocal": 1663434000000,
					  "arrivalDateLocal": 1663434300000,
					  "departureDateLocalStr": "17/09/2022 17:00:00",
					  "arrivalDateLocalStr": "17/09/2022 17:05:00",
					  "departureDateUTCStr": "17/09/2022 12:30:00",
					  "arrivalDateUTCStr": "17/09/2022 14:05:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": true,
					  "flightFull": false,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 6,
					  "legINFCount": 10,
					  "actualDirection": "OUTBOUND",
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": true,
				  "seatFull": false,
				  "fareProductDetails": [
					{
					  "fareProduct": "P1",
					  "fareIDs": [
						50814
					  ],
					  "availableSeatCount": 2,
					  "amount": "60000000",
					  "currencyCode": "IRR",
					  "selected": null,
					  "searchUuid": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0",
					  "fareProductPQ": {
						"paxTypes": [
						  "ADT"
						],
						"baseFares": [
						  {
							"segmentCode": "IKA-NJF",
							"marriedSegment": false,
							"bookingClass": "E",
							"fareClass": "EIRARBW",
							"fareRuleId": 6203,
							"fareId": 50814,
							"fareProduct": "P1",
							"salesType": "STANDARD",
							"bundled": false,
							"addOnFare": false,
							"addOnFareClass": null,
							"linkedFltSegIDs": [
							  109655
							],
							"paxPrice": {
							  "ADT": "23270000"
							},
							"vatAmount": {},
							"discountAmount": {},
							"freeVoucherAmount": null,
							"extraSeatAmount": {},
							"fareRuleSurchargeAmount": {},
							"segmentCodeWithMarriedStatus": "IKA-NJF"
						  }
						],
						"baseFaresTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "23270000"
						  },
						  "total": "46540000"
						},
						"taxes": [
						  {
							"name": "DEPARTURE_ARRIVAL TAX (IQ)",
							"chargeCode": "IQ",
							"chargeDescription": "DEPARTURE_ARRIVAL TAX",
							"paxPrice": {
							  "ADT": "3990000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "YQ Surcharge (YQ)",
							"chargeCode": "YQ",
							"chargeDescription": "YQ Surcharge",
							"paxPrice": {
							  "ADT": "1330000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "Airport Tax (IR)",
							"chargeCode": "IR",
							"chargeDescription": "Airport Tax",
							"paxPrice": {
							  "ADT": "350000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "Passenger Safety Oversight Services Fees (I6)",
							"chargeCode": "I6",
							"chargeDescription": "Passenger Safety Oversight Services Fees",
							"paxPrice": {
							  "ADT": "1060000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  }
						],
						"taxesTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "6730000"
						  },
						  "total": "13460000"
						},
						"charges": [],
						"chargesTotal": null,
						"perPaxTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "30000000"
						  },
						  "total": "60000000"
						},
						"total": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "60000000"
						  },
						  "total": "60000000"
						},
						"commissionTotal": null,
						"commissionOrTotal": null,
						"agentCommissionTotal": null,
						"discountTotal": null,
						"extraSeatAmountTotal": {},
						"currency": "IRR",
						"selectedCurrencyTotal": null,
						"ancillaryBundle": {
						  "109655": false
						},
						"pqLight": {
						  "fares": [
							{
							  "fareId": 50814,
							  "fareRuleID": 6203,
							  "fareSeq": 1,
							  "amount": {
								"ADT": "23270000"
							  },
							  "segmentSpan": [
								{
								  "fltSegID": 109655,
								  "cabinClassID": 2,
								  "rbdID": 20,
								  "rbdInvType": "STANDARD",
								  "invFltSegCCID": 191494,
								  "actualDirection": "OUTBOUND",
								  "flightID": 106982,
								  "segmentCode": "IKA-NJF",
								  "ownerAirlineID": 15,
								  "puGroupID": 2673052957599616500,
								  "flightType": "INT",
								  "flightOperationType": "STANDARD",
								  "overBookingOnDisruption": false,
								  "addonSegment": false
								}
							  ],
							  "segmentSpanCode": "IKA-NJF"
							}
						  ],
						  "fltSegs": {
							"109655": {
							  "fltSegID": 109655,
							  "cabinClassID": 2,
							  "rbdID": 20,
							  "rbdInvType": "STANDARD",
							  "invFltSegCCID": 191494,
							  "actualDirection": "OUTBOUND",
							  "flightID": 106982,
							  "segmentCode": "IKA-NJF",
							  "ownerAirlineID": 15,
							  "puGroupID": 2673052957599616500,
							  "flightType": "INT",
							  "flightOperationType": "STANDARD",
							  "overBookingOnDisruption": false,
							  "addonSegment": false
							}
						  },
						  "inventoryConsumption": {
							"adtConsumption": 2,
							"infConsumption": 0
						  },
						  "bookingCurrencyCode": "IRR",
						  "searchUID": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0"
						},
						"segmentFareAmount": {
						  "ADT": {
							"109655": "23270000"
						  }
						},
						"segmentChargeAmount": {},
						"segmentTaxAmount": {
						  "ADT": {
							"109655": "6730000"
						  }
						},
						"segmentFareChargeAmount": {
						  "ADT": {
							"109655": "30000000"
						  }
						},
						"openReturnBooking": false,
						"openReturnMaxDate": null,
						"openReturnMaxDateStr": null,
						"agentDiscountApplied": false,
						"addToDisruptionOverbookQueue": false,
						"searchUID": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0"
					  }
					}
				  ],
				  "flightSegments": [
					{
					  "flightID": 106982,
					  "fltSegID": 109655,
					  "flightScheduleID": 47057,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5321",
					  "direction": "OUTBOUND",
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A300B4-203",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1663398600000,
					  "arrivalDateUCT": 1663404300000,
					  "departureDateLocal": 1663414800000,
					  "arrivalDateLocal": 1663415100000,
					  "departureDateLocalStr": "17/09/2022 11:40:00",
					  "arrivalDateLocalStr": "17/09/2022 11:45:00",
					  "departureDateUTCStr": "17/09/2022 07:10:00",
					  "arrivalDateUTCStr": "17/09/2022 08:45:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": true,
					  "flightFull": false,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 2,
					  "legINFCount": 20,
					  "actualDirection": "OUTBOUND",
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": true,
				  "seatFull": false,
				  "fareProductDetails": [
					{
					  "fareProduct": "P1",
					  "fareIDs": [
						50814
					  ],
					  "availableSeatCount": 5,
					  "amount": "60000000",
					  "currencyCode": "IRR",
					  "selected": null,
					  "searchUuid": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0",
					  "fareProductPQ": {
						"paxTypes": [
						  "ADT"
						],
						"baseFares": [
						  {
							"segmentCode": "IKA-NJF",
							"marriedSegment": false,
							"bookingClass": "E",
							"fareClass": "EIRARBW",
							"fareRuleId": 6203,
							"fareId": 50814,
							"fareProduct": "P1",
							"salesType": "STANDARD",
							"bundled": false,
							"addOnFare": false,
							"addOnFareClass": null,
							"linkedFltSegIDs": [
							  85441
							],
							"paxPrice": {
							  "ADT": "23270000"
							},
							"vatAmount": {},
							"discountAmount": {},
							"freeVoucherAmount": null,
							"extraSeatAmount": {},
							"fareRuleSurchargeAmount": {},
							"segmentCodeWithMarriedStatus": "IKA-NJF"
						  }
						],
						"baseFaresTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "23270000"
						  },
						  "total": "46540000"
						},
						"taxes": [
						  {
							"name": "DEPARTURE_ARRIVAL TAX (IQ)",
							"chargeCode": "IQ",
							"chargeDescription": "DEPARTURE_ARRIVAL TAX",
							"paxPrice": {
							  "ADT": "3990000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "YQ Surcharge (YQ)",
							"chargeCode": "YQ",
							"chargeDescription": "YQ Surcharge",
							"paxPrice": {
							  "ADT": "1330000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "Airport Tax (IR)",
							"chargeCode": "IR",
							"chargeDescription": "Airport Tax",
							"paxPrice": {
							  "ADT": "350000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "Passenger Safety Oversight Services Fees (I6)",
							"chargeCode": "I6",
							"chargeDescription": "Passenger Safety Oversight Services Fees",
							"paxPrice": {
							  "ADT": "1060000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  }
						],
						"taxesTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "6730000"
						  },
						  "total": "13460000"
						},
						"charges": [],
						"chargesTotal": null,
						"perPaxTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "30000000"
						  },
						  "total": "60000000"
						},
						"total": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "60000000"
						  },
						  "total": "60000000"
						},
						"commissionTotal": null,
						"commissionOrTotal": null,
						"agentCommissionTotal": null,
						"discountTotal": null,
						"extraSeatAmountTotal": {},
						"currency": "IRR",
						"selectedCurrencyTotal": null,
						"ancillaryBundle": {
						  "85441": false
						},
						"pqLight": {
						  "fares": [
							{
							  "fareId": 50814,
							  "fareRuleID": 6203,
							  "fareSeq": 1,
							  "amount": {
								"ADT": "23270000"
							  },
							  "segmentSpan": [
								{
								  "fltSegID": 85441,
								  "cabinClassID": 2,
								  "rbdID": 20,
								  "rbdInvType": "STANDARD",
								  "invFltSegCCID": 191528,
								  "actualDirection": "OUTBOUND",
								  "flightID": 83188,
								  "segmentCode": "IKA-NJF",
								  "ownerAirlineID": 15,
								  "puGroupID": 6054734164545851000,
								  "flightType": "INT",
								  "flightOperationType": "STANDARD",
								  "overBookingOnDisruption": false,
								  "addonSegment": false
								}
							  ],
							  "segmentSpanCode": "IKA-NJF"
							}
						  ],
						  "fltSegs": {
							"85441": {
							  "fltSegID": 85441,
							  "cabinClassID": 2,
							  "rbdID": 20,
							  "rbdInvType": "STANDARD",
							  "invFltSegCCID": 191528,
							  "actualDirection": "OUTBOUND",
							  "flightID": 83188,
							  "segmentCode": "IKA-NJF",
							  "ownerAirlineID": 15,
							  "puGroupID": 6054734164545851000,
							  "flightType": "INT",
							  "flightOperationType": "STANDARD",
							  "overBookingOnDisruption": false,
							  "addonSegment": false
							}
						  },
						  "inventoryConsumption": {
							"adtConsumption": 2,
							"infConsumption": 0
						  },
						  "bookingCurrencyCode": "IRR",
						  "searchUID": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0"
						},
						"segmentFareAmount": {
						  "ADT": {
							"85441": "23270000"
						  }
						},
						"segmentChargeAmount": {},
						"segmentTaxAmount": {
						  "ADT": {
							"85441": "6730000"
						  }
						},
						"segmentFareChargeAmount": {
						  "ADT": {
							"85441": "30000000"
						  }
						},
						"openReturnBooking": false,
						"openReturnMaxDate": null,
						"openReturnMaxDateStr": null,
						"agentDiscountApplied": false,
						"addToDisruptionOverbookQueue": false,
						"searchUID": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0"
					  }
					}
				  ],
				  "flightSegments": [
					{
					  "flightID": 83188,
					  "fltSegID": 85441,
					  "flightScheduleID": 45508,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5319",
					  "direction": "OUTBOUND",
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A300B4-203",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1663380600000,
					  "arrivalDateUCT": 1663386300000,
					  "departureDateLocal": 1663396800000,
					  "arrivalDateLocal": 1663397100000,
					  "departureDateLocalStr": "17/09/2022 06:40:00",
					  "arrivalDateLocalStr": "17/09/2022 06:45:00",
					  "departureDateUTCStr": "17/09/2022 02:10:00",
					  "arrivalDateUTCStr": "17/09/2022 03:45:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": true,
					  "flightFull": false,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 5,
					  "legINFCount": 20,
					  "actualDirection": "OUTBOUND",
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": true,
				  "seatFull": false,
				  "fareProductDetails": [
					{
					  "fareProduct": "P1",
					  "fareIDs": [
						50814
					  ],
					  "availableSeatCount": 5,
					  "amount": "60000000",
					  "currencyCode": "IRR",
					  "selected": null,
					  "searchUuid": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0",
					  "fareProductPQ": {
						"paxTypes": [
						  "ADT"
						],
						"baseFares": [
						  {
							"segmentCode": "IKA-NJF",
							"marriedSegment": false,
							"bookingClass": "E",
							"fareClass": "EIRARBW",
							"fareRuleId": 6203,
							"fareId": 50814,
							"fareProduct": "P1",
							"salesType": "STANDARD",
							"bundled": false,
							"addOnFare": false,
							"addOnFareClass": null,
							"linkedFltSegIDs": [
							  103091
							],
							"paxPrice": {
							  "ADT": "23270000"
							},
							"vatAmount": {},
							"discountAmount": {},
							"freeVoucherAmount": null,
							"extraSeatAmount": {},
							"fareRuleSurchargeAmount": {},
							"segmentCodeWithMarriedStatus": "IKA-NJF"
						  }
						],
						"baseFaresTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "23270000"
						  },
						  "total": "46540000"
						},
						"taxes": [
						  {
							"name": "DEPARTURE_ARRIVAL TAX (IQ)",
							"chargeCode": "IQ",
							"chargeDescription": "DEPARTURE_ARRIVAL TAX",
							"paxPrice": {
							  "ADT": "3990000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "YQ Surcharge (YQ)",
							"chargeCode": "YQ",
							"chargeDescription": "YQ Surcharge",
							"paxPrice": {
							  "ADT": "1330000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "Airport Tax (IR)",
							"chargeCode": "IR",
							"chargeDescription": "Airport Tax",
							"paxPrice": {
							  "ADT": "350000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  },
						  {
							"name": "Passenger Safety Oversight Services Fees (I6)",
							"chargeCode": "I6",
							"chargeDescription": "Passenger Safety Oversight Services Fees",
							"paxPrice": {
							  "ADT": "1060000"
							},
							"vatAmount": {
							  "ADT": "0"
							}
						  }
						],
						"taxesTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "6730000"
						  },
						  "total": "13460000"
						},
						"charges": [],
						"chargesTotal": null,
						"perPaxTotal": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "30000000"
						  },
						  "total": "60000000"
						},
						"total": {
						  "currency": "IRR",
						  "paxPrice": {
							"ADT": "60000000"
						  },
						  "total": "60000000"
						},
						"commissionTotal": null,
						"commissionOrTotal": null,
						"agentCommissionTotal": null,
						"discountTotal": null,
						"extraSeatAmountTotal": {},
						"currency": "IRR",
						"selectedCurrencyTotal": null,
						"ancillaryBundle": {
						  "103091": false
						},
						"pqLight": {
						  "fares": [
							{
							  "fareId": 50814,
							  "fareRuleID": 6203,
							  "fareSeq": 1,
							  "amount": {
								"ADT": "23270000"
							  },
							  "segmentSpan": [
								{
								  "fltSegID": 103091,
								  "cabinClassID": 2,
								  "rbdID": 20,
								  "rbdInvType": "STANDARD",
								  "invFltSegCCID": 177228,
								  "actualDirection": "OUTBOUND",
								  "flightID": 100408,
								  "segmentCode": "IKA-NJF",
								  "ownerAirlineID": 15,
								  "puGroupID": -3801220797509906400,
								  "flightType": "INT",
								  "flightOperationType": "STANDARD",
								  "overBookingOnDisruption": false,
								  "addonSegment": false
								}
							  ],
							  "segmentSpanCode": "IKA-NJF"
							}
						  ],
						  "fltSegs": {
							"103091": {
							  "fltSegID": 103091,
							  "cabinClassID": 2,
							  "rbdID": 20,
							  "rbdInvType": "STANDARD",
							  "invFltSegCCID": 177228,
							  "actualDirection": "OUTBOUND",
							  "flightID": 100408,
							  "segmentCode": "IKA-NJF",
							  "ownerAirlineID": 15,
							  "puGroupID": -3801220797509906400,
							  "flightType": "INT",
							  "flightOperationType": "STANDARD",
							  "overBookingOnDisruption": false,
							  "addonSegment": false
							}
						  },
						  "inventoryConsumption": {
							"adtConsumption": 2,
							"infConsumption": 0
						  },
						  "bookingCurrencyCode": "IRR",
						  "searchUID": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0"
						},
						"segmentFareAmount": {
						  "ADT": {
							"103091": "23270000"
						  }
						},
						"segmentChargeAmount": {},
						"segmentTaxAmount": {
						  "ADT": {
							"103091": "6730000"
						  }
						},
						"segmentFareChargeAmount": {
						  "ADT": {
							"103091": "30000000"
						  }
						},
						"openReturnBooking": false,
						"openReturnMaxDate": null,
						"openReturnMaxDateStr": null,
						"agentDiscountApplied": false,
						"addToDisruptionOverbookQueue": false,
						"searchUID": "a0952fcd-0fcf-4390-a5bc-a6c331a8dad0"
					  }
					}
				  ],
				  "flightSegments": [
					{
					  "flightID": 100408,
					  "fltSegID": 103091,
					  "flightScheduleID": 46989,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5329",
					  "direction": "OUTBOUND",
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A321-200",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1663380000000,
					  "arrivalDateUCT": 1663385700000,
					  "departureDateLocal": 1663396200000,
					  "arrivalDateLocal": 1663396500000,
					  "departureDateLocalStr": "17/09/2022 06:30:00",
					  "arrivalDateLocalStr": "17/09/2022 06:35:00",
					  "departureDateUTCStr": "17/09/2022 02:00:00",
					  "arrivalDateUTCStr": "17/09/2022 03:35:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": true,
					  "flightFull": false,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 6,
					  "legINFCount": 9,
					  "actualDirection": "OUTBOUND",
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				},
				{
				  "seatAvailability": false,
				  "seatFull": false,
				  "fareProductDetails": [],
				  "flightSegments": [
					{
					  "flightID": 101260,
					  "fltSegID": 103943,
					  "flightScheduleID": 46975,
					  "segmentCode": "IKA-NJF",
					  "operatedCarrier": "IR",
					  "operatedAirlineID": 15,
					  "ownerAirlineID": 15,
					  "flightDesignator": "IR5333",
					  "direction": null,
					  "flightType": "INT",
					  "flightOperationType": "STANDARD",
					  "aircraftModel": "A321-200",
					  "departureCode": "IKA",
					  "arrivalCode": "NJF",
					  "depOJAirports": null,
					  "arrOJAirports": null,
					  "stopCount": 0,
					  "departureDateUTC": 1663398000000,
					  "arrivalDateUCT": 1663403700000,
					  "departureDateLocal": 1663414200000,
					  "arrivalDateLocal": 1663414500000,
					  "departureDateLocalStr": "17/09/2022 11:30:00",
					  "arrivalDateLocalStr": "17/09/2022 11:35:00",
					  "departureDateUTCStr": "17/09/2022 07:00:00",
					  "arrivalDateUTCStr": "17/09/2022 08:35:00",
					  "duration": {
						"standardMinutes": 95,
						"standardSeconds": 5700,
						"standardHours": 1,
						"standardDays": 0,
						"millis": 5700000
					  },
					  "depFromTransit": false,
					  "depFromStopover": false,
					  "status": "OPN",
					  "segmentAvailability": false,
					  "flightFull": false,
					  "eligibleForMarriedSegment": false,
					  "addOnFlight": false,
					  "internationalFlightOfAddon": false,
					  "addonInternationalFltSeg": null,
					  "restrictedRBDs": {},
					  "fltSegBeforeSalesDate": false,
					  "departureDateFromStopLocal": null,
					  "arrivalDateToStopLocal": null,
					  "departureDateFromStopLocalStr": null,
					  "arrivalDateToStopLocalStr": null,
					  "stopAirport": null,
					  "availableRBD": [],
					  "legADTCount": 1,
					  "legINFCount": 10,
					  "actualDirection": null,
					  "durationStr": "01h:35m",
					  "internationalFlight": true,
					  "flownSegment": false
					}
				  ]
				}
			  ],
			  "daySummaryDTOMap": {
				"0": {
				  "date": 1663372800000,
				  "amount": "60000000",
				  "currencyCode": "IRR"
				}
			  },
			  "daySeatFullMap": {
				"0": false
			  },
			  "journeyAvailabilityMap": {
				"0": true
			  }
			}
		  ],
		  "priceQuote": null,
		  "salesDate": "05/09/2022 15:08:00:000"
		}
	  }
	`
)
