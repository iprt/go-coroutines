package main

import "fmt"

// f
//
//	@Description:
//	@param from
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("go coroutines")

}
