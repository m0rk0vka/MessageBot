package main


import (
	"net/http"
	"io/ioutil"
)


func main() {
	botToken := "6197887318:AAHzZzekPsqzemXfWZKrwt0a2KXZ9yd26ck"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
}


// requesting updates
func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "getUpdates")
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
}


// answering for updates
func respond() {

}