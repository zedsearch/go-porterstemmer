package porterstemmer

import "testing"

func refIsConsonant(s []rune, i int) bool {
	switch s[i] {
	case 'a', 'e', 'i', 'o', 'u':
		return false
	case 'y':
		return i == 0 || !refIsConsonant(s, i-1)
	}
	return true
}

func refMeasure(s []rune) uint {
	var m uint
	for i := 1; i < len(s); i++ {
		if refIsConsonant(s, i) && !refIsConsonant(s, i-1) {
			m++
		}
	}
	return m
}

func TestConsonantMeasureMatchReference(t *testing.T) {
	alphabet := []rune("aeybcy")
	var gen func(prefix []rune, n int)
	gen = func(prefix []rune, n int) {
		s := append([]rune(nil), prefix...)
		for i := range s {
			if got, want := isConsonant(s, i), refIsConsonant(s, i); got != want {
				t.Fatalf("isConsonant(%q, %d) = %v, want %v", string(s), i, got, want)
			}
		}
		if got, want := measure(s), refMeasure(s); got != want {
			t.Fatalf("measure(%q) = %d, want %d", string(s), got, want)
		}
		if got, want := containsVowel(s), func() bool {
			for i := range s {
				if !refIsConsonant(s, i) {
					return true
				}
			}
			return false
		}(); got != want {
			t.Fatalf("containsVowel(%q) = %v, want %v", string(s), got, want)
		}
		if n == 0 {
			return
		}
		for _, r := range alphabet {
			gen(append(prefix, r), n-1)
		}
	}
	gen(nil, 7)
}

func TestLongYRun(t *testing.T) {
	s := make([]rune, 100000)
	for i := range s {
		s[i] = 'y'
	}
	if isConsonant(s, len(s)-1) != (len(s)%2 == 1) {
		t.Fatal("wrong alternation for long y run")
	}
	_ = StemString(string(s))
}
