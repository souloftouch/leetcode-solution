package leetcode_top_interview_150

import (
	"testing"
)

func TestRemove(t *testing.T) {
	rs := Constructor()

	// Test removing from an empty set
	if rs.Remove(1) {
		t.Errorf("Expected false, got true")
	}

	// Insert elements
	rs.Insert(1)
	rs.Insert(2)
	rs.Insert(3)

	// Test removing an element that exists
	if !rs.Remove(2) {
		t.Errorf("Expected true, got false")
	}

	// Test removing the same element again
	if rs.Remove(2) {
		t.Errorf("Expected false, got true")
	}

	// Test removing another element
	if !rs.Remove(1) {
		t.Errorf("Expected true, got false")
	}

	// Test removing the last element
	if !rs.Remove(3) {
		t.Errorf("Expected true, got false")
	}

	// Test removing from an empty set again
	if rs.Remove(3) {
		t.Errorf("Expected false, got true")
	}
}
