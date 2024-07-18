package internal

import "fmt"

func Maps() {
	mapString := map[string]string{"Chave": "Valor", "ChaveDois": "ValorDois"} // map de strings
	mapInt := map[string]int{"ChaveInt": 20, "ChaveIntDois": 30}               // map de inteiros

	// print do map completo
	fmt.Println("MapString: ", mapString)
	fmt.Println("MapInt: ", mapInt)

	// print do map em determinada posição
	fmt.Println("MapString na Posição Chave: ", mapString["Chave"])
	fmt.Println("MapInt na Posição ChaveInt: ", mapInt["ChaveInt"])

	// alterando valores
	mapString["Chave"] = "Outro Valor"
	mapInt["ChaveInt"] = 60

	// print do map em determinada posição
	fmt.Println("MapString na Posição Chave: ", mapString["Chave"])
	fmt.Println("MapInt na Posição ChaveInt: ", mapInt["ChaveInt"])

	// removendo valores
	delete(mapString, "ChaveDois")
	delete(mapInt, "ChaveIntDois")

	// print do map completo
	fmt.Println("MapString: ", mapString)
	fmt.Println("MapInt: ", mapInt)

	// iterando map
	for key, value := range mapString {
		fmt.Println("Chave: ", key, "Valor: ", value)
	}

	for key, value := range mapInt {
		fmt.Println("Chave: ", key, "Valor: ", value)
	}
}
