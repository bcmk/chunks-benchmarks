package chunks

import (
	"math/rand"
	"reflect"
	"testing"
)

var s1 = make([]string, 450)
var s2 = make([]string, 5000)
var s3 = make([]string, 15000)

func randomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func DoBenchmarks(f func(s []string, chunkSize int) [][]string) {
	f(s1, 20)
	f(s2, 20)
	f(s3, 20)
	f(s1, 200)
	f(s2, 200)
	f(s3, 200)
	f(s1, 1000)
	f(s2, 1000)
	f(s3, 1000)
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(chunks)
	}
	b.ReportAllocs()
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(chunks2)
	}
	b.ReportAllocs()
}
func TestEquality(t *testing.T) {
	if !reflect.DeepEqual(chunks(s1, 100), chunks2(s1, 100)) {
		t.Error()
	}
	if !reflect.DeepEqual(chunks(s2, 100), chunks2(s2, 100)) {
		t.Error()
	}
	if !reflect.DeepEqual(chunks(s3, 100), chunks2(s3, 100)) {
		t.Error()
	}
}
