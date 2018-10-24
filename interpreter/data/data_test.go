package data

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if HashData(hello1) != HashData(hello2) {
		t.Errorf("Expected same hash for strings with the same content")
	}

	if HashData(diff1) != HashData(diff2) {
		t.Errorf("Expected same hash for strings with the same content")
	}

	if HashData(hello1) == HashData(diff1) {
		t.Errorf("Expected same hash for strings with the same content")
	}
}
