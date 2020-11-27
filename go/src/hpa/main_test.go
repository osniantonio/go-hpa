package main

import "testing"

func TestRaizQuadrada(t *testing.T) {
	result := obterRaizQuadrada(81)
	var response float64 = 9
	if result != response {
		t.Errorf("Invalid result! :( return %f, wanted %f", result, response)
	}
}