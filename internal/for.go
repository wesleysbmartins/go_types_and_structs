package internal

import "fmt"

func For() {
	l := []string{}

	// for tradicional
	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}

	l = append(l, "x", "y")

	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}

	// for range
	for i, value := range l {
		fmt.Println("Index: ", i, "Value: ", value)
	}

	// for como um while
	n := 0

	// enquanto for difente de 10
	for n != 10 {
		fmt.Println("LOOP: ", n)
		n = 10
	}

	// loop só vai parar no break
	for {
		fmt.Println("LOOP: ", n)
		if n > 200 {
			break
		}
		n = n * n
	}

}
