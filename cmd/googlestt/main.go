package main

import (
	"os"
	"github.com/pavanchitta/go-scribe/src/googlestt"
)

func main() {
	audio_path := os.Args[1]
	googlestt.MakeLocalRequest(audio_path)
}
