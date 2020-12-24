package main


import (
	"os"
	"github.com/pavanchitta/go-scribe/src/googlestt"
	"fmt"
	"strings"
)


func main() {
	filepath := os.Args[1]
	fmt.Println("Enter comma delimited list of keywords: ")
	var keywords_str string
	fmt.Scanln(&keywords_str)
	keywords := strings.Split(keywords_str, ",")
	res := googlestt.SearchTranscript(filepath, keywords)
	for word, times := range res {
		fmt.Println("Keyword: ", word, "    Times (s): ", times)
	}
}
