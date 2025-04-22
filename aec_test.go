package aec

import (
	"testing"
)

func BenchmarkANSIGenerators(b *testing.B) {
	b.ReportAllocs()

	var result ANSI

	for i := 0; i < b.N; i++ {
		result = Up(3)
		result = Down(3)
		result = Left(3)
		result = Right(3)
		result = NextLine(2)
		result = PreviousLine(2)
		result = Column(10)
		result = Position(5, 15)
		result = EraseDisplay(EraseModes.All)
		result = EraseLine(EraseModes.Head)
		result = ScrollUp(5)
		result = ScrollDown(5)

		// Prevent compiler from optimizing away
		if result.String() == "" {
			b.Fatal("unexpected empty result")
		}
	}
}
