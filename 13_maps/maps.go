package main

import (
	"fmt"
	"strings"
)

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

func get_map_value() {
	fmt.Println("get_map_value():")

	m := map[string]int{
		"length": 50,
		"width":  60,
	}

	// get the value by the key
	length := m["length"]
	fmt.Println("The length in map:", length)

	// If a key does not exist in a map, the zero value of the value type is returned
	nokey := m["height"]
	fmt.Println("The value of nokey:", nokey) // The value of nokey: 0

	// The optional second return value when getting a value from a map
	// 		indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	_, present := m["height"]
	fmt.Println("The key height is present:", present) // The key height is present: false
}

func delete_map_key() {
	fmt.Println("delete_map_key():")

	// use the builtin function delete() to delete a key-value pair in a map
	m := map[string]int{
		"length": 50,
		"width":  60,
		"height": 70,
	}

	fmt.Println(m) // map[height:70 length:50 width:60]
	delete(m, "height")
	fmt.Println(m)         // map[length:50 width:60]
	delete(m, "nosuchkey") // not any error when delete non-exist key
	fmt.Println(m)         // map[length:50 width:60]
}

func modify_map_value() {
	fmt.Println("modify_map_value():")

	m := map[string]int{
		"length": 50,
		"width":  60,
		"height": 70,
	}

	fmt.Println(m)    // map[height:70 length:50 width:60]
	m["height"] = 140 // modify the value
	fmt.Println(m)    // map[height:140 length:50 width:60]
	m["timing"] = 90  // add new key-value pair
	fmt.Println(m)    // map[height:140 length:50 timing:90 width:60]

}

func modify_key_value(m map[string]int, key string, value int) {
	m[key] = value
}

func map_is_pointer() {
	fmt.Println("map_is_pointer():")

	m := map[string]int{
		"length": 50,
		"width":  60,
		"height": 70,
	}

	// if the underlying is modified, any refer to the map sees the change
	fmt.Println(m)                    // map[height:70 length:50 width:60]
	modify_key_value(m, "length", 20) // map is passed as argument, but if it is modified, the underlying is modified indeed
	fmt.Println(m)                    // map[height:70 length:20 width:60]
}

func iter_map() {
	fmt.Println("iter_map():")

	m := map[string]int{
		"length": 50,
		"width":  60,
		"height": 70,
	}

	// use range map_name to iter key-value pairs in a map
	for key, value := range m {
		fmt.Printf("key = %s, value = %d\n", key, value)
	}
}

func zero_value_of_map() {
	fmt.Println("zero_value_of_map():")
	// map is un-sorted collections of key-value pairs
	// To create an empty map, use the builtin function make
	// make(map[key-type]value-type)
	m := make(map[string]int)
	fmt.Println(m) // map[]
}

func WordCount(s string) map[string]int {
	fmt.Println("WordCounts():")
	fields := strings.Fields(s)
	fmt.Println(fields) // [This is a good day.]  why?
	m := make(map[string]int)
	// for word := range fields {
	//		m[word] += 1
	// }
	return m
}

func TestWordCount() {
	s := "  This is a good day. "
	m := WordCount(s)
	fmt.Println(m)
}

func main() {
	make_map()
	make_map2()
	make_map3()
	get_map_value()
	delete_map_key()
	modify_map_value()
	map_is_pointer()
	iter_map()
	zero_value_of_map()
	TestWordCount()
}
