package main

import (
	"fmt"
	"strings"
)

func main() {
	var dr string = "/+1-541-754-3010 156 Alphand_St. <J Steeve>\n 133, Green, Rd. <E Kustur> NY 56423 ;+1-541-914-3010\n" + "1-541-984-3012<P Reed> /PO Box 530; Pollocksville, NC-28573\n :+1-321-512-2222 <Paul Dive> Sequoia Alley PQ-67209 \n" + "1-741-984-3090 <Peter Reedgrave>_Chicago\n :+1-921-333-2222 <Anna Stevens> Haramburu_Street AA-67209\n" + "1-111-544-8973 <Peter Pan> LA\n +1-921-512-2222 <Wilfrid Stevens> Wild Street AA-67209\n" + "<Peter Gone> LA ?+1-121-544-8974 \n <R Steell> Quora Street AB-47209 +1-481-512-2222\n" + "<Arthur Clarke> San Antonio $+1-121-504-8974 TT-45120\n <Ray Chandler> Teliman Pk. !+1-681-512-2222! AB 47209, \n" + "<Sophia Loren> +1-421-674-8974 Bern TP-46017 \n <Peter O'Brien> High Street +1-908-512-2222; CC-47209\n" + "<Anastasia> +48-421-674-8974 Via Quirinal Roma \n <P Salinger> Main Street, +1-098-512-2222, Denver\n" + "<C Powel>*+19-421-674-8974 Chateau des Fosses Strasbourg F-68000\n <Bernard Deltheil> +1-498-512-2222; Mount Av. Eldorado\n" + "+1-099-500-8000 <Peter Crush> Labrador Bd.\n +1-931-512-4855 <William Saurin> Bison Street CQ-23071\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n"
	phone(dr, "1-098-512-2222")
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

func phone(strng string, num string) string {
	var books []string = strings.Split(strng, "\n")
	var name, adress string
	var count int = 0
	for _, book := range books {
		if i := strings.Index(book, num); i > 0 {
			count++
			name = between(book, "<", ">")
			adress = findAddress(book, name, num)
		}
	}
	if count == 0 {
		return fmt.Sprintf("Error => Not found: %s", num)
	} else if count > 1 {
		return fmt.Sprintf("Error => Too many people: %s", num)
	} else {
		return fmt.Sprintf("Phone => %s, Name => %s, Address => %s", num, name, adress)
	}
}

func between(value string, first string, last string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, first)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, last)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(first)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func findAddress(book string, name string, num string) string {
	book = strings.ReplaceAll(book, "<"+name+">", "")
	book = strings.ReplaceAll(book, num, "")
	book = strings.ReplaceAll(book, "$+", "")
	book = strings.ReplaceAll(book, "/+", "")
	book = strings.ReplaceAll(book, ";+", "")
	book = strings.ReplaceAll(book, ":+", "")
	book = strings.ReplaceAll(book, "+", "")
	book = strings.ReplaceAll(book, ";", "")
	book = strings.ReplaceAll(book, "_", " ")

	// adjust space
	book = strings.ReplaceAll(book, "  ", " ")
	return strings.Trim(book, " ")
}
