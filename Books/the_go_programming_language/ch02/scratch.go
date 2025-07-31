package main

var f, err = os.Open(name)

func main() {

	x := 1
	p := &x				// p, of type *int, points to x
	fmt.Println(*p)		// "1"
	*p = 2				// equivalent to x = 2
	fmt.Println(x)		// "2"

	//Each component of a variable of aggregate type&mdash;a field of a struct or an element of an array&mdash;is also a variable and thus has an address too.
	// For example:
	type Point struct {
		X int
		Y int
	}
	p := &Point{1, 2}	// p, of type *Point, points to a Point struct
	fmt.Println(p.X)	// "1"
	p.X = 3			// equivalent to (*p).X = 3
	fmt.Println((*p).X) // "3"

}
