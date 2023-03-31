package yr

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jonaslefdal/funtemps/conv"
)

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string) (string, error) {

	dividedString := strings.Split(line, ";")
	var err error

	if len(dividedString) == 4 {

		if dividedString[3] == "" {
			return "Data er basert paa gyldig data (per 18.03.2023) (CCBY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Jonas Lefdal", err
		}
		
		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil

	//return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
}
func CountLines(input string) int {

	var fileName = input

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines

}
func AverageTemp(sum int, count float64) float64 {

	src, err := os.Open("/home/BRUKERNAVN/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		if count == 0 {
			count++
		}
		dividedString := strings.Split(scanner.Text(), ";")
		
		if dividedString[3] == "Lufttemperatur" || dividedString[3] == "" {
			continue // skip
		}

		num, err := strconv.Atoi(dividedString[3])
		if err != nil {
			 log.Fatalln(err)
		}
		sum += num
		count++
	}
	avg := float64(sum) / float64(count-2)
	return avg

}

