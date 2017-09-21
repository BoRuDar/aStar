package main

import (
	"testing"
)

func TestGetH(t *testing.T) {

	a := Point{x:1, y:2}
	b := Point{x:2, y:3}

	expect := 2
	got := GetH(a, b)

	if got != expect {
		t.Errorf("Expected: %v, got %v",expect,  got)
	}
}

func TestPoint_GetG_14(t *testing.T) {

	a := Point{x:1, y:2}
	b := Point{x:2, y:3}

	expect := 14
	got := a.GetG(b)

	if got != expect {
		t.Errorf("Expected: %v, got %v",expect,  got)
	}
}

func TestPoint_GetG_10(t *testing.T) {

	a := Point{x:1, y:2}
	b := Point{x:2, y:2}

	expect := 10
	got := a.GetG(b)

	if got != expect {
		t.Errorf("Expected: %v, got %v",expect,  got)
	}
}

func TestPoint_GetF(t *testing.T) {
	p := Point{G:10, H:20}

	expect := 30
	got := p.GetF()

	if got != expect {
		t.Errorf("Expected: %v, got %v",expect,  got)
	}
}

func TestPoint_Key(t *testing.T) {
	p := Point{x:11, y:22}

	expect := "11:22"
	got := p.Key()

	if got != expect {
		t.Errorf("Expected: %v, got %v",expect,  got)
	}
}

func TestPoint_Equal_true(t *testing.T) {
	p1 := Point{x:1, y:2}
	p2 := Point{x:1, y:2}

	expect := true
	got := p1.Equal(p2)

	if got != expect {
		t.Errorf("Expected: %v, got %v",expect,  got)
	}
}

func TestPoint_Equal_false(t *testing.T) {
	p1 := Point{x:1, y:2}
	p2 := Point{x:1, y:1}

	expect := false
	got := p1.Equal(p2)

	if got != expect {
		t.Errorf("Expected: %v, got %v",expect,  got)
	}
}