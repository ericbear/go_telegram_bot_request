package message

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	uri "net/url"
	"fmt"
)

func GetUpdates(url string, lastUpdateId int) Response {
	u := fmt.Sprintf("%s%s", url, "getUpdates")
	
	if lastUpdateId != 0 {
		fmt.Sprintf("%s%s?offset=%d", url, "getUpdates", lastUpdateId)
	}
	
	resp, _ := http.Get(u)
	defer resp.Body.Close()
	
	body, _ := ioutil.ReadAll(resp.Body)
	
	var m Response
	json.Unmarshal(body, &m)
	
	return m
}

func SendMessage(url string, chat_id int, msg string, replyId int) Response {
	resp, _ := http.Get(fmt.Sprintf("%s%s?chat_id=%d&text=%s&reply_to_message_id=%d", url, "sendMessage", chat_id, uri.QueryEscape(msg), replyId))
	defer resp.Body.Close()
	
	body, _ := ioutil.ReadAll(resp.Body)
	
	var m Response
	json.Unmarshal(body, &m)
	
	return m
}