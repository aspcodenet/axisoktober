package main

import (
	"fmt"
)

func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)
	//area = l * b
	return
}

func getLongest(namn1 string, namn2 string) (string, error) {
	len1 := len(namn1)
	len2 := len(namn2)

	if len1 > len2 {
		return namn1, nil
	}
	return namn2, nil
}

func change(allaNamn [3]string) {
	allaNamn[1] = "Kerstin"
	for _, item := range allaNamn {
		fmt.Println(item)
	}
}

func changeSlice(slice []string) []string {
	slice[1] = "Kerstin"
	slice = append(slice, "Nisse")
	for _, item := range slice {
		fmt.Println(item)
	}
	return slice
}

func main() {

	// EFTER RASTEN
	//1. Flera filer???
	//2. Flera folders
	//3. User defined types - STRUCT
	//4. OOP "light"

	// Imorgon forts flera filer+folders och SCOPE
	//  dvs Access Modifiers

	allaNamnSlice := []string{"Stefan", "Josefine", "Oliver"}

	allaNamnSlice = append(allaNamnSlice, "Stefan", "31231", "234243423")

	for _, item := range allaNamnSlice {
		fmt.Println(item)
	}
	allaNamnSlice = changeSlice(allaNamnSlice)
	for _, item := range allaNamnSlice {
		fmt.Println(item)

	}

	allaNamnSlice = append(allaNamnSlice, "Kalle")
	//theLength := len(allaNamnSlice)

	allaNamn := [...]string{"Stefan", "Josefine", "Oliver"}

	for _, item := range allaNamn {
		fmt.Println(item)
	}
	change(allaNamn)
	for _, item := range allaNamn {
		fmt.Println(item)
	}

	// allaDagar := [...]string{"Måndag", "Tidasg"}
	// printAll(allaDagar)

	a := rectangle(12, 32)
	fmt.Print(a)
	namnen := []string{"hello", "world"}
	for _, s := range namnen {
		fmt.Println(s)
	}

	var name string
	var age int
	fmt.Print("vad heter du?:")
	fmt.Scanln(&name)
	fmt.Print("VHur gammal är du:")
	_, err := fmt.Scanln(&age) // trettio
	if err != nil {
		fmt.Println("jamen skriv in ett tal tack")
	}

	//var ii int = 12
	//var i int = 12
	// var i int = 12

	// i := 12
	// // i = 12
	// name := "Stefan"

	// name = strings.ToLower(name)

	// if i == 50 && name == "Stefan" {
	// 	fmt.Println("Du är ung")
	// }

	// fmt.Printf("Hej %d\n", i)
}
