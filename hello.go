package main /* executable package */

import(
	"fmt"
	"math"
)
var (
	a int
	b bool
)

type book struct {
	title string
	author string
	subject string
	book_id string
}

func swap(num1, num2 int) (int, int) {
	return num2, num1
}
func main() {
	/* numbers are 0 by default, bools are false, strings are "" */
	var v = 10
	v, k := 90, "abc" /* If k is not a new variable, a compilation error will occur */ 
	const t string = "const"
	const(
		m = iota
		n
		p = 3 << iota
	)
	if p < 10 {
		fmt.Println("p < 10")
	} else {
		fmt.Println("p > 10")
	}
	switch{
		case v >= 90 :
			fmt.Println("A")
			fallthrough
		case v >= 80 && v < 90 :
			fmt.Println("B")
		case v >= 70 && v < 80 :
			fmt.Println("C")
			fallthrough
		case v >= 60 && v < 70 :
			fmt.Println("D")
		default:
			fmt.Println("Failed")
	}
	var x interface{}

	switch i:= x.(type) {
		case nil:
			fmt.Printf("x is %T\n", i)
		case int:
			fmt.Println("x is int")
		case func(int) float64:
			fmt.Println("x is func(int)")
		default:
			fmt.Println("Unknown type")
	}
	fmt.Printf("%v %v %v %v %v %v %v %v\n", a, b, v, k, t, m, n, p)
	/* Objects start with Uppercase could be used outside the package */
	LOOP: for a < 6 {
		a++;
		if a == 3 {
			goto LOOP
		}
		fmt.Printf("a = %d\n", a)
	}
	a, v = swap(a, v)
	fmt.Printf("a, v = %d, %d\n", a, v)

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
	var balance = [...]float32{100.0, 2.0, 3.4, 7.0, 5.0}
	balance[2] = 30.0
	var ip *int /* pointer */
	ip = &a
	fmt.Printf("ip = %d\n", *ip)
	var book1 book
	book1.title = "Go language"
	var s []int /* slice */
	s = [] int{1,2,3}
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s));
	s = append(s, 5,6,7)
	s1 := make([]int, len(s), (cap(s)) * 2)
	copy(s1, s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s));
	for _, num := range s {
		fmt.Printf("%d, ", num);
	}
	fmt.Println();
	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Printf("%d -> %s\n", k, v)
	}
	delete(kvs, 1)
}
