package main

import "fmt"

func add(x, y int) int {
	return x + y // add: +
}

func minus(x, y float64) float64 {
	return x - y // minus: -
}

// more than: >
func greater(x, y float64) bool {
	return x > y
}

// less than: <
func smaller(x, y float64) bool {
	return x < y
}

// equal: =
func equal(x, y float64) bool {
	return x == y
}

// non-equal: !=
func nonequal(x, y float64) bool {
	return x != y
}

// more than or equal to: >=
func greater_equal(x, y float64) bool {
	return x >= y
}

// less than or equal to: <=
func smaller_equal(x, y float64) bool {
	return x <= y
}

// logic operation
// logic and: &&
func and(x, y bool) bool {
	return x && y
}

// logic or: ||
func or(x, y bool) bool {
	return x || y
}

// logic not: !
func not(x bool) bool {
	return !x
}

func main() {
	fmt.Println(add(5, 20))
	fmt.Println(minus(598.6, 45.26))
	fmt.Println("9 is greater than 6:", greater(9, 6))
	fmt.Println("9 is greater than 9.0:", greater(9, 9.0))
	fmt.Println("9 is smaller than 9.0:", smaller(9, 9.0))
	fmt.Println("9 is equal to 9.0:", equal(9, 9.0))
	fmt.Println("9 is not equal to 9.0:", nonequal(9, 9.0))
	fmt.Println("9 is greater than or equal to 9.0:", greater_equal(9, 9.0))
	fmt.Println("9 is less than or equal to 9.0:", smaller_equal(9, 9.0))
	// fmt.Println("0 && 1:", and(0, 1))  // Build Error: cannot use 0 (untyped int constant) as bool value in argument to and; cannot use 1 (untyped int constant) as bool value in argument to and
	fmt.Println("true && true:", and(5 > 2, 3 > 1))
	// i := 0
	fmt.Println("false && true:", and(5 < 2, true))
	fmt.Println("false || true:", or(5 < 2, true))
	fmt.Println("true || true:", or(5 >= 2, true))
	fmt.Println("not true:", not(true))

	// bit operator
	fmt.Println("0 & 0:", 0&0)
	fmt.Println("0 & 1:", 0&1)
	fmt.Println("1 & 0:", 1&0)
	fmt.Println("1 & 1:", 1&1)
	fmt.Println("0 | 0:", 0|0)
	fmt.Println("0 | 1:", 0|1)
	fmt.Println("1 | 0:", 1|0)
	fmt.Println("1 | 1:", 1|1)
	fmt.Println("0 ^ 0:", 0^0)
	fmt.Println("0 ^ 1:", 0^1)
	fmt.Println("1 ^ 0:", 1^0)
	fmt.Println("1 ^ 1:", 1^1)
	var sa = 5
	var sb = 30
	fmt.Println("sa & sb:", sa&sb)
	fmt.Println("sa | sb:", sa|sb)
	fmt.Println("sa ^ sb:", sa^sb)
	fmt.Println("1<<0:", 1<<0)
	fmt.Println("1<<1:", 1<<1)
	fmt.Println("1<<2:", 1<<2)
	fmt.Println("1<<3:", 1<<3)
	fmt.Println("1<<4:", 1<<4)
	fmt.Println("1<<5:", 1<<5)
	fmt.Println("0<<0:", 0<<0)
	fmt.Println("0<<1:", 0<<1)
	fmt.Println("0<<2:", 0<<2)
	fmt.Println("0<<3:", 0<<3)
	fmt.Println("32>>0:", 32>>0)
	fmt.Println("32>>1:", 32>>1)
	fmt.Println("32>>2:", 32>>2)
	fmt.Println("32>>3:", 32>>3)
	fmt.Println("32>>4:", 32>>4)
	fmt.Println("32>>5:", 32>>5)
	fmt.Println("32>>6:", 32>>6)

	// =
	// +=
	// -=
	// *=
	// /=
	// %=
	// <<=
	// >>=
	// &=
	// ^=
	// |=
	// & to get address
	// * pointer
	var ai = 68
	var ptr *int
	fmt.Println("value of ai:", ai)
	fmt.Println("memeory address &ai:", &ai)
	ptr = &ai
	fmt.Println("ptr = &ai")
	fmt.Println("*ptr:", *ptr)
	fmt.Println("ptr:", ptr)
}
