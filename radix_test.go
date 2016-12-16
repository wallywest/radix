package radix

import (
	crand "crypto/rand"
	"fmt"
	"testing"
)

func TestRadix(t *testing.T) {
	rad := New()

	out, ok := rad.Get("/")

	if out != nil {
		t.Fatalf("there should be no key")
	}
	if ok {
		t.Fatalf("there should be no key")
	}

	var min, max string

	input := make(map[string]interface{})

	for i := 0; i <= 1000; i++ {
		gen := generateUUID()
		input[gen] = i

		if gen < min || i == 0 {
			min = gen
		}

		if gen > max || i == 0 {
			max = gen
		}
	}

	//r := NewFromMap(input)

	//if r.Len() != len(input) {
	//t.Fatalf("the length should be equal: %v, %v", r.Len(), len(input))
	//}
}

func TestRoot(t *testing.T) {
	r := New()

	_, ok := r.Delete("")
	if ok {
		t.Fatalf("should not happen")
	}

	_, ok = r.Insert("", true)
	if ok {
		t.Fatalf("should not happen")
	}

	val, ok := r.Get("")
	if !ok || val != true {
		t.Fatalf("should not happen")
	}
}

func generateUUID() string {
	buf := make([]byte, 16)
	if _, err := crand.Read(buf); err != nil {
		panic(fmt.Errorf("failed to read random bytes: %v", err))
	}
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16])
}
