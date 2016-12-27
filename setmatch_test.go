package main

import (
	"testing"
)

func TestRemaining(t *testing.T) {
	shape := remaining([]byte{wave, pill, diamond}, wave, pill)
	if shape != diamond {
		t.Errorf("expected %v got %v\n", diamond, shape)
	}
	shape = remaining([]byte{wave, pill, diamond}, wave, diamond)
	if shape != pill {
		t.Errorf("expected %v got %v", pill, shape)
	}
}

func TestFindSets(t *testing.T) {
	cards := []card{
		card{wave, two, purple, hollow},
		card{wave, one, purple, hollow},
		card{wave, three, purple, hollow},
	}
	sets := findSets(cards)
	expectedSet := map[[3]int]bool {
		[3]int{0, 1, 2}: true,
	}
	if len(sets) != 1 || !sets[[3]int{0, 1, 2}] {
		t.Errorf("expected set %v, got %v", expectedSet, sets)
	}
}
