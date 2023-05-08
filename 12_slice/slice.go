package main

import "fmt"

// var identifier []type
// the number of the members in array is fixed, slice is dynamic
// identifier_of_array[lower:upper]
// slice can be part of array
func slice_array() {
	fmt.Println("slice_array()")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // a slice which is part of the array primes, started with the second memeber and ended with the fourth memeber
	fmt.Println(s)            // [3 5 7]
}

// slice does not store anything
// if slice is modifed, so is the underlying array, and any other slice will observe the modification
func slice_quote() {
	fmt.Println("slice_quote()")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // a slice which is part of the array primes, started with the second memeber and ended with the fourth memeber
	fmt.Println(s)            // [3 5 7]
	var s2 []int = primes[2:5]
	fmt.Println(s2) // [5 7 11]
	s[1] = 71
	fmt.Println(s)  // [3 71 7]
	fmt.Println(s2) // [3 71 7]
	fmt.Println(primes)
}

func slice_init() {
	fmt.Println("slice_init()")
	var s []int = []int{1, 2, 3}
	fmt.Println(s)

	f := []float64{1.0, 2.5, 3.6}
	fmt.Println(f)
}

// default lower bound is 0
// default upper bound is the lengh of the array
func slice_default_bound() {
	fmt.Println("slice_default_bound()")
	var s []int = []int{0, 1, 2, 3, 4, 5, 6, 7} // lengh is 8
	fmt.Println(s)
	s1 := s[0:8]
	s2 := s[0:]
	s3 := s[:8]
	s4 := s[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	// s5 := s[:9]  // "runtime error: slice bounds out of range [:9] with capacity 8"
	// fmt.Println(s5)
}

// the length of a slice is the number of the members
// the capacity of a slice is the number counting from the first member of the slice to the last member of the underlying array
func slice_len_cap() {
	fmt.Println("slice_len_cap():")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]     // the length is zero
	printSlice(s) // len=0 cap=6 []

	s = s[:4]     // extend the length
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	s = s[2:]     // drop the first two members
	printSlice(s) // len=2 cap=4 [5 7]

	s = s[:3]
	printSlice(s) // len=3 cap=4 [5 7 11]

	s = s[:4]
	printSlice(s) // len=4 cap=4 [5 7 11 13]

	// s = s[:5] // "runtime error: slice bounds out of range [:5] with capacity 4"
	// printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// the zero value of a slice is nil
// the length and cap of the slice of nil are both 0
// the slice of nil has no underlying array
func zero_slice() {
	fmt.Println("zero_slice():")
	var s []int                                         // nil slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=0 cap=0 []
}

func main() {
	slice_array()
	slice_quote()
	slice_init()
	slice_default_bound()
	slice_len_cap()
	zero_slice()
}
