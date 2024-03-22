package main

import "testing"

func TestCode(t *testing.T){
	var result = 100
	if result != 2 {
		t.Error("Excepted 2, got ", result)
		
	}
}