package main

import (
	"fmt"
	"math"
	"net/http"
)

func obterRaizQuadrada(x float64) float64 {
	return math.Sqrt(x)
}

func run(w http.ResponseWriter, r *http.Request) {
	var x float64 = 0.0001
	for i := 0; i < 1000000; i++ {
		x += obterRaizQuadrada(x)
	}
	fmt.Fprintf(w, "Code.education Rocks!")
}

func main() {
	http.HandleFunc("/", run)
	http.ListenAndServe(":8000", nil)
}
