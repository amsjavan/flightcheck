package airtour

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
	"1402/06/04": "شنبه",
	"1402/06/05": "یکشنبه",
	"1402/06/06": "دوشنبه",
	"1402/06/08": "سه شنبه",
	"1402/06/09": "چهارشنبه",
	"1402/06/10": "پنج شنبه",
	"1402/06/11": "جمعه",
	"1402/06/12": "شنبه",
	"1402/06/13": "یکشنبه",
	"1402/06/14": "دوشنبه",
	"1402/06/15": "سه شنبه",
	"1402/06/16": "چهارشنبه",
	"1402/06/17": "پنج شنبه",
	"1402/06/18": "جمعه",
	"1402/06/19": "شنبه",
	"1402/06/20": "یکشنبه",
	"1402/06/21": "دوشنبه",
	"1402/06/22": "سه شنبه"}

// "1401/06/26": "تسسسسست"}

var dateModels = []DateModel{
	{date: "1402/06/04"},
	{date: "1402/06/05"},
	{date: "1402/06/06"},
	{date: "1402/06/07"},
	{date: "1402/06/08"},
	{date: "1402/06/09"},
	{date: "1402/06/10"},
	{date: "1402/06/11"},
	{date: "1402/06/12"},
	{date: "1402/06/13"},
	{date: "1402/06/14"},
	{date: "1402/06/15"},
	{date: "1402/06/16"},
	// {date: "1402/06/17"},
	// {date: "1402/06/18"},
	// {date: "1402/06/19"},
	// {date: "1402/06/20"},
	// {date: "1402/06/21"},
	// {date: "1402/06/22"},
}

// {date: "1401/06/26"}}

type DateModel struct {
	date        string
	isAvailable bool
}

func Run() {
	defer func() {
		if r := recover(); r != nil {
			msg := "stacktrace from panic in airtour: \n" + string(debug.Stack())
			fmt.Println(msg)
			pkg.SendMessage(msg)
		}
	}()
	for {
		time.Sleep(20 * time.Second)

		fmt.Printf("check airtour flight ...\n")
		rsp, err := checkFlight()
		if err != nil {
			pkg.SendMessage("error in ckeck airtour flight: " + err.Error())
			println("error in ckeck airtour flight:", err.Error())
			continue
		}
		if rsp.HasError || !rsp.Result.Successful {
			pkg.SendMessage("bad response in airtour")
			println("bad response in airtour: ", rsp)
			continue
		}

		for i := 0; i < len(dateModels); i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("check airtour date: %s \n", dateModels[i].date)
			if checkAvailability(rsp, dateModels[i].date) && !dateModels[i].isAvailable {
				dateModels[i].isAvailable = true
				pkg.SendToGroup(
					fmt.Sprintf("پرواز ایرتور %v در تاریخ %v به نجف پیدا شد %v",
						dateToDay[dateModels[i].date],
						dateModels[i].date,
						"https://iranairtour.ir"),
				)
			}

			if !checkAvailability(rsp, dateModels[i].date) && dateModels[i].isAvailable {
				dateModels[i].isAvailable = false
				pkg.SendToGroup(
					fmt.Sprintf("پرواز ایرتور %v در تاریخ %v به نجف تموم شد",
						dateToDay[dateModels[i].date],
						dateModels[i].date),
				)
			}
		}

	}
}

func checkAvailability(rsp *Response, date string) bool {
	for _, flight := range rsp.Result.FlightOffers {
		if flight.Date == date {
			return true
		}
	}
	return false
}

func checkFlight() (*Response, error) {
	url := "https://iranairtour.ir/Flight/SearchFlightCalendar2"

	var jsonStr = []byte(getInputJson())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.Wrap(err, "error in creating request")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	println("sending airtour request...")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error in sending airtour request")
	}
	println("request airtour sent...")
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	res := Response{}
	err = json.Unmarshal([]byte(buf.String()), &res)
	if err != nil {
		return nil, errors.Wrap(err, "error in airtour unmarshal")
	}
	return &res, nil
}

func getInputJson() string {
	rsp := `
	{"Origin":"IKA","Destination":"NJF","CalendarType":1,"HasReturn":0}
	`

	return rsp

}
