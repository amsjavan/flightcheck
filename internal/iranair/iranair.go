package iranair

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flightcheck/internal/pkg"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var isAvailabe bool
var dateToDay = map[string]string{
	"24/08/2023": "پنج شنبه",
	"25/08/2023": "جمعه",
	"26/08/2023": "شنبه",
	"27/08/2023": "یکشنبه",
	"28/08/2023": "دوشنبه",
	"29/08/2023": "سه شنبه",
	"30/08/2023": "چهارشنبه",
	"31/08/2023": "پنج شنبه",
}

var miladiToHejri = map[string]string{
	"24/08/2023": "1402/06/02",
	"25/08/2023": "1402/06/03",
	"26/08/2023": "1402/06/04",
	"27/08/2023": "1402/06/05",
	"28/08/2023": "1402/06/06",
	"29/08/2023": "1402/06/07",
	"30/08/2023": "1402/06/08",
	"31/08/2023": "1402/06/09",
	"01/09/2023": "1402/06/10",
	"02/09/2023": "1402/06/11",
	"03/09/2023": "1402/06/12",
	"04/09/2023": "1402/06/13",
	"05/09/2023": "1402/06/14",
	"06/09/2023": "1402/06/15",
	"07/09/2023": "1402/06/16",
	"08/09/2023": "1402/06/17",
	"09/09/2023": "1402/06/18",
	"10/09/2023": "1402/06/19",
	"11/09/2023": "1402/06/20",
	"12/09/2023": "1402/06/21",
	"13/09/2023": "1402/06/22",
}

var toTehranDates = []DateModel{
	{date: "31/08/2023"},
	{date: "01/09/2023"},
	{date: "02/09/2023"},
	{date: "03/09/2023"},
	{date: "04/09/2023"},
	{date: "05/09/2023"},
	{date: "06/09/2023"},
	{date: "07/09/2023"},
	{date: "08/09/2023"},
	{date: "09/09/2023"},
	{date: "10/09/2023"},
	{date: "11/09/2023"},
	{date: "12/09/2023"},
	{date: "13/09/2023"},
}

var toNajafDates = []DateModel{
	{date: "24/08/2023"},
	{date: "25/08/2023"},
	{date: "26/08/2023"},
	{date: "27/08/2023"},
	{date: "28/08/2023"},
	{date: "29/08/2023"},
	{date: "30/08/2023"},
	{date: "31/08/2023"},
	{date: "01/09/2023"},
	{date: "02/09/2023"},
	{date: "03/09/2023"},
	{date: "04/09/2023"},
}

const (
	TEHRAN  = "IKA"
	MASHHAD = "MHD"
)

type DateModel struct {
	date        string
	isAvailable bool
}

func Run() {
	defer func() {
		if r := recover(); r != nil {
			msg := "stacktrace from panic in iranair: \n" + string(debug.Stack())
			fmt.Println(msg)
			pkg.SendMessage(msg)
		}
	}()
	for {
		runTehranToNajaf()
		// runNajafToTehran()

		// runMashhadToNajaf()
		// runNajafToMashhad()
		time.Sleep(5 * time.Minute)
	}
}

// Tehran
func runTehranToNajaf() {
	for i := 0; i < len(toNajafDates); i++ {
		time.Sleep(20 * time.Second)
		fmt.Printf("check iranair to najaf flight %s ...\n", toNajafDates[i].date)
		rsp, err := checkTehranToNajafFlight(toNajafDates[i].date)
		if err != nil {
			pkg.SendMessage("error in ckeck iranair to najaf flight: " + err.Error())
			println("error in ckeck iranair flight:", err.Error())
			continue
		}
		if !rsp.Success {
			continue
		}
		if checkAvailability(rsp) && !toNajafDates[i].isAvailable {
			toNajafDates[i].isAvailable = true
			pkg.SendToGroup(
				fmt.Sprintf("پرواز ایران ایر %v در تاریخ %v از تهران به نجف پیدا شد %v",
					dateToDay[toNajafDates[i].date],
					miladiToHejri[toNajafDates[i].date],
					"https://ebooking.iranair.com"),
			)
		}

		if !checkAvailability(rsp) && toNajafDates[i].isAvailable {
			toNajafDates[i].isAvailable = false
			pkg.SendToGroup(
				fmt.Sprintf("پرواز ایران ایر %v در تاریخ %v از تهران به نجف تمام شد",
					dateToDay[toNajafDates[i].date],
					miladiToHejri[toNajafDates[i].date]),
			)
		}
	}

}

