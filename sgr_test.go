package aec

import "testing"

func BenchmarkSGR(b *testing.B) {
	b.ReportAllocs()

	var a ANSI
	for i := 0; i < b.N; i++ {
		a = Color3BitF(3)
		a = Color3BitB(5)
		a = Color8BitF(128)
		a = Color8BitB(255)
		a = FullColorF(255, 128, 0)
		a = FullColorB(0, 64, 255)

		// Prevent compiler from optimizing away
		if a.String() == "" {
			b.Fatal("unexpected empty ANSI")
		}
	}
}
