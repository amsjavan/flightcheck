package atta

import (
	"crypto/tls"
	"encoding/json"
	"flightcheck/internal/pkg"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var dateToDay = map[int32]string{
	4:  "شنبه",
	5:  "یکشنبه",
	6:  "دوشنبه",
	7:  "سه شنبه",
	8:  "چهارشنبه",
	9:  "پنج‌شنبه",
	10: "جمعه",
	11: "شنبه",
	12: "یکشنبه",
	13: "یکشنبه",
	14: "یکشنبه",
	15: "یکشنبه",
}

var toTehranDates = []DateModel{
	{day: 7},
	{day: 8},
	{day: 9},
	{day: 10},
	{day: 11},
	{day: 12},
	{day: 13},
	{day: 14},
	{day: 15},
}

var toNajafDates = []DateModel{
	{day: 4},
	{day: 5},
	{day: 6},
	{day: 7},
	{day: 8},
	{day: 9},
	{day: 10},
	{day: 11},
	{day: 12},
	{day: 13},
	{day: 14},
	{day: 15},
}

type DateModel struct {
	day         int32
	isAvailable bool
}

func Run() {
	defer func() {
		if r := recover(); r != nil {
			msg := "stacktrace from panic in atta: \n" + string(debug.Stack())
			fmt.Println(msg)
			pkg.SendMessage(msg)
		}
	}()
	for {
		time.Sleep(4 * time.Minute)
		runToNajaf()
		runToTehran()
	}
}

func runToTehran() {
	for i := 0; i < len(toTehranDates); i++ {
		time.Sleep(10 * time.Second)
		fmt.Printf("check atta to tehran flight %d ...\n", toTehranDates[i].day)
		rsp, err := checkToTehranFlight(toTehranDates[i].day)
		if err != nil {
			pkg.SendMessage("error in ckeck atta flight: " + err.Error())
			println("error in ckeck atta flight:", err.Error())
			continue
		}
		if !rsp.Success {
			pkg.SendMessage("bad response in atta")
			println("bad response in atta: ", rsp)
			continue
		}
		if checkAvailability(rsp) && !toTehranDates[i].isAvailable {
			toTehranDates[i].isAvailable = true
			pkg.SendToGroup(
				fmt.Sprintf("پرواز آتا %v در تاریخ %v شهریور به تهران پیدا شد %v",
					dateToDay[toTehranDates[i].day],
					toTehranDates[i].day,
					"https://app.ataair.ir/flights"),
			)
		}

		if !checkAvailability(rsp) && toTehranDates[i].isAvailable {
			toTehranDates[i].isAvailable = false
			pkg.SendToGroup(
				fmt.Sprintf("پرواز آتا %v در تاریخ %v شهریور به تهران تموم شد",
					dateToDay[toTehranDates[i].day],
					toTehranDates[i].day),
			)
		}
	}
}

func runToNajaf() {
	for i := 0; i < len(toNajafDates); i++ {
		time.Sleep(10 * time.Second)
		fmt.Printf("check atta to najaf flight %d ...\n", toNajafDates[i].day)
		rsp, err := checkToNajafFlight(toNajafDates[i].day)
		if err != nil {
			pkg.SendMessage("error in ckeck atta flight: " + err.Error())
			println("error in ckeck atta flight:", err.Error())
			continue
		}
		if !rsp.Success {
			pkg.SendMessage("bad response in atta")
			println("bad response in atta: ", rsp)
			continue
		}
		if checkAvailability(rsp) && !toNajafDates[i].isAvailable {
			toNajafDates[i].isAvailable = true
			pkg.SendToGroup(
				fmt.Sprintf("پرواز آتا %v در تاریخ %v شهریور به نجف پیدا شد \n%v",
					dateToDay[toNajafDates[i].day],
					toNajafDates[i].day,
					"https://app.ataair.ir/flights"),
			)
		}

		if !checkAvailability(rsp) && toNajafDates[i].isAvailable {
			toNajafDates[i].isAvailable = false
			pkg.SendToGroup(
				fmt.Sprintf("پرواز آتا %v در تاریخ %v شهریور به نجف تموم شد",
					dateToDay[toNajafDates[i].day],
					toNajafDates[i].day),
			)
		}
	}
}

func checkAvailability(rsp *Response) bool {
	return len(rsp.Result) > 0 && !rsp.Result[0].IsCapacityFull
}

func checkToTehranFlight(day int32) (*Response, error) {
	return checkFlight(getToTehranUrl(day))
}

func checkToNajafFlight(day int32) (*Response, error) {
	return checkFlight(getToNajafUrl(day))
}

func checkFlight(url string) (*Response, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error in sending atta request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	res := Response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Wrap(err, "error in unmarshal")
	}
	if !res.Success {
		println("atta body ====", string(body))
	}
	return &res, nil
}

func getToTehranUrl(day int32) string {
	url := "https://api.app.ataair.ir/v1/booking-api/AvailableFlights?source=NJF&target=IKA&day=$$day&month=6&adultQty=1&childQty=0&infantQty=0"

	return strings.ReplaceAll(url, "$$day", fmt.Sprintf("%d", day))
}

func getToNajafUrl(day int32) string {
	url := "https://api.app.ataair.ir/v1/booking-api/AvailableFlights?source=IKA&target=NJF&day=$$day&month=6&adultQty=1&childQty=0&infantQty=0"

	return strings.ReplaceAll(url, "$$day", fmt.Sprintf("%d", day))
}
