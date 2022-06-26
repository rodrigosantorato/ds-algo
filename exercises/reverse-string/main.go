package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "Hello, my name is Kai"

	reversedString := reverseString(myString)

	fmt.Println(reversedString)
}

func reverseString(myString string) string {
	var reversedCharactersArray []string

	for i := len(myString) - 1; i >= 0; i-- {
		reversedCharactersArray = append(reversedCharactersArray, string(myString[i]))
	}
	return strings.Join(reversedCharactersArray, "")
}

// My first solution
//
// func reverseString(myString string) string {
// 	arrayOfCharacters := strings.Split(myString, "")
// 	arrayLength := len(arrayOfCharacters)
// 	var reversedCharactersArray []string
//
// 	for i := 0; i < arrayLength; i++ {
// 		reversedCharactersArray = append(reversedCharactersArray, arrayOfCharacters[arrayLength-1-i])
// 	}
//
// 	return strings.Join(reversedCharactersArray, "")
// }