func runNajafToTehran() {
	for i := 0; i < len(toTehranDates); i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("check iranair to tehran flight %s ...\n", toTehranDates[i].date)
		rsp, err := checkNajafToTehranFlight(toTehranDates[i].date)
		if err != nil {
			pkg.SendMessage("error in ckeck iranair flight: " + err.Error())
			println("error in ckeck iranair flight:", err.Error())
			continue
		}
		if !rsp.Success {
			continue
		}
		if checkAvailability(rsp) && !toTehranDates[i].isAvailable {
			toTehranDates[i].isAvailable = true
			pkg.SendToGroup(
				fmt.Sprintf("پرواز ایران ایر %v در تاریخ %v به تهران پیدا شد %v",
					dateToDay[toTehranDates[i].date],
					miladiToHejri[toTehranDates[i].date],
					"https://ebooking.iranair.com"),
			)
		}

		if !checkAvailability(rsp) && toTehranDates[i].isAvailable {
			toTehranDates[i].isAvailable = false
			pkg.SendToGroup(
				fmt.Sprintf("پرواز ایران ایر %v در تاریخ %v به تهران تمام شد",
					dateToDay[toTehranDates[i].date],
					miladiToHejri[toTehranDates[i].date]),
			)
		}
	}
}

// Mashhad
func runMashhadToNajaf() {
	for i := 0; i < len(toNajafDates); i++ {
		time.Sleep(20 * time.Second)
		fmt.Printf("check iranair to najaf flight %s ...\n", toNajafDates[i].date)
		rsp, err := checkMashhadToNajafFlight(toNajafDates[i].date)
		if err != nil {
			pkg.SendMessage("error in ckeck iranair to najaf flight: " + err.Error())
			println("error in ckeck iranair flight:", err.Error())
			continue
		}
		if !rsp.Success {
			continue
		}
		if checkAvailability(rsp) && !toNajafDates[i].isAvailable {
			toNajafDates[i].isAvailable = true
			pkg.SendToGroup(
				fmt.Sprintf("پرواز ایران ایر %v در تاریخ %v از مشهد به نجف پیدا شد %v",
					dateToDay[toNajafDates[i].date],
					miladiToHejri[toNajafDates[i].date],
					"https://ebooking.iranair.com"),
			)
		}

		if !checkAvailability(rsp) && toNajafDates[i].isAvailable {
			toNajafDates[i].isAvailable = false
			pkg.SendToGroup(
				fmt.Sprintf("پرواز ایران ایر %v در تاریخ %v از مشهد به نجف تمام شد",
					dateToDay[toNajafDates[i].date],
					miladiToHejri[toNajafDates[i].date]),
			)
		}
	}

}

func runNajafToMashhad() {
	for i := 0; i < len(toTehranDates); i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("check iranair to tehran flight %s ...\n", toTehranDates[i].date)
		rsp, err := checkNajafToMashhadFlight(toTehranDates[i].date)
		if err != nil {
			pkg.SendMessage("error in ckeck iranair flight: " + err.Error())
			println("error in ckeck iranair flight:", err.Error())
			continue
		}
		if !rsp.Success {
			continue
		}
		if checkAvailability(rsp) && !toTehranDates[i].isAvailable {
			toTehranDates[i].isAvailable = true
			pkg.SendToGroup(
				fmt.Sprintf("پرواز ایران ایر %v در تاریخ %v به مشهد پیدا شد %v",
					dateToDay[toTehranDates[i].date],
					miladiToHejri[toTehranDates[i].date],
					"https://ebooking.iranair.com"),
			)
		}

		if !checkAvailability(rsp) && toTehranDates[i].isAvailable {
			toTehranDates[i].isAvailable = false
			pkg.SendToGroup(
				fmt.Sprintf("پرواز ایران ایر %v در تاریخ %v به مشهد تمام شد",
					dateToDay[toTehranDates[i].date],
					miladiToHejri[toTehranDates[i].date]),
			)
		}
	}
}

func checkAvailability(rsp *Response) bool {
	for _, dto := range rsp.Result.Journeies[0].DTOs {
		if dto.SeatAvailability {
			return true
		}
	}
	return false
}

