package timestamp_test

import (
	"testing"
	"time"

	"github.com/jwilder/encoding/timestamp"
)

func Test_MarshalBinary(t *testing.T) {
	enc := timestamp.NewEncoder()

	x := []time.Time{}
	now := time.Unix(0, 0)
	x = append(x, now)
	enc.Write(now)
	for i := 1; i < 4; i++ {
		x = append(x, now.Add(time.Duration(i)*time.Second))
		enc.Write(x[i])
	}

	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dec := timestamp.NewDecoder(b)
	for i, v := range x {
		if !dec.Next() {
			t.Fatalf("Next == false, expected true")
		}

		if v != dec.Read() {
			t.Fatalf("Item %d mismatch, got %v, exp %v", i, dec.Read(), v)
		}
	}
}

func Test_MarshalBinary_NoValues(t *testing.T) {
	enc := timestamp.NewEncoder()
	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dec := timestamp.NewDecoder(b)
	if dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}
}

func Test_MarshalBinary_One(t *testing.T) {
	enc := timestamp.NewEncoder()
	tm := time.Unix(0, 0)

	enc.Write(tm)
	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dec := timestamp.NewDecoder(b)
	if !dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}

	if tm != dec.Read() {
		t.Fatalf("read value mismatch: got %v, exp %v", dec.Read(), tm)
	}
}

func Test_MarshalBinary_Two(t *testing.T) {
	enc := timestamp.NewEncoder()
	t1 := time.Unix(0, 0)
	t2 := time.Unix(0, 1)
	enc.Write(t1)
	enc.Write(t2)

	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dec := timestamp.NewDecoder(b)
	if !dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}

	if t1 != dec.Read() {
		t.Fatalf("read value mismatch: got %v, exp %v", dec.Read(), t1)
	}

	if !dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}

	if t2 != dec.Read() {
		t.Fatalf("read value mismatch: got %v, exp %v", dec.Read(), t2)
	}
}

func Test_MarshalBinary_Three(t *testing.T) {
	enc := timestamp.NewEncoder()
	t1 := time.Unix(0, 0)
	t2 := time.Unix(0, 1)
	t3 := time.Unix(0, 2)

	enc.Write(t1)
	enc.Write(t2)
	enc.Write(t3)

	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dec := timestamp.NewDecoder(b)
	if !dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}

	if t1 != dec.Read() {
		t.Fatalf("read value mismatch: got %v, exp %v", dec.Read(), t1)
	}

	if !dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}

	if t2 != dec.Read() {
		t.Fatalf("read value mismatch: got %v, exp %v", dec.Read(), t2)
	}

	if !dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}

	if t3 != dec.Read() {
		t.Fatalf("read value mismatch: got %v, exp %v", dec.Read(), t3)
	}
}

func Test_Encode_Large_Range(t *testing.T) {
	enc := timestamp.NewEncoder()
	t1 := time.Unix(0, 1442369134000000000)
	t2 := time.Unix(0, 1442369135000000000)
	enc.Write(t1)
	enc.Write(t2)
	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dec := timestamp.NewDecoder(b)
	if !dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}

	if t1 != dec.Read() {
		t.Fatalf("read value mismatch: got %v, exp %v", dec.Read(), t1)
	}

	if !dec.Next() {
		t.Fatalf("unexpected next value: got true, exp false")
	}

	if t2 != dec.Read() {
		t.Fatalf("read value mismatch: got %v, exp %v", dec.Read(), t2)
	}
}
