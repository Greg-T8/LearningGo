package main

var f, err = os.Open(name)

func incr(p *int) {
	*p++				// increments what p points to; does not change p
	return *p
}

v := 1
incr(&v)				// side effect:  v is now 2
fmt.Println(incr(&v))	// "3" (and v is 3

func main() {

}