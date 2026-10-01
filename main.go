package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type StructThree struct {
	strStruct string
	intStruct int
	bolStruct bool
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		var answer string
		fmt.Println()
		fmt.Print("Function 1: Print every X line from a file \nFunction 2: Print the number of characters and bytes in a string\nFunction 3: Printing and modifying a struct \nFunction 4: Printing and modifying a struct list \nFunction 5: End the program \nPlease type the number of the function you wish to run: ")
		answer, err := inputReader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if err != nil {
			return
		}
		switch answer {
		case "1":
			{
				function1()
			}
		case "2":
			{
				function2()
			}
		case "3":
			{
				function3()
			}
		case "4":
			{
				function4()
			}
		case "5":
			{
				fmt.Println("Goodbye!")
			}
		default:
			{
				fmt.Println("Invalid option, please try again. Type the number of the function you wish to run: ")
			}
		}
		if answer == "5" {
			break
		}
		println(err)
	}
}

// File Reader from in class
func getFile() []string {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter filename: ")
	var txtFile string
	txtFile, err := inputReader.ReadString('\n')
	txtFile = strings.TrimSpace(txtFile)
	contents, err := os.ReadFile(txtFile)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	bigString := string(contents)
	allLines := strings.Split(bigString, "\n")
	return allLines
}

func function1() {
	inputReader := bufio.NewReader(os.Stdin)
	allLines := getFile()
	var setLineNum int
	fmt.Print("Enter the number for which every X lines to print: ")
	setLine, err := inputReader.ReadString('\n')
	setLine = strings.TrimSpace(setLine)
	setLineNum, err = strconv.Atoi(setLine)
	for i := setLineNum - 1; i < len(allLines); i += setLineNum {
		fmt.Println(allLines[i])
	}
	println(err)
}

func function2() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a String: ")
	textSet, err := inputReader.ReadString('\n')
	textSet = strings.TrimSpace(textSet)
	//Consulted the internet to discover len() gives number of bytes in a string and utf8.RuneCountInString() gives number of characters in a string
	//https://www.gofaq.org/en/how-to-get-the-length-of-a-string-in-go-bytes-vs-runes/
	textByteValue := strconv.Itoa(len(textSet))
	textCharValue := strconv.Itoa(utf8.RuneCountInString(textSet))
	fmt.Println("Number of characters in " + textSet + ": " + textCharValue)
	fmt.Println("Number of bytes in " + textSet + ": " + textByteValue)
	println(err)
}

func function3() {

	strSet := StructThree{
		strStruct: "string",
		intStruct: 1,
		bolStruct: false,
	}
	fmt.Print("Current Struct: ")
	fmt.Println(strSet)

	function3b(strSet)

	fmt.Print("Struct After Return: ")
	fmt.Println(strSet)
}

func function3b(strSet StructThree) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a String to modify Struct: ")
	structString, err := inputReader.ReadString('\n')
	structString = strings.TrimSpace(structString)
	strSet.strStruct = structString
	fmt.Print("New Struct: ")
	fmt.Println(strSet)
	println(err)
}

func function4() {

	strSet := []StructThree{
		{strStruct: "hello", intStruct: 1, bolStruct: false},
		{strStruct: "goodbye", intStruct: 2, bolStruct: true},
	}
	fmt.Print("Current Struct: ")
	fmt.Println(strSet)

	function4b(strSet)

	fmt.Print("Struct After Return: ")
	fmt.Println(strSet)
}

func function4b(strSet []StructThree) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a String to adjust Struct 2: ")
	structString, err := inputReader.ReadString('\n')
	structString = strings.TrimSpace(structString)
	strSet[1].strStruct = structString
	fmt.Print("New Struct: ")
	fmt.Println(strSet)
	println(err)
}
