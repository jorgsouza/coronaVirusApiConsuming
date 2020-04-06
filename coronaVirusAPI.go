package main

import (
	_ "bytes"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	intro()
	fmt.Println("")
	for {
		menu()
		choose := startmenu()

		switch choose {
		case 1:
			commandCmd()
			fmt.Println("")
			submenu()
		case 2:
			commandCmd()
			fmt.Println("Showing all Coronavirus API data")
			coronaDataAllCountry()
		case 3:
			fmt.Println("Thanks for using")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			commandCmd()
			fmt.Println("You entered an invalid option")
			fmt.Println("")
		}
	}
}

func intro() {
	//
	fmt.Println("Exercise on how to extract data from an API.\nI am using an open data API on Coronavirus.")
	fmt.Println("Program Version: 1.0")
}

func menu() {
	//program menu
	fmt.Println("Choose Coronavirus data for... ")
	fmt.Println("1 - ...a specific country")
	fmt.Println("2 - ...or data from all countries")
	fmt.Println("3 - Exit")
}

func submenu() {
	fmt.Println("Which country?")
	coronaDataOnlyOne()
}

func startmenu() int {
	//collect the country name
	var requestMenu int
	fmt.Scan(&requestMenu)

	return requestMenu
}

func requestCountry() string {
	//collect the country name
	var readCountry string
	fmt.Scan(&readCountry)
	fmt.Println("Chosen country was...\n" + strings.ToUpper(readCountry))

	return readCountry
}

func coronaDataOnlyOne() {
	request := requestCountry()
	fmt.Println("Starting...\nperforming query in the API: https://coronavirus-19-api.herokuapp.com/countries/" + request)
	response, err := http.Get("https://coronavirus-19-api.herokuapp.com/countries/" + request)
	if err != nil {
		fmt.Println("The request failed \n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		meuip, err := http.Get("https://httpbin.org/ip")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(meuip.Body)
			fmt.Println("\nRequested by:", string(data))
		}
	}
	//jsonData := map[string]string{"firtname": "Nic", "lasname": "Raboy"}
	//jasonValue, _ := json.Marshal(jsonData)
	//response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jasonValue))
	//if err != nil {
	//	fmt.Println("The request failed \n", err)
	//} else {
	//	data, _ := ioutil.ReadAll(response.Body)
	//	fmt.Println(string(data))
	//}
	fmt.Println("\nClosing connection with API...")
}

func coronaDataAllCountry() {
	fmt.Println("Starting...\nperforming query in the API: https://coronavirus-19-api.herokuapp.com/countries/")
	response, err := http.Get("https://coronavirus-19-api.herokuapp.com/countries")
	if err != nil {
		fmt.Println("The request failed \n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		meuip, err := http.Get("https://httpbin.org/ip")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(meuip.Body)
			fmt.Println("\nRequested by:", string(data))
		}
	}
	//jsonData := map[string]string{"firtname": "Nic", "lasname": "Raboy"}
	//jasonValue, _ := json.Marshal(jsonData)
	//response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jasonValue))
	//if err != nil {
	//	fmt.Println("The request failed \n", err)
	//} else {
	//	data, _ := ioutil.ReadAll(response.Body)
	//	fmt.Println(string(data))
	//}
	fmt.Println("\nClosing connection with API...")
}

func commandCmd()  {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