func checkNajafToTehranFlight(date string) (*Response, error) {
	return checkFlight(getToTehranInputJson(date, TEHRAN), "")
}

func checkTehranToNajafFlight(date string) (*Response, error) {
	return checkFlight(getToNajafInputJson(date, TEHRAN), "")
}

func checkNajafToMashhadFlight(date string) (*Response, error) {
	return checkFlight(getToTehranInputJson(date, MASHHAD), "")
}

func checkMashhadToNajafFlight(date string) (*Response, error) {
	return checkFlight(getToNajafInputJson(date, MASHHAD), "")
}

func checkFlight(input, empty string) (*Response, error) {
	url := "https://ebooking.iranair.com/ibe/IR/availability/search?captchaInput="

	var jsonStr = []byte(input)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.Wrap(err, "error in creating request")
	}
	req.Header.Set("authority", "ebooking.iranair.com")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	println("sending iranair request...")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error in sending request")
	}
	println("request iranair sent...")
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error in read body")
	}
	res := Response{}
	json.Unmarshal(body, &res)
	if !res.Success {
		println("bad response! iranair body: ", string(body))
	}
	return &res, nil
}

func getToNajafInputJson(date, city string) string {
	rsp := `
	{
		"rand": "2023-08-18T18:16:23+03:30",
		"uuid": "",
		"travelPaxInfo": {
		  "adultCount": 3,
		  "infantCount": 0,
		  "paxQuantities": {
			"ADT": "2",
			"CHD": "1"
		  }
		},
		"availabilitySearchType": "OW",
		"pqPref": {},
		"searchedOnDInfos": [
		  {
			"sequence": 1,
			"origin": {
			  "airportCodes": [
				"IKA"
			  ],
			  "code": "IKA",
			  "isAirport": true
			},
			"destination": {
			  "airportCodes": [
				"NJF"
			  ],
			  "code": "NJF",
			  "isAirport": true
			},
			"selectedDepDateStr": "$date",
			"depDayVariance": "SameDay",
			"cabinClassCode": null
		  }
		],
		"depOrigin": {
		  "code": "IKA",
		  "airportCodes": [
			"IKA"
		  ],
		  "isAirport": true
		},
		"searchBehaviour": {
		  "quotePriceForAllJourneys": true,
		  "sameFareProductPerOnd": true
		},
		"anonymousUserDTO": {
		  "carrierCode": "IR",
		  "carrierID": 15,
		  "selectedLanguage": "fa",
		  "salesChannel": "WEB"
		},
		"searchRequestContext": {
		  "modifyBookingContext": false
		}
	  }
	`

	tmp := strings.ReplaceAll(rsp, "$date", date)
	// result := strings.ReplaceAll(tmp, "$city", city)
	// println("=========", result)
	return tmp

}

func getToTehranInputJson(date, city string) string {
	rsp := `
	{
		"rand": "2023-08-18T13:03:35+03:30",
		"uuid": "",
		"travelPaxInfo": {
		  "adultCount": 3,
		  "infantCount": 0,
		  "paxQuantities": {
			"ADT": "2",
			"CHD": "1"
		  }
		},
		"availabilitySearchType": "OW",
		"pqPref": {},
		"searchedOnDInfos": [
		  {
			"sequence": 1,
			"origin": {
			  "airportCodes": [
				"NJF"
			  ],
			  "code": "NJF",
			  "isAirport": true
			},
			"destination": {
			  "airportCodes": [
				"$city"
			  ],
			  "code": "$city",
			  "isAirport": true
			},
			"selectedDepDateStr": "$date",
			"depDayVariance": "SameDay",
			"cabinClassCode": null
		  }
		],
		"depOrigin": {
		  "code": "NJF",
		  "airportCodes": [
			"NJF"
		  ],
		  "isAirport": true
		},
		"searchBehaviour": {
		  "quotePriceForAllJourneys": true,
		  "sameFareProductPerOnd": true
		},
		"anonymousUserDTO": {
		  "carrierCode": "IR",
		  "carrierID": 15,
		  "selectedLanguage": "fa",
		  "salesChannel": "WEB"
		},
		"searchRequestContext": {
		  "modifyBookingContext": false
		}
	  }
	`

	tmp := strings.ReplaceAll(rsp, "$date", date)
	result := strings.ReplaceAll(tmp, "$city", city)
	println("==========", result)
	return result

}
