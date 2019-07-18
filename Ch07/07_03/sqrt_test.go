package sqrt

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation - %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

type testCase struct {
	value    float64
	expected float64
}

func TestMany(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatalf("error: can't open CSV file - %s", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error: can't read csv - %s", err)
	}

	testCases := []testCase{}

	for _, v := range records {
		v0, _ := strconv.ParseFloat(v[0], 64)
		v1, _ := strconv.ParseFloat(v[1], 64)
		testCases = append(testCases, testCase{v0, v1})
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatal("error")
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}
