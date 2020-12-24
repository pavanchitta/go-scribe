package googlestt

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"regexp"
	"log"
)

func SearchTranscript(filepath string, keywords []string) map[string][]string {

	// Construct the map that stores the timestamps for each keyword
	keymap := make(map[string][]string)
	for _, key := range keywords {
		keymap[strings.ToLower(strings.TrimSpace(key))] = make([]string, 0) 
	}
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("Coulldn't read file: ", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)	
	// parse the file one line at a time, populating the map if we 
	// find a hit
	for scanner.Scan() {

		line := scanner.Text()
		if !strings.HasPrefix(line, "Word") {
			continue
		}
		r := regexp.MustCompile(`Word: "(\S+)" \(startTime=([0-9]*[.])?[0-9]+, endTime=([0-9]*[.])?[0-9]+\)`)
		rs := r.FindStringSubmatch(line)
		fmt.Println("Match:", rs)
		if len(rs) < 3 {
			log.Println("Couldn't find complete match: ", rs)
			continue
		}
		word := strings.ToLower(rs[1])
		ts := rs[2]
		// TODO: Need to account for punctuation that might be present in words"
		// TODO: More generally, need a way to do fuzzy searching
		if val, ok := keymap[word]; ok {
			keymap[word]= append(val, ts)
		}
	}
	return keymap
}
