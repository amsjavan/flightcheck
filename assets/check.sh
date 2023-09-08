#!/bin/bash

curl 'https://ebooking.iranair.com/ibe/IR/availability/search?captchaInput=' \
  -H 'authority: ebooking.iranair.com' \
  -H 'accept: application/json, text/javascript, */*; q=0.01' \
  -H 'accept-language: en-US,en;q=0.9,fa;q=0.8' \
  -H 'content-type: application/json' \
  -H 'cookie: JSESSIONID=APP04~mRF-EUz_lmHXGFY0r97TwOny_nDRC1PVlkfq1jvC.server01d' \
  -H 'origin: https://ebooking.iranair.com' \
  -H 'referer: https://ebooking.iranair.com/ibe/IR/home?language=fa' \
  -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36' \
  -H 'x-requested-with: XMLHttpRequest' \
  --data-raw '{"rand":"2022-09-05T19:45:21+04:30","uuid":"","travelPaxInfo":{"adultCount":2,"infantCount":0,"paxQuantities":{"ADT":"2"}},"availabilitySearchType":"OW","pqPref":{},"searchedOnDInfos":[{"sequence":1,"origin":{"airportCodes":["IKA"],"code":"IKA","isAirport":true},"destination":{"airportCodes":["NJF"],"code":"NJF","isAirport":true},"selectedDepDateStr":"11/09/2022","depDayVariance":"SameDay","cabinClassCode":null}],"depOrigin":{"code":"IKA","airportCodes":["IKA"],"isAirport":true},"searchBehaviour":{"quotePriceForAllJourneys":true,"sameFareProductPerOnd":true},"anonymousUserDTO":{"carrierCode":"IR","carrierID":15,"selectedLanguage":"fa","salesChannel":"WEB"},"searchRequestContext":{"modifyBookingContext":false}}' \
  --compressed