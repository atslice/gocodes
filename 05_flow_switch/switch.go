package main

import "fmt"

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
	case "video/mp4": //
		media_type = "video"
	case "audio/mp3":
		media_type = "audio"
	default:
		media_type = "not support"
	}
	return media_type
}

func main() {

	fmt.Println(check_media("amazon"))
	fmt.Println(check_media("image/png"))
	fmt.Println(check_media("audio/mp3"))

}
