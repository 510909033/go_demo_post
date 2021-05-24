package my_json

import (
	"testing"
)

func BenchmarkReflect_DemoMyJson1(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DemoMyJson1()
	}
}

func BenchmarkReflect_DemoMyJson2(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DemoMyJson2()
	}
}
