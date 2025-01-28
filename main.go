package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
	"strconv"
	"test_7/logic"
)

var scanner *bufio.Scanner

func main() {

	if err := keyboard.Open(); err != nil {
		fmt.Println("Error opening keyboard:", err)
		return
	}
	defer keyboard.Close()

	scanner = bufio.NewScanner(os.Stdin)
	selected := 0
	options := []string{
		"[1] test_1: Find the most valuable path",
		"[2] test_2: Catch me left-right-equal",
		"[3] test_3: Pie Fire Dire",
		"[Q] Exit: 🚫",
	}

	for {
		printMenu(options, selected)

		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			break
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selected > 0 {
				selected--
			}
		case keyboard.KeyArrowDown:
			if selected < len(options)-1 {
				selected++
			}
		case keyboard.KeyEnter:
			fmt.Print("\033[H\033[2J")
			fmt.Printf("You selected: [ %s ]\n", options[selected])
			switch selected {
			// test_1: Find the most valuable path
			case 0:
				// Receive user input
				filePath := "input/InputFindTheMostValuablePath.txt"
				fmt.Printf("\nPaste your input data into the file. [test_7/%s]\n", filePath)
				fmt.Printf("\nPlease enter word.\n")
				fmt.Printf("\"default\" - input with your long example message.\n")
				fmt.Printf("\"ok\" - if you successfully placed the input.\n")
				fmt.Printf("\"q\" - if you want to quit.\n")
				fmt.Print("Enter your word: ")
				for {
					if scanner.Scan() {
						acceptInput := scanner.Text()

						if acceptInput == "ok" || acceptInput == "default" {
							if acceptInput == "default" {
								filePath = "input/DefaultInputFindTheMostValuablePath.txt"
							}
							data, err := ioutil.ReadFile(filePath)
							if err != nil {
								data, err = ioutil.ReadFile("../" + filePath)
								if err != nil {
									fmt.Printf("\ncannot reading file input,%s\n", err.Error())
								}
							}
							// Call logic
							result, err := logic.FindTheMostValuablePath(string(data))
							if err != nil {
								fmt.Printf("\nError: %s \n\n", err)
								selectMenu()
							} else {
								fmt.Printf("Input:\n%s \n\n", string(data))
								// Print result
								fmt.Printf("\nResult: %s \n\n", strconv.FormatFloat(result, 'f', -1, 64))
								selectMenu()
							}
							break

						} else if acceptInput == "q" {
							break

						} else {
							fmt.Print("Please enter \"ok\" or \"q\": ")
						}
					}
				}

			// test_2: Catch me left-right-equal
			case 1:
				// Receive user input
				var input string
				var result string
				fmt.Println("Input your data [string of (L, R, =)]:")
				if scanner.Scan() {
					input = scanner.Text()
				}

				// Call logic
				result, err = logic.CatchMeLeftRightEqual(input)
				if err != nil {
					fmt.Printf("\nError: %s \n\n", err)
					selectMenu()
					break
				}
				fmt.Printf("\nResult: %s \n\n", result)
				selectMenu()

			case 2:
				app := fiber.New()
				app.Post("/beef/summary", logic.PieFireDire)
				fmt.Printf("\nYou can test with method POST at this link: http://127.0.0.1:3000/beef/summary ")
				fmt.Printf("\nExample Body: {\"message\": \"Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.\"}\n\n")
				err := app.Listen(":3000")
				if err != nil {
					panic(err)
				}

			// Exit: App
			case 3:
				fmt.Println("Exiting the program...")
				os.Exit(0)
			}
		default:
		}
	}
}

func printMenu(options []string, selected int) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Select an option with [↑] and [↓]:")
	for i, option := range options {
		if i == selected {
			fmt.Printf("> %s\n", option)
		} else {
			fmt.Printf("  %s\n", option)
		}
	}
}

func selectMenu() {
	var input string
	fmt.Println("[m] -> menu")
	fmt.Println("[q] -> exit program...")
	fmt.Print("Enter menu: ")
	if scanner.Scan() {
		input = scanner.Text()
	}
	switch input {
	case "m", "M":
	case "q", "Q":
		fmt.Println("Exiting the program...")
		os.Exit(0)
	}
}
