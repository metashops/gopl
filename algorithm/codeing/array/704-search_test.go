package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	if ans := Search([]int{1, 3, 4, 5, 6, 7, 8}, 3); ans != 1 {
		t.Errorf("expected 1, got %d", ans)
	}

}
