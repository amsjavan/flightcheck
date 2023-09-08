package pkg

import (
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {
	// SendToGroup("test")
	str := fmt.Sprintf("پرواز ایرتور %v در تاریخ %v به نجف پیدا شد %v",
		"12344",
		"12344",
		"https://iranairtour.ir")
	println(str)
	SendMessage(str)
	// SendMessage(fmt.Sprintf("%s", "error in ckeck iranair flight: error in sending request: Post \"https://ebooking.iranair.com/ibe/IR/availability/search?captchaInput=\": dial tcp 91.98.28.36:443: i/o timeout"))
	// SendMessage("salam")
}
