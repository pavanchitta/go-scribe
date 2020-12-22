package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"os"
)
func RequestDG(audio_url string) string {
	/* :param audio_url : url to remote audio file hosted on the web */
	requestBody, err := json.Marshal(map[string]string{
		"url": audio_url,
	})

	if err != nil {
		log.Fatalln(err)
	}

	client := http.Client{}
	main_url := "https://brain.deepgram.com/v2/listen"
	request, err := http.NewRequest("POST", main_url, bytes.NewBuffer(requestBody))
	request.SetBasicAuth("pchitta@caltech.edu", "tenjId-jeswe9-xocfav")
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("here")
	log.Println(string(body))
	return string(body)
}

func main() {
	audio_url := os.Args[1]
	resp := RequestDG(audio_url)
	log.Println(resp)
}
