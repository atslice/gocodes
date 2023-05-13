package main

import (
	"fmt"
	"strings"
)

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

// The make built-in function allocates and initializes an object of type slice, map, or chan (only).
// Like new, the first argument is a type, not a value.
// Unlike new, make's return type is the same as the type of its argument, not a pointer to it.
// The specification of the result depends on the type:
// Slice: The size specifies the length. The capacity of the slice is equal to its length.
// A second integer argument may be provided to specify a different capacity;
// it must be no smaller than the length.
// For example, make([]int, 0, 10) allocates an underlying array of size 10
// and returns a slice of length 0 and capacity 10 that is backed by this underlying array.
func make_slice() {
	fmt.Println("make_slice():")
	a := make([]int, 5) // if only one integer argument is specified, make function creates an array with the int memembers with zero-values, and returns a slice that is backed by this underlying array with length and capacity of the argument int
	fmt.Println("make_slice(): a")
	printSlice(a) // len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	fmt.Println("make_slice(): b")
	printSlice(b) // len=0 cap=5 []

	c := b[:2]
	fmt.Println("make_slice(): c")
	printSlice(c) // len=2 cap=5 [0 0]

	d := c[2:5]
	fmt.Println("make_slice(): d")
	printSlice(d) // len=3 cap=3 [0 0 0]

	e := make([]int, 3, 7) // if two integer arguments are specified, the first integer is as length, the second integer is as capacity
	fmt.Println("make_slice(): e")
	printSlice(e) // len=3 cap=7 [0 0 0]

	f := e[:7]
	fmt.Println("make_slice(): f")
	printSlice(f) // len=7 cap=7 [0 0 0 0 0 0 0]

	g := e[2:5]
	fmt.Println("make_slice(): g")
	printSlice(g) // len=3 cap=5 [0 0 0]
}

func game_slice() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func slice_append() {
	fmt.Println("slice_append(): ")
	var s []int   // empty slice
	printSlice(s) // len=0 cap=0 []

	s = append(s, 0) // append a memeber to the slice
	printSlice(s)    // len=1 cap=1 [0]

	s = append(s, 1) // append a memeber to the new slice
	printSlice(s)    // len=2 cap=2 [0 1]

	s = append(s, 2, 3, 4) // append mutliple members once
	printSlice(s)          // len=5 cap=6 [0 1 2 3 4]

	s = append(s, 5) // try to see how compiler decide the cap
	printSlice(s)    // len=6 cap=6 [0 1 2 3 4 5]

	// when the underlying array can not cap the new members, a larger array will be created and a slice that underlys the array is returned
	s = append(s, 6, 7) // try to see how compiler decide the cap
	printSlice(s)       // len=8 cap=12 [0 1 2 3 4 5 6 7]
}

// use for ...range iteration
func slice_iteration() {
	// var alphabet = []string{'a', 'b', 'c', "d", "e"}  //  cannot use 'a' (untyped rune constant 97) as string value in array or slice literal
	// cannot use 'b' (untyped rune constant 98) as string value in array or slice literal
	var alphabet = []string{"a", "b", "c", "d", "e"}
	for index, value := range alphabet { // for ... range iterate slice or map
		fmt.Printf("%d: %s\n", index, value) // range returns the index and a copy of the value in slice
	}
}

func main() {
	slice_array()
	slice_quote()
	slice_init()
	slice_default_bound()
	slice_len_cap()
	zero_slice()
	make_slice()
	game_slice()
	slice_append()
	slice_iteration()
}
