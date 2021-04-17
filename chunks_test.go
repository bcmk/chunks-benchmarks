package chunks

import (
	"reflect"
	"testing"
)

func TestChunks(t *testing.T) {
	xs := []string{"a", "b", "c", "d", "e", "f"}
	ys := chunks(xs, 1)
	zs := [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}}
	if !reflect.DeepEqual(ys, zs) {
		t.Errorf("%v is not equal to %v", ys, zs)
	}
	ys = chunks(xs, 2)
	zs = [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	if !reflect.DeepEqual(ys, zs) {
		t.Errorf("%v is not equal to %v", ys, zs)
	}
	ys = chunks(xs, 4)
	zs = [][]string{{"a", "b", "c", "d"}, {"e", "f"}}
	if !reflect.DeepEqual(ys, zs) {
		t.Errorf("%v is not equal to %v", ys, zs)
	}
	ys = chunks(xs, 5)
	zs = [][]string{{"a", "b", "c", "d", "e"}, {"f"}}
	if !reflect.DeepEqual(ys, zs) {
		t.Errorf("%v is not equal to %v", ys, zs)
	}
	ys = chunks(xs, 6)
	zs = [][]string{{"a", "b", "c", "d", "e", "f"}}
	if !reflect.DeepEqual(ys, zs) {
		t.Errorf("%v is not equal to %v", ys, zs)
	}
	ys = chunks(xs, 7)
	zs = [][]string{{"a", "b", "c", "d", "e", "f"}}
	if !reflect.DeepEqual(ys, zs) {
		t.Errorf("%v is not equal to %v", ys, zs)
	}
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
