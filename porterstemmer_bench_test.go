package porterstemmer

import "testing"

func BenchmarkStemString(b *testing.B) {
	words := []string{"generalizations", "running", "caresses", "hopefulness", "relational", "controlling", "a", "sky"}
	b.ReportAllocs()
	for b.Loop() {
		for _, w := range words {
			StemString(w)
		}
	}
}
