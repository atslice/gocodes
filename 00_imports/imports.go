package main

import (
	"fmt"
	"math"
)

/*
if not using double quote:
import math
error: missing import path
*/

/*
	use () to group import
import "fmt"
import "math"
*/

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}