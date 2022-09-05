package helpers


import (
	"fmt"
	"log"
    "net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"
)


type KavenegarResponseMessage struct {
	Return struct {
		Status  int `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries       []struct {
		Messageid        int64  `json:"messageid"`
		Message          int    `json:"message"`
		Status           int    `json:"status"`
		Statustext       int    `json:"statustext"`
		Sender           int    `json:"sender"`
		Receptor         int    `json:"receptor"`
		Date             int16  `json:"date"`
		Cost             int    `json:"cost"`
	} `json:"entries"`
}


func KavenagarSendMessage(mobile string,code string) bool{

	token := "53792F626"

	urlA, err := url.Parse("https://api.kavenegar.com/v1/" + token + "/verify/lookup.json")
	if err != nil {
		log.Fatal(err)
	}
	values := urlA.Query()
	values.Add("receptor", mobile)
    values.Add("token", code)
    values.Add("template", "newuncodev")
    values.Add("type", "sms")
	urlA.RawQuery = values.Encode()
	resp, err := http.Get(urlA.String())
    if err != nil {
        fmt.Println("No response from request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body) // response body is []byte
    var result KavenegarResponseMessage
    if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
        //fmt.Println("Can not unmarshal JSON")
    }
    fmt.Println(result.Entries[0].Messageid)
	return true
}
