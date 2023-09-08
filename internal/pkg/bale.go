package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendToGroup(input string) {
	fmt.Println(input)
	// send(873587972, input) // گروه خاندان
	send(1019326949, input) // گروه اطلاع رسانی
}

var dests = []int{2058851620}

func SendMessage(input string) {
	for _, peer := range dests {
		send(peer, input)
	}
}

func send(chatId int, input string) {
	input = strings.ReplaceAll(input, "\"", "")
	httpposturl := "https://tapi.bale.ai/1719568750:vZvuSc8TbD4mGrsVrbX9q2ucVfxLRIuOGJNRU6aS/sendMessage"
	fmt.Println("HTTP JSON POST URL:", httpposturl)
	var jsonData = []byte(fmt.Sprintf(`{"chat_id": "%d","text": "%s"}`, chatId, input))
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		println("error in sending bot message: ", error)
		return
	}
	if response.StatusCode != 200 {
		println("error in sending bot message. status: ", response.Status)
		return
	}
	defer response.Body.Close()

	fmt.Println("bot response Status:", response.Status)
	fmt.Println("bot response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("bot response Body:", string(body))
}
