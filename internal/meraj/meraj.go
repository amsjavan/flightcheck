package meraj

import (
	"bytes"
	"encoding/json"
	"flightcheck/internal/pkg"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var isAvailabe bool
var dateToDay = map[string]string{
	"2022-09-11T00:00:00.000Z": "یکشنبه",
	"2022-09-12T00:00:00.000Z": "دوشنبه",
	"2022-09-13T00:00:00.000Z": "سه شنبه",
	"2022-09-14T00:00:00.000Z": "چهارشنبه",
	"2022-09-15T00:00:00.000Z": "پنج شنبه",
	"2022-09-16T00:00:00.000Z": "جمعه",
	"2022-09-17T00:00:00.000Z": "شنبه",
	"2022-09-18T00:00:00.000Z": "یکشنبه",
}

var toTehranDates = []DateModel{
	{date: "2022-09-11T00:00:00.000Z"},
	{date: "2022-09-12T00:00:00.000Z"},
	{date: "2022-09-13T00:00:00.000Z"},
	{date: "2022-09-14T00:00:00.000Z"},
	{date: "2022-09-15T00:00:00.000Z"},
	{date: "2022-09-16T00:00:00.000Z"},
	{date: "2022-09-17T00:00:00.000Z"},
	{date: "2022-09-18T00:00:00.000Z"},
}

var toNajafDates = []DateModel{
	{date: "2022-09-11T00:00:00.000Z"},
	{date: "2022-09-12T00:00:00.000Z"},
	{date: "2022-09-13T00:00:00.000Z"},
	{date: "2022-09-14T00:00:00.000Z"},
	{date: "2022-09-15T00:00:00.000Z"},
	{date: "2022-09-16T00:00:00.000Z"},
	{date: "2022-09-17T00:00:00.000Z"},
	{date: "2022-09-18T00:00:00.000Z"},
}

type DateModel struct {
	date        string
	isAvailable bool
}

func Run() {
	defer func() {
		if r := recover(); r != nil {
			msg := "stacktrace from panic in meraj: \n" + string(debug.Stack())
			fmt.Println(msg)
			pkg.SendMessage(msg)
		}
	}()
	for {
		time.Sleep(10 * time.Second)
		runToTehran()
		runToNajaf()
	}
}

func runToTehran() {
	for i := 0; i < len(toTehranDates); i++ {
		time.Sleep(2 * time.Second)
		fmt.Printf("check meraj to tehran flight %s ...\n", toTehranDates[i].date)
		rsp, err := checkToTehranFlight(toTehranDates[i].date)
		if err != nil {
			println("error in ckeck meraj flight:", err.Error())
			pkg.SendMessage("error in ckeck meraj flight: " + err.Error())
			continue
		}
		if !rsp.Success {
			println("bad response in meraj: ", rsp.Success, rsp.Result)
			pkg.SendMessage("bad response error for meraj")
			continue
		}
		if checkAvailability(rsp) && !toTehranDates[i].isAvailable {
			toTehranDates[i].isAvailable = true
			pkg.SendToGroup(" پرواز معراج " + dateToDay[toTehranDates[i].date] + " به تهران " + " پیدا شد" + " https://meraj.aero ")
		}

		if !checkAvailability(rsp) && toTehranDates[i].isAvailable {
			toTehranDates[i].isAvailable = false
			pkg.SendToGroup("!پرواز معراج " + dateToDay[toTehranDates[i].date] + " به تهران " + " تموم شد")
		}
	}

}

func runToNajaf() {
	for i := 0; i < len(toNajafDates); i++ {
		time.Sleep(2 * time.Second)
		fmt.Printf("check meraj to najaf flight %s ...\n", toNajafDates[i].date)
		rsp, err := checkToNajafFlight(toNajafDates[i].date)
		if err != nil {
			println("error in ckeck meraj flight:", err.Error())
			pkg.SendMessage("error in ckeck meraj flight: " + err.Error())
			continue
		}
		if !rsp.Success {
			println("bad response in meraj: ", rsp.Success, rsp.Result)
			pkg.SendMessage("bad response error for meraj")
			continue
		}
		if checkAvailability(rsp) && !toNajafDates[i].isAvailable {
			toNajafDates[i].isAvailable = true
			pkg.SendToGroup(" پرواز معراج " + dateToDay[toNajafDates[i].date] + " به نجف " + " پیدا شد" + " https://meraj.aero ")
		}

		if !checkAvailability(rsp) && toNajafDates[i].isAvailable {
			toNajafDates[i].isAvailable = false
			pkg.SendToGroup("!پرواز معراج " + dateToDay[toNajafDates[i].date] + " به نجف " + " تموم شد")
		}
	}
}

func checkAvailability(rsp *Response) bool {
	if len(rsp.Result) == 0 {
		println("there is no any result in meraj response: ", rsp.Result)
		pkg.SendMessage("there is no any result in meraj response")
		return false
	}
	if len(rsp.Result[0].FlightGroups) > 0 {
		for _, flight := range rsp.Result[0].FlightGroups {
			if len(flight.Flights[0].Prices) > 0 && flight.Flights[0].Prices[0].Amount >= 30000000 && flight.Flights[0].Prices[0].Amount <= 50000000 {
				println("=========peyda shod==========")
				return true
			}
		}
	}
	return false
}

func checkToNajafFlight(date string) (*Response, error) {
	return checkFlight(getTehranToNajafInputJson(date))
}

func checkToTehranFlight(date string) (*Response, error) {
	return checkFlight(getNajafToTehranInputJson(date))
}

func checkFlight(input string) (*Response, error) {
	url := "https://flight-api-v1.meraj.aero/Flight/availability/GetFlightCityToCityAvailability"

	var jsonStr = []byte(input)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.Wrap(err, "error in creating request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Cookie", "_ga=GA1.2.191592258.1662472777; _session=aa6eeba8-4d3e-4872-9894-743a52b4b9ad")
	// req.Header.Set("Accept-Encoding", " gzip")
	req.Header.Set("Accept-Language", " en-US,en;q=0.9,fa;q=0.8")
	req.Header.Set("Connection", " keep-alive")
	req.Header.Set("Content-Length", " 464")
	req.Header.Set("Host", " flight-api-v1.meraj.aero")
	req.Header.Set("Origin", " https://meraj.aero")
	req.Header.Set("Referer", " https://meraj.aero/")
	req.Header.Set("Sec-Fetch-Dest", " empty")
	req.Header.Set("Sec-Fetch-Mode", " cors")
	req.Header.Set("Sec-Fetch-Site", " same-site")
	// req.Header.Set("User-Agent", " Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-mobile", "?0")

	client := &http.Client{}
	println("sending meraj request...")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error in sending request")
	}
	println("request meraj sent...")
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)

	// println("meraj body: ", buf.String())

	res := Response{}
	err = json.Unmarshal([]byte(buf.String()), &res)
	if err != nil {
		return nil, errors.Wrap(err, "error in unmarshal")
	}
	if !res.Success {
		println("meraj body: ", buf.String())
	}
	return &res, nil
}

