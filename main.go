package main

import (
	ascii "ascii/func"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	args := os.Args
	counter := 0

	wordTofind, file, error := ascii.Validation(args)
	if error != nil {
		fmt.Println(error)
		return
	}
	//turn the file into an array
	fileArray := ascii.Convert(file)

	//convert the wordToFind into an array for each \n
	splitstring := strings.Split(wordTofind, "\\n")
	result := ""

	//count how many times does \n appear in wordTofind
	for i := 0; i < len(wordTofind); i++ {
		if wordTofind[i] == '\\' && len(wordTofind) > i+1 && wordTofind[i+1] == 'n' {
			counter++
		}
	}

	//check if wordTofind only has \n
	lenWord := len(wordTofind)
	if math.Round(float64(lenWord)/2) == float64(counter) {
		i := 0
		for i != counter {
			fmt.Print("\n")
			i++
		}
		return
	} else {
		// For loop to generate for each split
		for _, word := range splitstring {

			if word == "" {
				result += "\n"
			} else {

				asciiArt, err := ascii.GenerateAscii(word, fileArray)
				if err != nil {
					fmt.Println(err)
					return
				}
				result += asciiArt
			}

		}
	}

	//print the word as an ascii art into the terminal
	fmt.Print(result)
}
