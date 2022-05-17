package main

import (
	"fmt"
	"os"
	"strings"
)


func checkSentenceStart(sentence string) bool {
	var start = sentence[0:1]
	return strings.ToUpper(start) == start
}

func checkSentenceEnd(sentence string) bool {
	if strings.HasSuffix(sentence, ".") {
		return true
	} else if strings.HasSuffix(sentence, "!") {
		return true
	} else if strings.HasSuffix(sentence, "?") {
		return true
	}
	return false
}


func main() {
	var sb strings.Builder

	for index, element := range os.Args {
		if index > 0 {
			sb.WriteString(element)
			sb.WriteString(" ")
		}
	}
	var sentence = sb.String()
	sentence = strings.Trim(sentence, " ")

	fmt.Println(sentence)
	if checkSentenceStart(sentence) {
		fmt.Println("Start OK")
	} else {
		fmt.Println("Start Invalid")
	}

	if checkSentenceEnd(sentence) {
		fmt.Println("End OK")
	} else {
		fmt.Println("End Invalid")
	}
}
