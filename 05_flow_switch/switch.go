package main

import (
	"fmt"
	"time"
)

// switch example
func check_media(x string) string {
	media_type := "unknown"
	switch x {
	case "image/jpg": // put a : following
		media_type = "image" // no need to break, once one case is excuted, go auto breaks
	case "image/png": // variable of x can by any type, but case statement or the value of the expression of case statement, should be the same type
		media_type = "image"
	case "video/mp4": //
		media_type = "video"
	case "audio/mp3":
		media_type = "audio"
	default:
		media_type = "not support"
	}
	return media_type
}

// fallthrough
func check_fall(x string) string {
	media_type := "unknown"
	switch x {
	case "image/jpg": // put a : following
		media_type = "image" // no need to break, once one case is excuted, go auto breaks
	case "image/png": // variable of x can by any type, but case statement or the value of the expression of case statement, should be the same type
		media_type = "image"
		fallthrough // if fallthrough, next case block will be executed without calculating the case statement
	case "video/mp4": // becareful if using fallthrough
		media_type = "video"
	case "audio/mp3":
		media_type = "audio"
	default:
		media_type = "not support"
	}
	return media_type
}

// multiple match
func menu(x int) {
	switch x {
	case 0, 1, 2:
		fmt.Println("low")
	case 3, 4, 5:
		fmt.Println("mid")
	case 6, 7, 8:
		fmt.Println("high")
	case 9:
		fmt.Println("good")
	case 10:
		fmt.Println("perfect")
	default:
		fmt.Println("^-^")
	}
}

// switch without condition
func switch_condition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {

	fmt.Println(check_media("amazon"))
	fmt.Println(check_media("image/png"))
	fmt.Println(check_media("audio/mp3"))

	fmt.Println(check_fall("amazon"))

	fmt.Println(check_fall("image/png")) // match 'case "image/png":', the block is executed(media_type is given value "image") and then fallthrough to next case '"case "video/mp4":"'
	// though not matched, the block is executed(media_type is given value "video"), then break for no "fallthrough"
	fmt.Println(check_fall("audio/mp3"))
	menu(0)
	menu(2)
	menu(200)
	menu(4)
	menu(8)
	menu(9)
	menu(10)
	// menu("sfe")  // cannot use "sfe" (untyped string constant) as int value in argument to menu (exit status 1)
	switch_condition()

}
