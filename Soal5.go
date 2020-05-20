package main

import (
	"fmt"
	"time"
)

type Development struct {
	DEVELOPMENT_PORT int
	DEVELOPMENT_APP_NAME string
	DEVELOPMENT_DEBUG_MODE bool
}

type Production struct {
	PRODUCTION_PORT int
	PRODUCTION_APP_NAME string
	PRODUCTION_DEBUG_MODE bool
}

func main() {
	fmt.Println("")
	development := Development{1234, "Development_APP", true}
	production := Production{2345, "Production_APP", false}
	exit_environment := 0
	
	for ;exit_environment == 0; {
		fmt.Println("Choose your environment: ")
		fmt.Println("1. Development")
		fmt.Println("2. Production")
		fmt.Println("3. Exit")
		var environment_choice int
		fmt.Print("/")
		fmt.Scanf("%d\n", &environment_choice)
		var action string
		switch environment_choice {
		case 1:
			fmt.Println("")
			fmt.Println("Entering Development Environment...")
			for {
				fmt.Print("/Development/")
				fmt.Scanf("%s\n", &action)

				if action == "" {
					continue
				} else if action == "ping" {
					fmt.Println("pong")
				} else if action == "time" {
					dt := time.Now()
					fmt.Println("Current Date (dd/mm/YYYY): ", dt.Format("02-1-2006"))
					fmt.Printf("Current time: %s WIB\n", dt.Format("15:04:05"))
				} else if action == "status" {
					fmt.Printf("Port: %d\nApp Name: %s\n", development.DEVELOPMENT_PORT, development.DEVELOPMENT_APP_NAME)
				} else if action == "exit" {
					fmt.Println("Leaving Development Environment")
					break
				}else {
					fmt.Printf("No action called %s!\n", action)
				}

			}

		case 2:
			fmt.Println("Entering Production Environment...")
			for {
				fmt.Print("/")
				fmt.Scanf("%s\n", &action)

				if action == "" {
					continue
				} else if action == "ping" {
					fmt.Println("pong")
				} else if action == "time" {
					fmt.Println("waktu")
				} else if action == "status" {
					fmt.Printf("Port: %d\nApp Name: %s\n", production.PRODUCTION_PORT, production.PRODUCTION_APP_NAME)
				} else if action == "exit" {
					fmt.Println("Leaving Development Environment")
					break
				}else {
					fmt.Printf("No action called %s!\n", action)
				}

			}


		case 3:
			fmt.Println("Goodbye!")
			exit_environment = 1

		default:
			fmt.Println("Invalid Environment type !!!")
		}
		fmt.Println("")
		action = ""
	}
}