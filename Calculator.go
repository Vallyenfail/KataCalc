package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input an expression \nType exit to terminate the program: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		if cmd == "exit" {
			os.Exit(1)
		}
		CalcRom(cmd)
	}
}

func CalcArabic(msg string) {
	words := strings.Fields(msg)
	result := 0
	var numbers [3]int
	for idx, word := range words {
		if idx == 1 {
			numbers[idx] = 0 // Don't know hot to use steps here in iteration
			continue
		}
		numbers[idx], _ = strconv.Atoi(word)
	}
	switch { // make negative results!
	case words[1] == "+":
		result = numbers[0] + numbers[2]
	case words[1] == "-":
		result = numbers[0] - numbers[2]
	case words[1] == "/":
		result = numbers[0] / numbers[2] // check the division!!!
	case words[1] == "*":
		result = numbers[0] * numbers[2]
	default:
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	fmt.Println(result)
}

func CalcRom(msg string) {
	RomToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	ArabToRom := map[int]string{
		0:  "0",
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VII",
		9:  "IX",
		10: "X",
	}
	words := strings.Fields(msg)
	if len(words) != 3 {
		fmt.Println("Check your input, must be a mistake")
		os.Exit(1)
	}
	var numbers [3]int
	result := 0
	for idx, word := range words {
		numbers[idx] = RomToArab[word]
	}
	counter := 0
	for _, zero := range numbers {
		if zero == 0 { // if there are two zeroes then it's an error if there are three of them then proceed to arabic calc
			counter += 1
		}
	}
	if counter == 2 {
		fmt.Println("Check your input, there's an error")
		os.Exit(1)
	} else if counter == 3 { // If there is something wrong with the input and the counter = 3, the CalcArabic will fix it
		CalcArabic(msg)

	} else { // If counter is 1 it continues to calculate the Romanian
		switch {
		case words[1] == "+":
			result = numbers[0] + numbers[2]
		case words[1] == "-":
			result = numbers[0] - numbers[2]
		case words[1] == "/":
			result = numbers[0] / numbers[2]
		case words[1] == "*":
			result = numbers[0] * numbers[2]
		default:
			fmt.Println("Check the input, something is wrong")
			os.Exit(1)
		}
		switch {
		case result > 10:
			var result1 = "X" + ArabToRom[result-10]
			fmt.Println(result1)
		case result <= 10 && result >= 0:
			var result2 = ArabToRom[result]
			fmt.Println(result2)
		case result < 0:
			fmt.Println("The result is below zero, I can't calculate it")
			os.Exit(1)
		default:
			fmt.Println("XX")
		}
	}
}

// check the number of numbers (like len of the text) if len of the text is good check ten then rome or arabic return error if nothing's proven
// first check whether two numbers are arabic or not (len of the text again unless it's ten)
// then if yes check whether there is a sign or not
// check the result if it's romanian whether it's below zero
