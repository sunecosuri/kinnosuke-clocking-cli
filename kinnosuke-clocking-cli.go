package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"gopkg.in/headzoo/surf.v1"
)

const kinnosukeUrl string = "https://www.4628.jp/"
const ua string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"
const clockingIdIn string = "1"
const clockingIdOut string = "2"

func choise() bool {
	result := true
	fmt.Print("[Y/n]:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()

		if i == "Y" || i == "y" {
			break
		} else if i == "N" || i == "n" {
			result = false
			break
		} else {
			fmt.Println("Please answer Y or n")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}

func attendance(clockingOut *bool) {
	var clockingId string
	if *clockingOut {
		clockingId = clockingIdOut
		fmt.Println("Clocking out...")
	} else {
		clockingId = clockingIdIn
		fmt.Println("Clocking in...")
	}

	browser := surf.NewBrowser()
	browser.SetUserAgent(ua)

	err := browser.Open(kinnosukeUrl)
	if err != nil {
		panic(err)
	}

	loginForm, _ := browser.Form("[id='form1']")
	loginForm.Input("y_companycd", os.Getenv("KINNOSUKE_COMPANYCD"))
	loginForm.Input("y_logincd", os.Getenv("KINNOSUKE_LOGINCD"))
	loginForm.Input("password", os.Getenv("KINNOSUKE_PASSWORD"))
	if loginForm.Submit() != nil {
		panic(err)
	}

	timeRecorderForm, _ := browser.Form("[id='tr_submit_form']")
	timeRecorderForm.Input("timerecorder_stamping_type", clockingId)
	if timeRecorderForm.Submit() != nil {
		panic(err)
	}
}

func main() {

	answer := choise()
	if answer {
		clockingOut := flag.Bool("out", false, "Clocking out")
		flag.Parse()

		attendance(clockingOut)

		fmt.Println("Success")
	} else {
		fmt.Println("Cancel")
	}
}
