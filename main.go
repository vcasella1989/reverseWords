package main

import (
	"strings"
	"fmt"
)


func main() {

	word := "Hello World"

	fmt.Println(word)

	//running reverse func and combining back into one string
	fmt.Println(strings.Join(reverseWords(word), " "))

	//Just here to keep the program from auto closing
	fmt.Println("enter any key to exit")
	fmt.Scanln()
    fmt.Println("closing")

	
}

func reverseWords(word string) []string {
	//Splitting Strings on empty space and storing in array
	words := strings.Split(word, " ")
	var wordsReversed []string

	//Looping over strings that we placed in array
	for _, x := range words{
		//Break string into Runes(chars). 
		individRunes := strings.Split(x, "")
		var reversedRunes []string

		for i := len(individRunes) - 1; i >= 0; i-- {
			//Add runes in reverse order to Array. 
			reversedRunes = append(reversedRunes, individRunes[i])
		}
		//combining runes back into strings and adding to an array
		wordsReversed = append(wordsReversed, strings.Join(reversedRunes, ""))
	}
	
	//retuning reversed words back to main
	return wordsReversed
}