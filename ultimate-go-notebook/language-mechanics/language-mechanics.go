package main

import "fmt"

func main() {

	//Declare and Initialize
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var a string \t %T [%v]\n", b, b)
	fmt.Printf("var a float64 \t %T [%v]\n", c, c)
	fmt.Printf("var a bool \t %T [%v]\n", d, d)

	a = 10
	b = "hello"
	c = 3.14159
	d = true

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var a string \t %T [%v]\n", b, b)
	fmt.Printf("var a float64 \t %T [%v]\n", c, c)
	fmt.Printf("var a bool \t %T [%v]\n", d, d)

	// Conversion vs Casting

	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)

	//Struct and Construction Mechanics

	var e1 example
	fmt.Printf("%+v\n", e1)

	e2 := example{flag: false,
		counter: 10,
		pi:      3.14}

	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    false,
		counter: 20,
		pi:      3.14,
	}

	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)

	//Padding and Alignment

	//If I need to minimize the amount of padding bytes, I must lay out the fields from
	//highest allocation to lowest allocation. This will push any necessary padding bytes
	//down to the bottom of the struct and reduce the total number of padding bytes
	//necessary.

	//type example struct {
	//	pi float32
	//} // 0xc000100020 <- Starting Address counter	//// int16 // 0xc000100024 <- 2 byte alignment
	////flag bool // 0xc000100026 <- 1
	////byte alignment
	////flag2 bool // 0xc000100027 <- 1 byte alignment
	////}

	//Assigning Values

	//var ex1 example
	//var ex2 example1

	//	ex1 = ex2  Not allowed, compiler error

	//ex1 = example(ex2) Allowed, NO compiler error

	//var ex2 struct {
	//	flag    bool
	//	counter int16
	//	pi      float32
	//}

	// ex1=ex2 Allowed, NO need for conversion syntax

	//Pointers
	//Pass By Value

	count := 10
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	increment1(count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	increment2(&count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	//Escape Analysis

}

func increment1(inc int) {
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

func increment2(inc *int) {
	*inc++
	println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tPoints To[", *inc, "]")
}

type example struct {
	flag    bool
	counter int16
	pi      float32
}

type example1 struct {
	flag    bool
	counter int16
	pi      float32
}
