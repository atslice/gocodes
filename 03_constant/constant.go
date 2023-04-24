package main

import (
	"fmt"
	"unsafe"
)

// use key word const to define constant
const Pi = 3.1415

// const identifier [type] = value
const Name string = "Alice"

// or ignore type
// const identifier = value
const Height = 100

// define constants in a line
const ip, port = "127.0.0.1", 5638

// define constants in group, for example, enum
const (
	Anoy        = 1
	High        = 0
	Transparent = 2
	Unknown     = 3
)

// import "unsafe"  // Build Error: imports must appear before other declarations

// constant can be declared by built-in function, such as len(), cap(), unsafe.Sizeof()

const (
	sentence = "This is a sentence."
	length   = len(sentence)           // 19
	size     = unsafe.Sizeof(sentence) // 16
)

/*

It is said that constant can not be declared by user's function, otherwise build-error.

func get_int() int {
	return 100
}

const haha = get_int()  // get_int() (value of type int) is not constant

func get_Height() {
	return
}

const hh = get_Height()  // get_Height() (no value) used as value
*/

func get_int() int {
	return 100
}

var bint = get_int()

const (
	a = 2
	b // in group declaration of constants, declared constant without value, the value of the last line will be used as the default value
	c
	d
)

const aa = 3

// can not miss expression in non-group declaration of constant
// const bb  // Build Error:  missing init expr for bb
// const cc  // Build Error:  missing init expr for cc

func main() {
	fmt.Println("Pi: ", Pi)
	fmt.Println("Name: ", Name)
	fmt.Println("Hight: ", Height)
	fmt.Println("ip: ", ip)
	fmt.Println("port: ", port)
	fmt.Println("sentence:", sentence, "length:", length, "size:", size)
	fmt.Println("bint:", bint)
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d)
	fmt.Println("aa:", aa)

	const s1 = 1 << 1
	fmt.Println("s1:", s1) // 2

	const s2 = 1 << 2
	fmt.Println("s2:", s2) // 4

	const s3 = 1 << 3
	fmt.Println("s3:", s3) // 8

	const s32 = 1 << 32
	fmt.Println("s32:", s32) // 4294967296

	const s33 = 1 << 33
	fmt.Println("s33:", s33) // 8589934592

	const s48 = 1 << 48
	fmt.Println("s48:", s48) // 281474976710656

	const s60 = 1 << 60
	fmt.Println("s60:", s60) // 1152921504606846976

	const s62 = 1 << 62
	fmt.Println("s62:", s62) // 4611686018427387904

	const s63 = 1 << 63
	//fmt.Println("s63:", s63) // cannot use s63 (untyped int constant 9223372036854775808) as int value in argument to fmt.Println (overflows)

	const s64 = 1 << 64
	//fmt.Println("s64:", s64) // cannot use s64 (untyped int constant 18446744073709551616) as int value in argument to fmt.Println (overflows)

	const big = 1 << 100
	//fmt.Println("big:", big) // cannot use big (untyped int constant 1267650600228229401496703205376) as int value in argument to fmt.Println (overflows)

	const s500 = 1 << 500
	// fmt.Println("s500:", s500) // cannot use s500 (untyped int constant 3273390607896141870013189696827599152216642046043064789483291368096133796404674554883270092325904157150886684127560071009217256545885393053328527589376) as int value in argument to fmt.Println (overflows)

	const s511 = 1 << 511
	//fmt.Println("s511:", s511)  // cannot use s511 (untyped int constant 6703903964971298549787012499102923063739682910296196688861780721860882015036773488400937149083451713845015929093243025426876941405973284973216824503042048) as int value in argument to fmt.Println (overflows) (exit status 1)

	//const s512 = 1 << 512  //constant shift overflow (exit status 1)
	//const s1000 = 1 << 1000  // constant shift overflow (exit status 1)

	// iota is special constant, once the keyword const presents, iota will be reset as 0, every line is added in const group declaration, iota will be added by 1 (still constant)
	const (
		ia = iota // 0
		ib = iota // 1
		ic = iota // 2
	)
	fmt.Println("ia:", ia, "ib:", ib, "ic:", ic)

	const (
		iaa = iota // 0
		ibb        // 1
		icc        // 2
	)
	fmt.Println("iaa:", iaa, "ibb:", ibb, "icc:", icc)

	const (
		i_aa = iota  // 0
		i_bb         // 1
		i_cc         // 2
		i_dd = "yes" // yes, iota == 3
		i_ee         // yes, iota == 4
		i_ff = 50    // 50,  iota == 5
		i_gg         // 50,  iota == 6
		i_hh = iota  // 7,  iota == 7
		i_ii         // 8,  iota == 8
	)
	fmt.Println("i_aa:", i_aa, "i_bb:", i_bb, "i_cc:", i_cc, "i_dd:", i_dd, "i_ee:", i_ee, "i_ff:", i_ff, "i_gg:", i_gg, "i_hh:", i_hh, "i_ii:", i_ii)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i:", i, "j:", j, "k:", k, "l", l)
}
