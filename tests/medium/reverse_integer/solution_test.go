package reverse_integer

import (
	"testing"
)

func Test_reverse_1(t *testing.T) {
	t.SkipNow()
	got := reverse(123)
	expected := 321
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_reverse_2(t *testing.T) {
	t.SkipNow()
	got := reverse(-123)
	expected := -321
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_reverse_3(t *testing.T) {
	t.SkipNow()
	got := reverse(120)
	expected := 21
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_reverse_4(t *testing.T) {
	t.SkipNow()
	got := reverse(1534236469)
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_reverse_5(t *testing.T) {
	t.SkipNow()
	got := reverse(-2147483412)
	expected := -2143847412
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_reverse_6(t *testing.T) {
	t.SkipNow()
	got := reverse(1463847412)
	expected := 2147483641
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
