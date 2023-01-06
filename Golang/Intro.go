package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Hello world")

	s := "Go tutorial"
	fmt.Println(s)
	i := 123
	fmt.Println(i)
	fmt.Printf("value of i = %d value of s = %s\n", i, s)

	arr := [4]int8{1, 2, 3, 4}
	fmt.Println(arr, len(arr), cap(arr))

	slice := []int8{1, 2, 3, 4}
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 5)
	fmt.Println(slice, len(slice), cap(slice))

	d1 := map[int](string){1: "A", 2: "B"}
	fmt.Println(d1)
	fmt.Printf("Map %v\n", d1)

	d2 := make(map[string]string)
	d2["A"] = "Apple"
	d2["B"] = "Big Apple"
	d2["C"] = "Cute Apple"
	fmt.Printf("Map %v\n", d2)

	for i, a := range slice {
		fmt.Println(i, a)
	}
	for k, v := range d2 {
		fmt.Println(k, v)
	}

	t1 := tyre{
		radius: 1.0, width: .8,
	}
	t2 := tyre{
		radius: 1.0, width: .8,
	}
	t3 := tyre{
		radius: 1.0, width: .8,
	}
	t4 := tyre{
		radius: 1.0, width: .8,
	}
	c := car{
		tyre: [4]tyre{t1, t2, t3, t4},
	}
	fmt.Println(c)
	fmt.Printf("%s %v \n ", c, c)
	fmt.Println(reflect.TypeOf(c))

	go print("Goroutine")
	print("Normal")
	time.Sleep(time.Second)
}

func print(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

type tyre struct {
	radius float32
	width  float32
}
type car struct {
	tyre [4]tyre
}

func (c car) String() string {
	return fmt.Sprintf("Car has %d tyre", len(c.tyre))
}
