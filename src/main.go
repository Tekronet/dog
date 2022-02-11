package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

//program information
const (
	PROGRAM_NAME        = "Dog"
	PROGRAM_VERSION     = "1.2"
	PROGRAM_DESCRIPTION = "Simple program for displaying text files in diffrent colors nad their info"
	PROGRAM_COPYRIGHT   = "Tekronet 2022 <tekronet.github.io>"
	PROGRAM_HELP_INFO   = "USAGE:\n    dog <filename/-v/-h> <flag (optional)>\n\n" +
		"FLAGS:\n    -v, --version - Print version information\n" +
		"    -h, --help - Print this help information\n" +
		"    -c, --color <COLOR> - Change output color\n" +
		"    -i, --info - Print file characters, words and lines rather than file content"
)

func main() {

	//checking if file is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file.",
			"For more help type 'dog --help'")
	} else {
		//version & help information
		if os.Args[1] == "--version" || os.Args[1] == "-v" {
			fmt.Println(PROGRAM_NAME, PROGRAM_VERSION)
			fmt.Println(PROGRAM_DESCRIPTION)
			fmt.Println(PROGRAM_COPYRIGHT)

		} else if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(PROGRAM_HELP_INFO)
			//reading file
		} else {
			fileName := os.Args[1]

			fileContent, err := ioutil.ReadFile(fileName)
			if err != nil {
				color.Set(color.FgRed)
				fmt.Println(fileName + ": file not found")
				color.Unset()
			}

			text := string(fileContent)

			//checking if color or info is provided
			if len(os.Args) > 2 {
				if os.Args[2] == "--color" || os.Args[2] == "-c" {
					switch os.Args[3] {
					case "black":
						color.Black(text)
					case "red":
						color.Red(text)
					case "green":
						color.Green(text)
					case "yellow":
						color.Yellow(text)
					case "blue":
						color.Blue(text)
					case "magenta":
						color.Magenta(text)
					case "cyan":
						color.Cyan(text)
					case "white":
						color.White(text)
					default:
						color.White(text)
					}
				} else if os.Args[2] == "--info" || os.Args[2] == "-i" {
					printInfo(fileName)
				} else {
					fmt.Println(text)
				}
			} else {
				fmt.Println(text)
			}
		}
	}
}

func printInfo(path string) {
	//word count
	fileHandle, err := os.Open(path)
	if err != nil {
		os.Exit(1)
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanWords)

	wordCount := 0

	for fileScanner.Scan() {
		wordCount++
	}

	//characters, letters, lines
	file, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}
	text := string(file)

	lineCount := (strings.Count(text, "\n") + 1)

	charCount := (utf8.RuneCountInString(text) - (lineCount - 1))
	letterCount := (charCount - strings.Count(text, " "))

	fmt.Println("Char count: " + strconv.Itoa(charCount))
	fmt.Println("Letter count: " + strconv.Itoa(letterCount))
	fmt.Println("Word count: " + strconv.Itoa(wordCount))
	fmt.Println("Line count: " + strconv.Itoa(lineCount))
}
