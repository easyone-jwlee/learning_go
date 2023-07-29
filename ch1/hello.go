package main

import (
	"fmt"
)

type Param struct {
	x int
}

func main() {
	// fmt.Println("Hello, World!")
	// p :=
	p := 200
	fmt.Printf("----------------------- p: %v\n", p)
	fmt.Printf("----------------------- &p: %p\n", &p)
	update(&p)
	fmt.Printf("----------------------- p: %v\n", p)
	fmt.Printf("----------------------- &p: %p\n", &p)
}

func update(ppp *int) {
	fmt.Printf("----------------------- ppp1: %p\n", ppp)
	c := 400
	fmt.Printf("----------------------- &c: %p\n", &c)
	*ppp = c
	fmt.Printf("----------------------- ppp2: %p\n", ppp)
}

// func update(ppp *Param) {
// 	p := Param{x: 400}
// 	ppp = &p
// }

// void update(int *p) {
//     *p = 100;
// }

// int main() {
//     int x = 50;
//     printf("Before update: %d\n", x);
//     update(&x);
//     printf("After update: %d\n", x);
//     return 0;
// }
