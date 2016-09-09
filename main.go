package main

func main() {

	// a := 0
	// a := 1
	a := 2
	// a := 3

	switch a {
	case 0:
		println(`I'm zero`)

	case 1:
		println(`I'm one`)
		fallthrough

	case 2:
		println(`I'm two`)
		fallthrough

	case 3:
		println(`I'm three`)

	default:
		println(`I know who i am!`)
	}

}
