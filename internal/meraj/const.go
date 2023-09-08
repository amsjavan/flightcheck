package meraj

const (
	Success1 = `
	{
		"Result": [
		  {
			"FlightGroups": [
			  {
				"Flights": [
				  {
					"Id": 601018091,
					"SegmentGroups": [
					  {
						"Index": 0,
						"Segments": [
						  {
							"Id": 0,
							"Capacity": 2,
							"ArrivalDateTime": "2022-09-10T22:00:00Z",
							"DepartureDateTime": "2022-09-10T22:00:00Z",
							"Airline": {
							  "Id": 333,
							  "SmallLogoImage": {
								"Id": 76257,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "85c1f49b-00f5-4962-a80a-75132788bd71",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=76257&password=85c1f49b-00f5-4962-a80a-75132788bd71&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "LogoImage": {
								"Id": 76256,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "5a7107f0-4a8d-461b-842e-3912a10ef895",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=76256&password=5a7107f0-4a8d-461b-842e-3912a10ef895&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "AircraftImage": {
								"Id": 438169,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "528cd0fa-144d-4105-993b-f86e83d63a6c",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=438169&password=528cd0fa-144d-4105-993b-f86e83d63a6c&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "Name": "Meraj Air",
							  "PersianName": "معراج",
							  "IATA": "JI",
							  "ICAO": "MRJ"
							},
							"Aircraft": {
							  "Name": "320",
							  "PersianName": null,
							  "CabinName": null
							},
							"Terminal": "پروازهای خارجی",
							"Destination": {
							  "CityId": 1209,
							  "CountryId": 2067,
							  "PersianCityName": "نجف",
							  "EnglishCityName": "Najaf",
							  "ArabicCityName": "نجف",
							  "PersianAirportName": "نجف الاشرف",
							  "EnglishAirportName": "Al Najaf International Airport",
							  "ArabicAirportName": "مطار النجف الدولی",
							  "Code": "NJF"
							},
							"Origin": {
							  "PersianCityName": "تهران",
							  "EnglishCityName": "Tehran",
							  "ArabicCityName": "تهران",
							  "PersianAirportName": "بین المللی امام خمینی",
							  "EnglishAirportName": "Tehran Imam Khomeini International Airport",
							  "ArabicAirportName": "مطار الإمام الخمینی الدولی فی طهران",
							  "Code": "IKA"
							},
							"FareBasisCode": "C",
							"CabinTypeName": "Business",
							"ReservationBookingDesignator": "CC99",
							"FlightNumber": "4853",
							"CancellationPolicies": null,
							"AllowedLuggage": null,
							"ArrivalTerminalName": null,
							"DepartureTerminalName": "پروازهای خارجی",
							"FromCityName": "Tehran",
							"FromCountryName": "Iran",
							"PersianFromCityName": "تهران",
							"PersianFromCountryName": "ایران",
							"ArabicFromCityName": "تهران",
							"ArabicFromCountryName": "ایران",
							"ToCityName": "Najaf",
							"ToCountryName": "Iraq",
							"PersianToCityName": "نجف",
							"PersianToCountryName": "عراق",
							"ArabicToCityName": "نجف",
							"ArabicToCountryName": "عراق",
							"TotalFlightTime": "01:30:00",
							"Passengers": null,
							"UniqueIndex": 0,
							"Supplier": null,
							"IsFromAirline": true
						  }
						]
					  }
					],
					"Prices": [
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  }
					],
					"PassengerPrices": [
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  }
					],
					"Taxes": [],
					"CancellationPolicies": [
					  {
						"Title": null,
						"PercentText": null,
						"FromHour": null,
						"FromDay": null,
						"FromFormulaType": 1,
						"ToHour": null,
						"ToDay": null,
						"ToFormulaType": 1,
						"IsCurrentPenalty": false
					  }
					],
					"TicketFiles": null,
					"InvoiceFiles": null,
					"ServiceType": 6,
					"IsCharter": true,
					"IsInternational": true,
					"IsRefundable": false,
					"Key": "MerajairSepehrMata24",
					"Credit": 0,
					"Priority": 0,
					"FlightUniqueGroupNumber": "1_2022-09-10T22:00_IKA_NJF_333_4853_Business_True_C",
					"ContactNumber": "05132222223",
					"IsFromAirline": true
				  }
				]
			  }
			],
			"Flights": null,
			"IsMerged": false,
			"TripType": 0,
			"UniqueIndex": 0,
			"IsResetCacheHappened": false
		  }
		],
		"Success": true,
		"Key": null,
		"FailedReason": 0,
		"Error": null,
		"Validations": [],
		"ProjectsTree": null
	  }
	`
	Success2 = `
	{
		"Result": [
		  {
			"FlightGroups": [
			  {
				"Flights": [
				  {
					"Id": 600981377,
					"SegmentGroups": [
					  {
						"Index": 0,
						"Segments": [
						  {
							"Id": 0,
							"Capacity": 3,
							"ArrivalDateTime": "2022-09-14T17:30:00Z",
							"DepartureDateTime": "2022-09-14T17:30:00Z",
							"Airline": {
							  "Id": 333,
							  "SmallLogoImage": {
								"Id": 76257,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "85c1f49b-00f5-4962-a80a-75132788bd71",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=76257&password=85c1f49b-00f5-4962-a80a-75132788bd71&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "LogoImage": {
								"Id": 76256,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "5a7107f0-4a8d-461b-842e-3912a10ef895",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=76256&password=5a7107f0-4a8d-461b-842e-3912a10ef895&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "AircraftImage": {
								"Id": 438169,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "528cd0fa-144d-4105-993b-f86e83d63a6c",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=438169&password=528cd0fa-144d-4105-993b-f86e83d63a6c&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "Name": "Meraj Air",
							  "PersianName": "معراج",
							  "IATA": "JI",
							  "ICAO": "MRJ"
							},
							"Aircraft": {
							  "Name": "320",
							  "PersianName": null,
							  "CabinName": null
							},
							"Terminal": "پروازهای خارجی",
							"Destination": {
							  "CityId": 1209,
							  "CountryId": 2067,
							  "PersianCityName": "نجف",
							  "EnglishCityName": "Najaf",
							  "ArabicCityName": "نجف",
							  "PersianAirportName": "نجف الاشرف",
							  "EnglishAirportName": "Al Najaf International Airport",
							  "ArabicAirportName": "مطار النجف الدولی",
							  "Code": "NJF"
							},
							"Origin": {
							  "PersianCityName": "تهران",
							  "EnglishCityName": "Tehran",
							  "ArabicCityName": "تهران",
							  "PersianAirportName": "بین المللی امام خمینی",
							  "EnglishAirportName": "Tehran Imam Khomeini International Airport",
							  "ArabicAirportName": "مطار الإمام الخمینی الدولی فی طهران",
							  "Code": "IKA"
							},
							"FareBasisCode": "C",
							"CabinTypeName": "Business",
							"ReservationBookingDesignator": "CC99",
							"FlightNumber": "4845",
							"CancellationPolicies": null,
							"AllowedLuggage": null,
							"ArrivalTerminalName": null,
							"DepartureTerminalName": "پروازهای خارجی",
							"FromCityName": "Tehran",
							"FromCountryName": "Iran",
							"PersianFromCityName": "تهران",
							"PersianFromCountryName": "ایران",
							"ArabicFromCityName": "تهران",
							"ArabicFromCountryName": "ایران",
							"ToCityName": "Najaf",
							"ToCountryName": "Iraq",
							"PersianToCityName": "نجف",
							"PersianToCountryName": "عراق",
							"ArabicToCityName": "نجف",
							"ArabicToCountryName": "عراق",
							"TotalFlightTime": "01:30:00",
							"Passengers": null,
							"UniqueIndex": 0,
							"Supplier": null,
							"IsFromAirline": true
						  }
						]
					  }
					],
					"Prices": [
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  }
					],
					"PassengerPrices": [
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  }
					],
					"Taxes": [],
					"CancellationPolicies": [
					  {
						"Title": null,
						"PercentText": null,
						"FromHour": null,
						"FromDay": null,
						"FromFormulaType": 1,
						"ToHour": null,
						"ToDay": null,
						"ToFormulaType": 1,
						"IsCurrentPenalty": false
					  }
					],
					"TicketFiles": null,
					"InvoiceFiles": null,
					"ServiceType": 6,
					"IsCharter": true,
					"IsInternational": true,
					"IsRefundable": false,
					"Key": "MerajairSepehrMata24",
					"Credit": 0,
					"Priority": 0,
					"FlightUniqueGroupNumber": "1_2022-09-14T17:30_IKA_NJF_333_4845_Business_True_C",
					"ContactNumber": "05132222223",
					"IsFromAirline": true
				  }
				]
			  },
			  {
				"Flights": [
				  {
					"Id": 601022918,
					"SegmentGroups": [
					  {
						"Index": 0,
						"Segments": [
						  {
							"Id": 0,
							"Capacity": 8,
							"ArrivalDateTime": "2022-09-14T22:00:00Z",
							"DepartureDateTime": "2022-09-14T22:00:00Z",
							"Airline": {
							  "Id": 333,
							  "SmallLogoImage": {
								"Id": 76257,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "85c1f49b-00f5-4962-a80a-75132788bd71",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=76257&password=85c1f49b-00f5-4962-a80a-75132788bd71&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "LogoImage": {
								"Id": 76256,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "5a7107f0-4a8d-461b-842e-3912a10ef895",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=76256&password=5a7107f0-4a8d-461b-842e-3912a10ef895&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "AircraftImage": {
								"Id": 438169,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "528cd0fa-144d-4105-993b-f86e83d63a6c",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=438169&password=528cd0fa-144d-4105-993b-f86e83d63a6c&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "Name": "Meraj Air",
							  "PersianName": "معراج",
							  "IATA": "JI",
							  "ICAO": "MRJ"
							},
							"Aircraft": {
							  "Name": "320",
							  "PersianName": null,
							  "CabinName": null
							},
							"Terminal": "پروازهای خارجی",
							"Destination": {
							  "CityId": 1209,
							  "CountryId": 2067,
							  "PersianCityName": "نجف",
							  "EnglishCityName": "Najaf",
							  "ArabicCityName": "نجف",
							  "PersianAirportName": "نجف الاشرف",
							  "EnglishAirportName": "Al Najaf International Airport",
							  "ArabicAirportName": "مطار النجف الدولی",
							  "Code": "NJF"
							},
							"Origin": {
							  "PersianCityName": "تهران",
							  "EnglishCityName": "Tehran",
							  "ArabicCityName": "تهران",
							  "PersianAirportName": "بین المللی امام خمینی",
							  "EnglishAirportName": "Tehran Imam Khomeini International Airport",
							  "ArabicAirportName": "مطار الإمام الخمینی الدولی فی طهران",
							  "Code": "IKA"
							},
							"FareBasisCode": "W",
							"CabinTypeName": "Economy",
							"ReservationBookingDesignator": "WAT1",
							"FlightNumber": "4853",
							"CancellationPolicies": null,
							"AllowedLuggage": null,
							"ArrivalTerminalName": null,
							"DepartureTerminalName": "پروازهای خارجی",
							"FromCityName": "Tehran",
							"FromCountryName": "Iran",
							"PersianFromCityName": "تهران",
							"PersianFromCountryName": "ایران",
							"ArabicFromCityName": "تهران",
							"ArabicFromCountryName": "ایران",
							"ToCityName": "Najaf",
							"ToCountryName": "Iraq",
							"PersianToCityName": "نجف",
							"PersianToCountryName": "عراق",
							"ArabicToCityName": "نجف",
							"ArabicToCountryName": "عراق",
							"TotalFlightTime": "01:30:00",
							"Passengers": null,
							"UniqueIndex": 0,
							"Supplier": null,
							"IsFromAirline": true
						  }
						]
					  }
					],
					"Prices": [
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 5000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 5000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 5000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 5000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 5000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 5000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 5000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 5000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  }
					],
					"PassengerPrices": [
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 30000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 30000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 5000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 5000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 5000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 5000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 5000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 5000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 5000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 5000000,
						"IncomePaymentAmount": 0
					  }
					],
					"Taxes": [],
					"CancellationPolicies": [
					  {
						"Title": null,
						"PercentText": null,
						"FromHour": null,
						"FromDay": null,
						"FromFormulaType": 1,
						"ToHour": null,
						"ToDay": null,
						"ToFormulaType": 1,
						"IsCurrentPenalty": false
					  }
					],
					"TicketFiles": null,
					"InvoiceFiles": null,
					"ServiceType": 6,
					"IsCharter": true,
					"IsInternational": true,
					"IsRefundable": false,
					"Key": "MerajairSepehrMata24",
					"Credit": 0,
					"Priority": 0,
					"FlightUniqueGroupNumber": "1_2022-09-14T22:00_IKA_NJF_333_4853_Economy_True_W",
					"ContactNumber": "05132222223",
					"IsFromAirline": true
				  }
				]
			  },
			  {
				"Flights": [
				  {
					"Id": 600981392,
					"SegmentGroups": [
					  {
						"Index": 0,
						"Segments": [
						  {
							"Id": 0,
							"Capacity": 2,
							"ArrivalDateTime": "2022-09-14T22:00:00Z",
							"DepartureDateTime": "2022-09-14T22:00:00Z",
							"Airline": {
							  "Id": 333,
							  "SmallLogoImage": {
								"Id": 76257,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "85c1f49b-00f5-4962-a80a-75132788bd71",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=76257&password=85c1f49b-00f5-4962-a80a-75132788bd71&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "LogoImage": {
								"Id": 76256,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "5a7107f0-4a8d-461b-842e-3912a10ef895",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=76256&password=5a7107f0-4a8d-461b-842e-3912a10ef895&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "AircraftImage": {
								"Id": 438169,
								"CreatedDateTime": "0001-01-01T00:00:00Z",
								"Password": "528cd0fa-144d-4105-993b-f86e83d63a6c",
								"Url": "https://cdn.utravs.com/Storage/DownloadFile?fileId=438169&password=528cd0fa-144d-4105-993b-f86e83d63a6c&Tag=",
								"LanguageCode": null,
								"Tag": null
							  },
							  "Name": "Meraj Air",
							  "PersianName": "معراج",
							  "IATA": "JI",
							  "ICAO": "MRJ"
							},
							"Aircraft": {
							  "Name": "320",
							  "PersianName": null,
							  "CabinName": null
							},
							"Terminal": "پروازهای خارجی",
							"Destination": {
							  "CityId": 1209,
							  "CountryId": 2067,
							  "PersianCityName": "نجف",
							  "EnglishCityName": "Najaf",
							  "ArabicCityName": "نجف",
							  "PersianAirportName": "نجف الاشرف",
							  "EnglishAirportName": "Al Najaf International Airport",
							  "ArabicAirportName": "مطار النجف الدولی",
							  "Code": "NJF"
							},
							"Origin": {
							  "PersianCityName": "تهران",
							  "EnglishCityName": "Tehran",
							  "ArabicCityName": "تهران",
							  "PersianAirportName": "بین المللی امام خمینی",
							  "EnglishAirportName": "Tehran Imam Khomeini International Airport",
							  "ArabicAirportName": "مطار الإمام الخمینی الدولی فی طهران",
							  "Code": "IKA"
							},
							"FareBasisCode": "C",
							"CabinTypeName": "Business",
							"ReservationBookingDesignator": "CC99",
							"FlightNumber": "4853",
							"CancellationPolicies": null,
							"AllowedLuggage": null,
							"ArrivalTerminalName": null,
							"DepartureTerminalName": "پروازهای خارجی",
							"FromCityName": "Tehran",
							"FromCountryName": "Iran",
							"PersianFromCityName": "تهران",
							"PersianFromCountryName": "ایران",
							"ArabicFromCityName": "تهران",
							"ArabicFromCountryName": "ایران",
							"ToCityName": "Najaf",
							"ToCountryName": "Iraq",
							"PersianToCityName": "نجف",
							"PersianToCountryName": "عراق",
							"ArabicToCityName": "نجف",
							"ArabicToCountryName": "عراق",
							"TotalFlightTime": "01:30:00",
							"Passengers": null,
							"UniqueIndex": 0,
							"Supplier": null,
							"IsFromAirline": true
						  }
						]
					  }
					],
					"Prices": [
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 0,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  }
					],
					"PassengerPrices": [
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 50000000,
						"PassengerType": 1,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 50000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 49000000,
						"PassengerType": 2,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 49000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 2,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 1,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 3,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 10,
						"CurrencyCode": 364,
						"Amount": 0,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 0,
						"IncomePaymentAmount": 0
					  },
					  {
						"PriceType": 9,
						"CurrencyCode": 364,
						"Amount": 6000000,
						"PassengerType": 3,
						"ClassType": null,
						"WithoutIncomePaymentAmount": 6000000,
						"IncomePaymentAmount": 0
					  }
					],
					"Taxes": [],
					"CancellationPolicies": [
					  {
						"Title": null,
						"PercentText": null,
						"FromHour": null,
						"FromDay": null,
						"FromFormulaType": 1,
						"ToHour": null,
						"ToDay": null,
						"ToFormulaType": 1,
						"IsCurrentPenalty": false
					  }
					],
					"TicketFiles": null,
					"InvoiceFiles": null,
					"ServiceType": 6,
					"IsCharter": true,
					"IsInternational": true,
					"IsRefundable": false,
					"Key": "MerajairSepehrMata24",
					"Credit": 0,
					"Priority": 0,
					"FlightUniqueGroupNumber": "1_2022-09-14T22:00_IKA_NJF_333_4853_Business_True_C",
					"ContactNumber": "05132222223",
					"IsFromAirline": true
				  }
				]
			  }
			],
			"Flights": null,
			"IsMerged": false,
			"TripType": 0,
			"UniqueIndex": 0,
			"IsResetCacheHappened": false
		  }
		],
		"Success": true,
		"Key": null,
		"FailedReason": 0,
		"Error": null,
		"Validations": [],
		"ProjectsTree": null
	  }
	`
	Failed = `
	{
		"Result": [
		  {
			"FlightGroups": [],
			"Flights": null,
			"IsMerged": false,
			"TripType": 0,
			"UniqueIndex": 0,
			"IsResetCacheHappened": false
		  }
		],
		"Success": true,
		"Key": null,
		"FailedReason": 0,
		"Error": null,
		"Validations": [],
		"ProjectsTree": null
	  }
	`
)
