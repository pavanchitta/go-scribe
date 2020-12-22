package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"os"
	"bufio"
	"fmt"
	"strings"
)
func RequestDG(audio_url string, params map[string]string) string {
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
	request.SetBasicAuth(params["username"], params["password"])
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


func RequestDGLocal(audio_filepath string, params map[string]string) string {
	
	data, err := ioutil.ReadFile(audio_filepath)
	if err != nil {
		log.Fatalln(err)
	}
	requestBody, err := json.Marshal(map[string]string{
		"data":string(data),
	})

	if err != nil {
		log.Fatalln(err)
	}

	client := http.Client{}
	main_url := "https://brain.deepgram.com/v2/listen"
	request, err := http.NewRequest("POST", main_url, bytes.NewBuffer(requestBody))
	request.SetBasicAuth(params["username"], params["password"])
	request.Header.Set("Content-Type", "audio/m4a")
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
	audio_path := os.Args[1]
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	user, _ := reader.ReadString('\n')
	user = strings.Replace(user, "\n", "", -1)
	fmt.Print("Enter password: ")
	pw, _ := reader.ReadString('\n')
	pw = strings.Replace(pw, "\n", "", -1)
	fmt.Println(user, pw)
	params := map[string]string{"username":user, "password":pw}
	var resp string
	if strings.HasPrefix(audio_path, "http") {
		resp = RequestDG(audio_path, params)
	} else {
		resp = RequestDGLocal(audio_path, params)
	}
	log.Println(resp)
}
