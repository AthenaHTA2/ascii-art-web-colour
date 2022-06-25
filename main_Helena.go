package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cut(s, sep string) (before, after string, found bool) { // Position of letters to be coloured is given after the "_" character, e.g.: red_3
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

func Find(slice []int, sliceItem int) bool {
	for item := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

func PrintArt(n string, y map[int][]string) {
	// prints horizontally
	colorChoice := os.Args[2][8:]
	var l []int
	mRGB := make(map[string]string) // Table of colours linked to terminal input
	mRGB["red"] = "\033[38;2;" + "255;" + "0;" + "0" + "m"
	mRGB["orange"] = "\033[38;2;" + "255;" + "165;" + "0" + "m"
	mRGB["green"] = "\033[38;2;" + "11;" + "232;" + "11" + "m"
	mRGB["yellow"] = "\033[38;2;" + "242;" + "244;" + "78" + "m"
	mRGB["blue"] = "\033[38;2;" + "0;" + "0;" + "255" + "m"
	mRGB["purple"] = "\033[38;2;" + "177;" + "30;" + "228" + "m"
	mRGB["cyan"] = "\033[38;2;" + "31;" + "239;" + "239" + "m"
	mRGB["white"] = "\033[38;2;" + "255;" + "255;" + "255" + "m"
	mRGB["gray"] = "\033[38;2;" + "199;" + "187;" + "187" + "m"
	mRGB["31"] = "\033[38;2;" + "255;" + "0;" + "0" + "m"     // red
	mRGB["39"] = "\033[38;2;" + "255;" + "165;" + "0" + "m"   // orange
	mRGB["32"] = "\033[38;2;" + "255;" + "165;" + "0" + "m"   // green
	mRGB["33"] = "\033[38;2;" + "242;" + "244;" + "78" + "m"  // yellow
	mRGB["34"] = "\033[38;2;" + "0;" + "0;" + "255" + "m"     // blue
	mRGB["35"] = "\033[38;2;" + "177;" + "30;" + "228" + "m"  // purple
	mRGB["36"] = "\033[38;2;" + "31;" + "239;" + "239" + "m"  // cyan
	mRGB["37"] = "\033[38;2;" + "255;" + "255;" + "255" + "m" // white
	mRGB["38"] = "\033[38;2;" + "199;" + "187;" + "187" + "m" // gray

	before, after, found := Cut(colorChoice, "_") // check if a command to colour single letters was given
	color := before
	chooseLetters := after

	if !found { // if the suffix '_#' is not present, colour the entire word
		for j := 0; j < 8; j++ {
			for _, letter := range n { // for each letter in string n
				output := fmt.Sprintf("%s%s", mRGB[color], y[int(letter)][j])
				fmt.Print(output)
			}
			fmt.Println()
		}
	} else {
		numLetters := len(chooseLetters)
		switch numLetters {
		case 1: // colour only one letter in the word
			g, _ := strconv.Atoi(string(chooseLetters[0]))
			if g < 0 || g > len(n) {
				fmt.Println("Index can not be negative or longer than string length.")
				fmt.Println()
				os.Exit(0)
			} else {
				if g == len(n) {
					g = g - 1
				}
				l = append(l, g) // appending the letter in int format
			}

		case 2: // colour two letters in the word
			if chooseLetters[1] == ':' { // e.g. --color=red_3: will colour letters from position 4 to the end, as index for the first letter is 0.
				g, _ := strconv.Atoi(string(chooseLetters[0]))
				if g < 0 || g > len(n) {
					fmt.Println("Index can not be negative or longer than string length.")
					fmt.Println()
					os.Exit(0)
				} else {
					if g == len(n) {
						g = g - 1
					}
					for i := g; i < len(n); i++ {
						l = append(l, i)
					}
				}
				break
			}
			if chooseLetters[0] == ':' { // e.g. --color=red_:3 will colour letters from position 1 to 4.
				g, _ := strconv.Atoi(string(chooseLetters[1]))
				if g < 0 || g > len(n) {
					fmt.Println("Index can not be negative or longer than string length.")
					fmt.Println()
					os.Exit(0)
				} else {
					if g == len(n) {
						g = g - 1
					}
					for i := 0; i <= g; i++ {
						l = append(l, i)
					}
				}
				break
			}
			g, _ := strconv.Atoi(string(chooseLetters[0]))
			k, _ := strconv.Atoi(string(chooseLetters[1]))
			if g < 0 || k < 0 || g > len(n) || k > len(n) {
				fmt.Println("Index can not be negative or longer than string length.")
				fmt.Println()
				os.Exit(0)
			} else {
				if g == len(n) {
					g = len(n) - 1
				}
				if k == len(n) {
					k = len(n) - 1
				}
				if g > k {
					temp := g
					g = k
					k = temp
				}
				l = append(l, g)
				l = append(l, k)
			}

		case 3: // colour a range of contiguous letters
			if chooseLetters[1] == ':' { // e.g. --color=red_1:3 will colour letters in positions 2, 3, and 4, as index for the first letter is 0.
				g, _ := strconv.Atoi(string(chooseLetters[0]))
				k, _ := strconv.Atoi(string(chooseLetters[2]))
				if g > k { // if first index is greater than the second index swap them over
					temp := g
					g = k
					k = temp
				}
				if g < 0 || k < 0 || g > len(n) || k > len(n) {
					fmt.Println("Index can not be negative or longer than string length.")
					fmt.Println()
					os.Exit(0)
				} else {
					if g == len(n) {
						g = len(n) - 1
					}
					if k == len(n) {
						k = len(n) - 1
					}
					for i := g; i <= k; i++ {
						l = append(l, i)
					}
				}
			} else { // e.g. --color=red_145 will colour the second, fifth and sixth letters, as the first letter has index 0.
				for i := 0; i < numLetters; i++ {
					g, _ := strconv.Atoi(string(chooseLetters[i]))
					if g < 0 || g > len(n) {
						fmt.Println("Index can not be negative or longer than string length.")
						fmt.Println()
						os.Exit(0)
					} else {
						l = append(l, g)
					}
				}
			}
		default:
			msg := fmt.Sprintf("%s%s", "\033[38;2;"+"255;"+"165;"+"0"+"m", "Position of letters to be coloured not provided.") // if there is nothing after the underscore exit program
			fmt.Println(msg)
			os.Exit(0)
		}

		for j := 0; j < 8; j++ { // for each of 8 lines
			index := 0
			for c, letter := range n { // for each letter in the string
				if l[index] == c { //if letter position equals the position given in terminal command
					output := fmt.Sprintf("%s%s", mRGB[color], y[int(letter)][j]) //colour that letter
					fmt.Print(output)

					if index < len(l)-1 {
						index++ //move to the next position given in terminal command
					} else {
						index = 0
					}
				} else { //letter positions that are not equal to positions in terminal command must be coloured white
					output := fmt.Sprintf("%s%s", mRGB["white"], y[int(letter)][j])
					fmt.Print(output)
				}
			} //--> loops letters in word
			fmt.Println()

		} // --> loops eight lines in banner
	}
}

func main() {
	// open the text file
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [OPTION]\n\nEX: go run . something --color=<color>")
		os.Exit(0)
	} else {

		if os.Args[2][:8] != "--color=" {
			fmt.Println("Usage: go run . [STRING] [OPTION]\n\nEX: go run . something --color=<color>")
			os.Exit(0)
		}

		file, err := os.Open("standard.txt")
		if err != nil {
			fmt.Println("Please double check the txt file name")
		}

		defer file.Close()
		// read the file
		Scanner := bufio.NewScanner(file)

		// identify the letters with ascii code
		var lines []string
		for Scanner.Scan() {
			lines = append(lines, Scanner.Text())
		}
		asciiChrs := make(map[int][]string) // map keys are ascii symbols' decimal numbers, and the map values are corresponding images' lines.
		dec := 31

		for _, line := range lines {
			if line == "" {
				dec++
			} else {
				asciiChrs[dec] = append(asciiChrs[dec], line)
			}
		}

		args := os.Args[1]
		for i := 0; i < len(args); i++ {
			if args[i] == 92 && args[i+1] == 110 { // checking if a "\n" was typed in the terminal
				PrintArt(string(args[:i]), asciiChrs)   // if yes, print the string characters up to "\n"
				PrintArt(string(args[i+2:]), asciiChrs) // then skip i and i+1 that correspond to "\n", and print the rest ona new line.
			}
		}
		if !strings.Contains(args, "\\n") {
			PrintArt(args, asciiChrs)

			//	fmt.Println()
			//	fmt.Println(asciiChrs[65]) //prints the lines corresponding to 'A' = 65
			//	fmt.Println()
		}
	}
}
