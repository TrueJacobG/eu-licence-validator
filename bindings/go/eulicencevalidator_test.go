package eulicencevalidator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type testCase struct {
	Plate    string `json:"plate"`
	Country  string `json:"country"`
	Expected bool   `json:"expected"`
}

func loadTestCases(t *testing.T) []testCase {
	t.Helper()
	path := filepath.Join("..", "..", "test_cases.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	var cases []testCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("failed to parse %s: %v", path, err)
	}
	if len(cases) == 0 {
		t.Fatal("test_cases.json is empty")
	}
	return cases
}

func TestIsValid(t *testing.T) {
	for _, tc := range loadTestCases(t) {
		got := IsValid(tc.Plate, tc.Country)
		if got != tc.Expected {
			t.Errorf("IsValid(%q, %q) = %v, want %v", tc.Plate, tc.Country, got, tc.Expected)
		}
	}
}
