package main

import (
	"fmt"
	"math"
	"net/http"
)

func greeting(str string) string {
	x := 0.0001

	for i := 0; i < 100000; i++ {
		x = math.Sqrt(x)
	}

	return fmt.Sprintf("<b>%s</b>", str)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", greet)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", nil)
}