func getTehranToNajafInputJson(date string) string {
	rsp := `
	{
		"availabilityRequest": {
		  "$id": 1,
		  "passengers": {
			"$id": 2,
			"$values": [
			  {
				"$id": 3,
				"type": 1,
				"quantity": 2,
				"optionalServices": {
				  "$id": 4,
				  "$values": []
				}
			  },
			  {
				"$id": 5,
				"type": 2,
				"quantity": 1,
				"optionalServices": {
				  "$id": 6,
				  "$values": []
				}
			  }
			]
		  },
		  "segments": {
			"$id": 7,
			"$values": [
			  {
				"$id": 8,
				"origin": "Tehran",
				"destination": "NJF",
				"departureDateTime": "$$date"
			  }
			]
		  },
		  "travelDetails": {
			"$id": 9,
			"cabinType": 1,
			"airTripType": 1,
			"stopQuantityType": 3,
			"pricingSourceType": 3
		  },
		  "availabilityType": 0
		}
	  }
	`

	return strings.ReplaceAll(rsp, "$$date", date)

}

func getNajafToTehranInputJson(date string) string {
	rsp := `
	{
		"availabilityRequest": {
		  "$id": 1,
		  "passengers": {
			"$id": 2,
			"$values": [
			  {
				"$id": 3,
				"type": 1,
				"quantity": 1,
				"optionalServices": {
				  "$id": 4,
				  "$values": []
				}
			  }
			]
		  },
		  "segments": {
			"$id": 5,
			"$values": [
			  {
				"$id": 6,
				"origin": "NJF",
				"destination": "Tehran",
				"departureDateTime": "$$date"
			  }
			]
		  },
		  "travelDetails": {
			"$id": 7,
			"cabinType": 1,
			"airTripType": 1,
			"stopQuantityType": 3,
			"pricingSourceType": 3
		  },
		  "availabilityType": 0
		}
	  }
	`

	return strings.ReplaceAll(rsp, "$$date", date)

}
