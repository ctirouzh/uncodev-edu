package helpers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func SendMessage(chat_id string,text string){

	token := "5xM"

	telegramUrl, err := url.Parse("https://api.telegram.org/bot" + token + "/sendMessage")
	if err!=nil {
		log.Println(err)
	}

	values := telegramUrl.Query()
	values.Add("chat_id",chat_id)
	values.Add("text",text)
	values.Add("parse_mode","html")
	telegramUrl.RawQuery = values.Encode()

	response , err := http.Get(telegramUrl.String())
	if err!=nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	///

}


