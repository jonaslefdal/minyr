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
			fmt.Println("Please select in degrees Celsius or Fahrenheit? (c/f)")
			scanner.Scan()
			input = scanner.Text()
			if strings.ToLower(input) == "c" {
				averageTempCelcsius()
			} else if strings.ToLower(input) == "f" {
				averageTempFahr()
			}
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

func averageTempCelcsius() {
	fmt.Println("Finding the average temp in Celsius")
	
	count := 0
	sum := 0

	// Calculate the average
	avg := yr.AverageTemp(sum, float64(count))

	fmt.Printf("Average: %.2f\n", avg)
}
func averageTempFahr() {
	fmt.Println("Finding the average temp in Fahrenheit")
	
	count := 0
	sum := 0

	// Calculate the average
	avg := yr.AverageTemp(sum, float64(count))
	avgFahr := conv.CelsiusToFahrenheit(avg)

	fmt.Printf("Average: %.2f\n", avgFahr)
}
