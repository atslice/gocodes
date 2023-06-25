package main

import "fmt"

func make_map() {
	fmt.Println("make_map():")
	// map is un-sorted collections of key-value pairs
	// To create an empty map, use the builtin function make
	// make(map[key-type]value-type)
	m := make(map[string]int)
	// set key-value pairs, map_name[key] = value
	m["length"] = 50
	m["width"] = 100

	// Printing a map, will show all of its key-value pairs
	fmt.Println("map:", m) // map: map[length:50 width:100]
}

func make_map2() {
	fmt.Println("make_map2():")
	// map is un-sorted collections of key-value pairs
	// To create an empty map, use the builtin function make
	// make(map[key-type]value-type, initialCapacity)
	// the argument initialCapacity is optional
	m := make(map[int]string, 10)

	// The builtin function len returns the number of key-value pairs in a map
	fmt.Println("length of map:", len(m)) // legnth of map: 0

	// set key-value pairs, map_name[key] = value
	m[10] = "male"
	m[20] = "female"

	// Printing a map, will show all of its key-value pairs
	fmt.Println("map:", m) // map: map[10:male 20:female]

	// The builtin function len returns the number of key-value pairs in a map
	fmt.Println("length of map:", len(m)) // legnth of map: 2
}

func make_map3() {
	fmt.Println("make_map3():")
	// make map with initialization
	m := map[string]bool{
		"Apple":  true,
		"Google": false,
		"Meta":   false,
	} // Do Not forget the last , otherwise syntax error: unexpected newline in composite literal

	// Printing a map, will show all of its key-value pairs
	fmt.Println("map:", m) // map: map[Apple:true Google:false Meta:false]

	n := map[string]float64{"pi": 3.1415926, "rate": 0.618} // initialize in a single line, the last , can be ignored

	// Printing a map, will show all of its key-value pairs
	fmt.Println("map:", n) // map: map[pi:3.1415926 rate:0.618]
}

func main() {
	make_map()
	make_map2()
	make_map3()
}
