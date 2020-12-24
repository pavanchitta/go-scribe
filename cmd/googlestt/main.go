package main

import (
	"os"
	"github.com/pavanchitta/go-scribe/src/googlestt"
	"strings"
)

func main() {
	audio_path := os.Args[1]
	if strings.HasPrefix(audio_path, "gs://") {
		googlestt.MakeRemoteRequest(audio_path)
	} else {
		googlestt.MakeLocalRequest(audio_path)
	}
}
