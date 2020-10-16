package main

import (
	"strings"
)

func main() {
	stringDecoder("WUBWEWUBAREWUBWUBTHEWUB BACKYARD WUBMYWUBFRIENDWUB")
	stringDecoder("AWUBBWUBC")
	stringDecoder("AWUBWUBWUBBWUBWUBWUBC")
	stringDecoder("WUBAWUBBWUBCWUB")
}

func stringDecoder(song string) string {
	// remove WUB
	song = strings.ReplaceAll(song, "WUB", " ")
	// adjust space
	song = strings.ReplaceAll(song, "   ", " ")
	song = strings.ReplaceAll(song, "  ", " ")
	song = strings.Trim(song, " ")
	
	return song
}


