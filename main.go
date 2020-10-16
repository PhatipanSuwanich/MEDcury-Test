package main

import (
	"fmt"
	"strings"
)

func main() {
	stringDecoder("WUBWEWUBAREWUBWUBTHEWUB BACKYARD WUBMYWUBFRIENDWUB")
	stringDecoder("AWUBBWUBC")
	stringDecoder("AWUBWUBWUBBWUBWUBWUBC")
	stringDecoder("WUBAWUBBWUBCWUB")
}

//use package strings
func stringDecoder(song string) {
	// remove WUB
	song = strings.ReplaceAll(song, "WUB", " ")
	// adjust spac
	song = strings.ReplaceAll(song, "   ", " ")
	song = strings.ReplaceAll(song, "  ", " ")
	song = strings.Trim(song, " ")
	// print
	fmt.Println(song)
}
