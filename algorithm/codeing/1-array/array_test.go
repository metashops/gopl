package main

import (
	`testing`
)

func TestRemoveElement(t *testing.T) {
	if ans := RemoveElement([]int{1, 3, 4, 5, 6, 7, 8}, 3); ans != 6 {
		t.Errorf("expected 1, got %d", ans)
	}

}

func TestSearch(t *testing.T) {
	if ans := Search([]int{1, 3, 4, 5, 6, 7, 8}, 3); ans != 1 {
		t.Errorf("expected 1, got %d", ans)
	}
}

func TestFindKthLargestK(t *testing.T) {
	if ans := findKthLargestK([]int{3, 6, 10, 2, 15}, 2); ans != 10 {
		t.Errorf("expected 10, got %v", ans)
	}
}
