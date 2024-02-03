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

}

type example struct {
	flag    bool
	counter int16
	pi      float32
}
