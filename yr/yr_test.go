package yr

import (
	"testing"
	"math"
)

func TestCelsiusToFahrenheitLine(t *testing.T) {

	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologiskinstitutt (MET);;;", 
		want: "Data er basert paa gyldig data (per 18.03.2023) (CCBY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Jonas Lefdal"},
	}
	for _, tc := range tests {
		got, err := CelsiusToFahrenheitLine(tc.input)
		if got != tc.want {
			t.Errorf("expected: %s, got: %s", tc.want, got)
		}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
func TestCountLines(t *testing.T) {

	type test struct {
		input string
		want  int
	}
	tests := []test{
		{input: "/home/BRUKERNAVN/minyr/kjevik-temp-fahr-20220318-20230318.csv", want: 16756},
	}
	for _, tc := range tests {
		got := CountLines(tc.input)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestAverageTemp(t *testing.T) {

	type test struct {
		sum int
		count float64
		want float64
	}
	tests := []test{
		{sum: 143397, count: 16754, want: 8.56},
	}
	for _, tc := range tests {
		got := AverageTemp(tc.sum, tc.count)
		if math.Round(got*100)/100 != tc.want {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}
