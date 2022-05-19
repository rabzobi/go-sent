package main

import (
	"fmt"
	"os"
	"strings"
)


func checkPalindrome(sentence string) bool {
	var length = len(sentence)
	for pos, char := range sentence {
		var char1 = byte(char)
		var char2 = []byte(sentence[length-pos-1:length-pos])[0]
		//fmt.Println("char1=",char1," char2=",char2)
		if (char1 != char2) {
			return false
		}
	}	
	return true
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
	if checkPalindrome(sentence) {
		fmt.Println("Valid palindrome")
	} else {
		fmt.Println("Not palindrome")
	}

}
