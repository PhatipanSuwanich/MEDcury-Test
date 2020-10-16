package main

import (
	"testing"
)

func TestStringDecoder(t *testing.T) {

	if r := stringDecoder("WUBWEWUBAREWUBWUBTHEWUB BACKYARD WUBMYWUBFRIENDWUB"); r != "WE ARE THE BACKYARD MY FRIEND" {
		t.Error("WUBWEWUBAREWUBWUBTHEWUB BACKYARD WUBMYWUBFRIENDWUB should be 'WE ARE THE BACKYARD MY FRIEND' but have", r)
	}

	if r := stringDecoder("AWUBBWUBC"); r != "A B C" {
		t.Error("AWUBBWUBCk should be 'A B C' but have", r)
	}
	if r := stringDecoder("AWUBWUBWUBBWUBWUBWUBC"); r != "A B C" {
		t.Error("AWUBWUBWUBBWUBWUBWUBC should be 'A B C' but have", r)
	}
	if r := stringDecoder("WUBAWUBBWUBCWUB"); r != "A B C" {
		t.Error("WUBAWUBBWUBCWUB should be 'A B C' but have", r)
	}
}
