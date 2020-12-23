package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"log"
	"github.com/pavanchitta/go-scribe/src/deepgram"
)
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
		resp = deepgram.RequestDG(audio_path, params)
	} else {
		resp = deepgram.RequestDGLocal(audio_path, params)
	}
	log.Println(resp)
}
