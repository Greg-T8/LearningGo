package main

func main() {

func newInt() *int {
	return new(int)
}

func newInt() *int {
	var dummy int
	return &dummy
}

p := new(int)
q := new(int)
fmt.Println(p == q)		// false, p and q point to different memory addresses

func delta(old, new int) int { return new - old }

}
