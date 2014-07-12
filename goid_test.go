package goid

import (
	"testing"
)

func TestSimpleGenerate(t *testing.T) {
	values := make([]byte, 4)
	values[0] = '0'
	values[1] = '1'
	values[2] = '2'
	values[3] = 'A'

	for i := 0; i < 4; i++ {
		if GenerateNext(values, 0, 0) != "0" {
			t.Fatalf("Got %s instead of 1", GenerateNext(values, 0, 0))
		}
		if GenerateNext(values, 1, 0) != "1" {
			t.Fatalf("Got %s instead of 1", GenerateNext(values, 1, 0))
		}
		if GenerateNext(values, 2, 0) != "2" {
			t.Fatalf("Got %s instead of 1", GenerateNext(values, 2, 0))
		}
		if GenerateNext(values, 3, 0) != "A" {
			t.Fatalf("Got %s instead of A", GenerateNext(values, 3, 0))
		}
		if GenerateNext(values, 4, 0) != "01" {
			t.Fatalf("Got %s instead of 01", GenerateNext(values, 4, 0))
		}
		if GenerateNext(values, 5, 0) != "11" {
			t.Fatalf("Got %s instead of 11", GenerateNext(values, 5, 0))
		}
		if GenerateNext(values, 6, 0) != "21" {
			t.Fatalf("Got %s instead of 21", GenerateNext(values, 6, 0))
		}
		if GenerateNext(values, 7, 0) != "A1" {
			t.Fatalf("Got %s instead of A1", GenerateNext(values, 7, 0))
		}
		if GenerateNext(values, 8, 0) != "02" {
			t.Fatalf("Got %s instead of 02", GenerateNext(values, 8, 0))
		}
		if GenerateNext(values, 9, 0) != "12" {
			t.Fatalf("Got %s instead of 12", GenerateNext(values, 9, 0))
		}
		if GenerateNext(values, 10, 0) != "22" {
			t.Fatalf("Got %s instead of 22", GenerateNext(values, 10, 0))
		}
		if GenerateNext(values, 11, 0) != "A2" {
			t.Fatalf("Got %s instead of A2", GenerateNext(values, 11, 0))
		}
		if GenerateNext(values, 12, 0) != "0A" {
			t.Fatalf("Got %s instead of 0A", GenerateNext(values, 12, 0))
		}
		if GenerateNext(values, 13, 0) != "1A" {
			t.Fatalf("Got %s instead of 1A", GenerateNext(values, 13, 0))
		}
		if GenerateNext(values, 14, 0) != "2A" {
			t.Fatalf("Got %s instead of 2A", GenerateNext(values, 14, 0))
		}
		if GenerateNext(values, 15, 0) != "AA" {
			t.Fatalf("Got %s instead of AA", GenerateNext(values, 15, 0))
		}
		if GenerateNext(values, 16, 0) != "001" {
			t.Fatalf("Got %s instead of 001", GenerateNext(values, 16, 0))
		}
		if GenerateNext(values, 17, 0) != "101" {
			t.Fatalf("Got %s instead of A3", GenerateNext(values, 17, 0))
		}
		if GenerateNext(values, 18, 0) != "201" {
			t.Fatalf("Got %s instead of 201", GenerateNext(values, 18, 0))
		}
		if GenerateNext(values, 19, 0) != "A01" {
			t.Fatalf("Got %s instead of A01", GenerateNext(values, 19, 0))
		}
		if GenerateNext(values, 20, 0) != "011" {
			t.Fatalf("Got %s instead of 011", GenerateNext(values, 20, 0))
		}
		if GenerateNext(values, 21, 0) != "111" {
			t.Fatalf("Got %s instead of 111", GenerateNext(values, 21, 0))
		}
		if GenerateNext(values, 22, 0) != "211" {
			t.Fatalf("Got %s instead of 211", GenerateNext(values, 22, 0))
		}
	}

	g := GenerateNext(values, 5, 4);
	if len(g) < 4 {
		t.Fatalf("Got %s instead something which should have 4 digits", g)
	}
	g = GenerateNext(values, 55, 4);
	if len(g) < 4 {
		t.Fatalf("Got %s instead something which should have 4 digits", g)
	}
	g = GenerateNext(values, 43432423435, 4);
	if len(g) < 4 {
		t.Fatalf("Got %s instead something which should have 4 digits", g)
	}
	g = GenerateNext(values, 5, 12);
	if len(g) < 12 {
		t.Fatalf("Got %s instead something which should have 12 digits", g)
	}

	g = GenerateNext(values, 0, 6);
	if g != "000000" {
		t.Fatalf("Should generate 000000 and not : %s", g);
	}
}
