package main

import (
	"fmt"
	"time"
)

//Available action in environment :
//1. ping
//2. time
//3. status
//4. exit (keluar dari environment)


/* ----- Kekurangan : Tidak menampilkan HTTP Status ----- */

type Development struct { //membuat struct development environment
	DEVELOPMENT_PORT int
	DEVELOPMENT_APP_NAME string
	DEVELOPMENT_DEBUG_MODE bool
}

type Production struct { //membuat struct production environment
	PRODUCTION_PORT int
	PRODUCTION_APP_NAME string
	PRODUCTION_DEBUG_MODE bool
}

func getTime() { //mengeprint waktu dan tanggal local
	dt := time.Now()
	fmt.Println("Current Date (dd/mm/YYYY): ", dt.Format("02-1-2006"))
	fmt.Printf("Current time: %s\n", dt.Format("15:04:05"))
}

func developmentEnvironment(dev Development, action string) { //function untuk development environment
	
	fmt.Println("")
	fmt.Println("Entering Development Environment...")
	for {
		fmt.Print("/Development/")
		fmt.Scanf("%s\n", &action)

		if action == "" {
			continue
		} else if action == "ping" { //memberikan return pong
			fmt.Println("pong")
		} else if action == "time" { //memberikan return current time
			getTime()
		} else if action == "status" { //memberikan status dari development environment
			fmt.Printf("Port: %d\nApp Name: %s\n", dev.DEVELOPMENT_PORT, dev.DEVELOPMENT_APP_NAME)
		} else if action == "exit" { //keluar dari environment
			fmt.Println("Leaving Development Environment")
			break
		}else { //handle error jika action yang dimasukkan invalid
			fmt.Printf("No action called %s!\n", action)
		}
	}
}

func productionEnvironment(prod Production, action string) {
	fmt.Println("Entering Production Environment...")
	for {
		fmt.Print("/Production/")
		fmt.Scanf("%s\n", &action)

		if action == "" {
			continue
		} else if action == "ping" { //memberikan return pong
			fmt.Println("pong")
		} else if action == "time" { //memberikan return current time
			getTime()
		} else if action == "status" { //memberikan status dari production environment
			fmt.Printf("Port: %d\nApp Name: %s\n", prod.PRODUCTION_PORT, prod.PRODUCTION_APP_NAME)
		} else if action == "exit" { //keluar dari environment
			fmt.Println("Leaving Production Environment")
			break
		}else { //handle error jika action yang dimasukkan invalid
			fmt.Printf("No action called %s!\n", action)
		}
	}
}


func main() {
	fmt.Println("")
	development := Development{1234, "Development_APP", true} //membuat variable development berdasarkan struct Development
	production := Production{2345, "Production_APP", false} //membuat variable development berdasarkan struct Production
	exit_environment := 0
	
	for ;exit_environment == 0; { //looping selama pilihan 3 tidak dipilih
		fmt.Println("Choose your environment: ")
		fmt.Println("1. Development")
		fmt.Println("2. Production")
		fmt.Println("3. Exit")
		var environment_choice int
		fmt.Print("/")
		fmt.Scanf("%d\n", &environment_choice)
		switch environment_choice {
		case 1: //masuk kedalam development environment
			developmentEnvironment(development, "")

		case 2: //masuk kedalam production environment
			productionEnvironment(production, "")

		case 3: //keluar dari looping
			fmt.Println("Goodbye!")
			exit_environment = 1

		default:
			fmt.Println("Invalid Environment type !!!")
		}
		fmt.Println("")
	}
}