package internal

import "fmt"

func Switch() {
	i := 2

	fmt.Println("SWITCH CASE TO: ", i)

	switch i {
	case 1:
		fmt.Println("UM")
	case 2:
		fmt.Println("DOIS")
	case 3:
		fmt.Println("TRES")
	}

	t := "TESTE"

	switch t {
	case "UM":
		fmt.Println("UM")
	case "DOIS":
		fmt.Println("DOIS")
	case "TRES":
		fmt.Println("TRES")
	default:
		fmt.Println("ANOTHER VALUE: ", t)
	}

}
