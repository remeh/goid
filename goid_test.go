package goid

import "testing"

func TestSimpleGenerate(t *testing.T) {
	values := make([]byte, 0)
	values = append(values, '0')
	values = append(values, '1')
	values = append(values, '2')
	values = append(values, 'A')

	for i := 0; i < 4; i++ {
		if GenerateNext(values, 0) != "0" {
			t.Fatalf("Got %s instead of 1", GenerateNext(values, 0))
		}
		if GenerateNext(values, 1) != "1" {
			t.Fatalf("Got %s instead of 1", GenerateNext(values, 1))
		}
		if GenerateNext(values, 2) != "2" {
			t.Fatalf("Got %s instead of 1", GenerateNext(values, 2))
		}
		if GenerateNext(values, 3) != "A" {
			t.Fatalf("Got %s instead of A", GenerateNext(values, 3))
		}
		if GenerateNext(values, 4) != "01" {
			t.Fatalf("Got %s instead of 01", GenerateNext(values, 4))
		}
		if GenerateNext(values, 5) != "11" {
			t.Fatalf("Got %s instead of 11", GenerateNext(values, 5))
		}
		if GenerateNext(values, 6) != "21" {
			t.Fatalf("Got %s instead of 21", GenerateNext(values, 6))
		}
		if GenerateNext(values, 7) != "A1" {
			t.Fatalf("Got %s instead of A1", GenerateNext(values, 7))
		}
		if GenerateNext(values, 8) != "02" {
			t.Fatalf("Got %s instead of 02", GenerateNext(values, 8))
		}
		if GenerateNext(values, 9) != "12" {
			t.Fatalf("Got %s instead of 12", GenerateNext(values, 9))
		}
		if GenerateNext(values, 10) != "22" {
			t.Fatalf("Got %s instead of 22", GenerateNext(values, 10))
		}
		if GenerateNext(values, 11) != "A2" {
			t.Fatalf("Got %s instead of A2", GenerateNext(values, 11))
		}
		if GenerateNext(values, 12) != "0A" {
			t.Fatalf("Got %s instead of 0A", GenerateNext(values, 12))
		}
		if GenerateNext(values, 13) != "1A" {
			t.Fatalf("Got %s instead of 1A", GenerateNext(values, 13))
		}
		if GenerateNext(values, 14) != "2A" {
			t.Fatalf("Got %s instead of 2A", GenerateNext(values, 14))
		}
		if GenerateNext(values, 15) != "AA" {
			t.Fatalf("Got %s instead of AA", GenerateNext(values, 15))
		}
		if GenerateNext(values, 16) != "001" {
			t.Fatalf("Got %s instead of 001", GenerateNext(values, 16))
		}
		if GenerateNext(values, 17) != "101" {
			t.Fatalf("Got %s instead of A3", GenerateNext(values, 17))
		}
		if GenerateNext(values, 18) != "201" {
			t.Fatalf("Got %s instead of 201", GenerateNext(values, 18))
		}
		if GenerateNext(values, 19) != "A01" {
			t.Fatalf("Got %s instead of A01", GenerateNext(values, 19))
		}
		if GenerateNext(values, 20) != "011" {
			t.Fatalf("Got %s instead of 011", GenerateNext(values, 20))
		}
		if GenerateNext(values, 21) != "111" {
			t.Fatalf("Got %s instead of 111", GenerateNext(values, 21))
		}
		if GenerateNext(values, 22) != "211" {
			t.Fatalf("Got %s instead of 211", GenerateNext(values, 22))
		}
	}
}
