// add_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	if res := Add(2, 2); res != 4 {
		t.Fatalf("Expected 4, got %d", res)
	}
}
