package tcharter

import (
	"crypto/tls"
	"flightcheck/internal/pkg"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var staurdayTicket bool
var sundayTicket bool

func Run() {
	url := "https://www.tcharter.ir/tickets/dates/IKA-NJF-airplane"

	now := time.Now()
	for {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		resp, err := http.Get(url)
		if err != nil {
			println("error in tcharter : ", err.Error())
			continue
		}
		defer resp.Body.Close()

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		// Find the review items
		doc.Find(".daterow").Each(func(i int, s *goquery.Selection) {
			day := s.Find("div span").Eq(0).Text()
			price := s.Find("div span").Eq(1).Text()
			msg := check(day, price)
			if msg != "" {
				pkg.SendToGroup(msg)
			}
		})
		println("checking tcharter...")
		time.Sleep(10 * time.Second)

		if time.Since(now) > 15*time.Minute {
			now = time.Now()
			pkg.SendMessage("من زنده ام")
		}
	}
}

const (
	FindSaturday  = "بلیط تی جارتر برای شنبه پیدا شد"
	FindSunday    = "بلیط تی جارتر برای یکشنبه پیدا شد"
	CloseSaturday = "بلیط تی جارتر شنبه تموم شد"
	CloseSunday   = "بلیط تی جارتر یکشنبه تموم شد"
	CloseText     = "کلیک کنید"
	Friday        = "جمعه18 شهریور"
	Saturday      = "شنبه19 شهریور"
	//Saturday = "جمعه18 شهریور"
	Sunday = "یکشنبه20 شهریور"
)

func check(day, price string) string {
	if strings.Contains(day, Saturday) {
		if price != CloseText && !staurdayTicket {
			println(price)
			staurdayTicket = true
			return FindSaturday + " " + price
		}
		if price == CloseText && staurdayTicket {
			staurdayTicket = false
			return CloseSaturday
		}
	}

	if strings.Contains(day, Sunday) {
		if price != CloseText && !sundayTicket {
			println(price)
			// newPrice := strings.ReplaceAll(price, "تومان", "")
			// newPrice = strings.ReplaceAll(newPrice, ",", "")

			// priceVal, err := strconv.Atoi(newPrice)
			// if err != nil {
			// 	println("error in converting to int")
			// 	println(price)
			// 	println(newPrice)
			// }
			// if priceVal > 11427000 {

			// }

			sundayTicket = true
			return FindSunday + " " + price
		}
		if price == CloseText && sundayTicket {
			sundayTicket = false
			return CloseSunday
		}
	}
	return ""
}
