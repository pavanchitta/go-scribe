package deepgram

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
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

func MakeConcurrentRequests(audio_filepath string, params map[string]string) {

	data, err := ioutil.ReadFile(audio_filepath)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Size of data byte array:", len(data))
	chunk_size := 1000000
	num_threads := 0
	if len(data) % chunk_size == 0 {
		num_threads = len(data) / chunk_size
	} else {
		num_threads = len(data) / chunk_size + 1
	}
	done := make(chan bool, num_threads)
	curr_arr := make([]byte, 0)
	for i, val := range data {
		if i % chunk_size == 0 {
			log.Println("Making go request on iter", i)
			go MakeRequest(curr_arr, params, done)
			curr_arr = curr_arr[:0]
		} else {
			curr_arr = append(curr_arr, val)
		}
	}

	if len(curr_arr) > 0 {
		go MakeRequest(curr_arr, params, done)
	}
	log.Println("Waiting for threads")
	for i := 1; i < num_threads; i++ {
		<-done
	}
}

func MakeRequest(data []byte, params map[string]string, done chan bool) string {
	log.Println("In make request")
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
		done <- true
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		done <- true
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		done <- true
		log.Fatalln(err)
	}
	log.Println("here")
	log.Println(string(body))
	done <- true
	return string(body)

}
	

func RequestDGLocal(audio_filepath string, params map[string]string) string {
	
	data, err := ioutil.ReadFile(audio_filepath)
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan bool, 1)
	res := MakeRequest(data, params, done)
	<-done
	return string(res)
}

