package main

import "fmt"

func main() {

	var jawaban int

	fmt.Println("Permainan Ganjil Genap")
	_, err := fmt.Scan(&jawaban)

	if err != nil {
		fmt.Println("input tidak valid")
		return
	}

	if jawaban%2 == 0 {
		fmt.Println("genap")
	} else {
		fmt.Println("ganjil")
	}

}
