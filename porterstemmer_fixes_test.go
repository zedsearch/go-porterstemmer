package porterstemmer

import (
	"testing"
)

// Test for issue listed here:
// https://github.com/reiver/go-porterstemmer/issues/1
//
// StemString("ion") was causing runtime exception
func TestStemStringIon(t *testing.T) {

	expected := "ion"

	s := "ion"
	actual := StemString(s)
	if expected != actual {
		t.Errorf("Input: [%s] -> Actual: [%s]. Expected: [%s]", s, actual, expected)
	}
}

// Test for issue listed here:
// https://github.com/reiver/go-porterstemmer/pull/10
//
// StemString("eeg") was causing runtime exception
func TestStemStringEeg(t *testing.T) {

	expected := "eeg"

	s := "eeg"
	actual := StemString(s)
	if expected != actual {
		t.Errorf("Input: [%s] -> Actual: [%s]. Expected: [%s]", s, actual, expected)
	}
}

// step5a on a lone "e" leaves an empty subSlice, which must not be indexed.
func TestStep5aLoneE(t *testing.T) {
	if actual := string(step5a([]rune("e"))); actual != "e" {
		t.Errorf("step5a(\"e\") = [%s]. Expected: [e]", actual)
	}
}
