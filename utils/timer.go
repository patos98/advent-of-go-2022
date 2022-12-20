package utils

import (
	"fmt"
	"time"
)

// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
//	defer timer("sum")()
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

// Example use
//
// func main() {
//     defer timer("main")()  // <-- The trailing () is the deferred call
//     time.Sleep(time.Second * 2)
// }   // prints: main took 2s
