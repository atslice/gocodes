package main

import "fmt"

// a struct is a group of field
// define a struct
type Vertex struct {
	X int
	Y int
}

// memeber of struct can be of different type
// the identifier of struct can be started with lower-case
type person struct {
	name string
	age  int
}

// access member of struct by dot
func struct_memeber() {
	fmt.Println("struct_memeber()")
	p := person{"Belinda", 18} // declare a struct of person with identifier p
	fmt.Println(p.name, p.age) // access member of struct by dot
	fmt.Println(p)             // output: {Belinda 18}
}

// pointer to struct
func struct_pointer() {
	fmt.Println("struct_pointer()")
	type credits struct {
		name    string
		math    float64
		english float64
		phy     float64
	}
	// cindy := credits{"Belinda", 96.0}      // Build Error: too few values in struct literal of type credits (exit status 1)
	cindy := credits{"Belinda", 96.0, 100, 59.8}
	fmt.Println(cindy)                     // output: {Belinda 96 100 59.8}
	fmt.Println(cindy.name, cindy.english) // access member of struct by dot
	p := &cindy                            // pointer to struct
	fmt.Println((*p).name, (*p).english)   // *p is the value of the variable that the pointer points to, so of course (*p).name can be used to access the member of the struct cindy
	fmt.Println(p.name, p.english)         // go supports simplifing grammar, use p.name instead of (*p).name to access the member of a struct
}

// struct init
func struct_init() {
	fmt.Println("struct_init()")
	// is it possible to initialize when defining a struct
	type credits struct {
		name    string
		math    float64
		english float64
		phy     float64
	}
}

func main() {
	// fmt.Println(person{1, 2})  // cannot use 1 (untyped int constant) as string value in struct literal (exit status 1)
	fmt.Println(person{"Alice", 16}) // output: {Alice 16}
	// fmt.Println(person.name)  // Build Error: person.name undefined (type person has no method name) (exit status 1)
	fmt.Println(Vertex{1, 2})
	struct_memeber()
	struct_pointer()
}
