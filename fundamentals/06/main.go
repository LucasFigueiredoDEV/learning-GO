package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len = %d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len = %d cap=%d %v\n", len(s), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[:2]), s[:2])

	s = append(s, 12) //tomar cuidado ao utilizar append em slices, pois basicamente o Go dobra o array por trás do slice, o que acaba sendo um processo mais "doloroso".
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[:2]), s[:2])
}
