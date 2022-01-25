package main

import "fmt"

func main() {
	numChild := 5

	switch numChild {
	case 5:
		fmt.Println("Mantap")
	case 3:
		fmt.Println("Bagus")
	default:
		fmt.Println("Semangat")
	}
	// switch {
	// case numChild > 5:
	// 	fmt.Println("Lebih dari lima")
	// case numChild > 3:
	// 	fmt.Println("Lebih dari tiga")
	// default:
	// 	fmt.Println("Kurang dari tiga")
	// }
}
