package passphrase

import (
	"testing"
)

func BenchmarkGeneratePassphrase(b *testing.B) {
	b.StopTimer()
	pg := New(NONE, "dict/wordsEn.txt")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		pg.Generate(4)
	}
}
