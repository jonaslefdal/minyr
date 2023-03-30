package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"minyr/yr"
	"github.com/jonaslefdal/funtemps/conv"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()

		switch input {
		case "q", "exit":
			fmt.Println("exit")
			os.Exit(0)
		case "convert":
			_, err := os.Stat("/home/BRUKERNAVN/minyr/kjevik-temp-fahr-20220318-20230318.csv")
			if os.IsNotExist(err) {
				convertCelsiusToFahrenheit()
			} else {
				fmt.Println("The output file already exists. Are you sure you want to convert all measurements given in degrees Celsius to degrees Fahrenheit? (y/n)")
				scanner.Scan()
				input = scanner.Text()
				if strings.ToLower(input) == "y" {
					convertCelsiusToFahrenheit()
				} else {
					fmt.Println("Conversion canceled.")
				}
			}
		case "average":
			averageTemp()
		default:
			fmt.Println("Please select convert, average or exit:")
		}
	}
}

func convertCelsiusToFahrenheit() {
	fmt.Println("Converting all measurements given in degrees Celsius to degrees Fahrenheit.")
	src, err := os.Open("/home/BRUKERNAVN/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create("/home/BRUKERNAVN/minyr/kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dst)

	if scanner.Scan() {
		line1 := (scanner.Text())
		fmt.Fprintln(writer, line1)
	}

	for scanner.Scan() {
		line := scanner.Text()
		fahrLine, err := yr.CelsiusToFahrenheitLine(line)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Fprintln(writer, fahrLine)
	}

	writer.Flush()
	fmt.Println("Conversion completed successfully.\nResults saved in kjevik-temp-fahr-20220318-20230318.csv.")
}

func averageTemp() {
	fmt.Println("Please select in degrees Celsius or Fahrenheit? (c/f)")
	count := 0
	sum := 0

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()

			if strings.ToLower(input) == "c" {
				fmt.Println("Finding the average temp in Celsius")
				avg := yr.AverageTemp(sum, float64(count))
				fmt.Printf("Average: %.2f\n", avg)
				fmt.Println("Please select (c/f) for new average or (q/exit) to exit")

			} else if strings.ToLower(input) == "f" {

				fmt.Println("Finding the average temp in Fahrenheit")
				// Calculate the average
				avg := yr.AverageTemp(sum, float64(count))
				avgFahr := conv.CelsiusToFahrenheit(avg)
				fmt.Printf("Average: %.2f\n", avgFahr)

			} else {
				fmt.Println("Please select (c/f) or (q/exit)")
			}
			if strings.ToLower(input) == "q" || strings.ToLower(input) == "exit" {
				fmt.Println("exit")
				os.Exit(0)
		}
	}
}
