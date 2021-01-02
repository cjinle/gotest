package encode

import (
	"testing"
)

func TestSth(t *testing.T) {
	// AesEncode()
	// ByteEncode()
	// BinPacks()
	// GobEncode()
	JsonEncode()
	JsonDecode()
}

func BenchmarkSth(b *testing.B) {
	for n := 0; n < b.N; n++ {
		JsonEncode()
		JsonDecode()
	}
}
