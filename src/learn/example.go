package main

import (
	"fmt"
	"strconv"
)

//  variable
var word string
var fname, lname = "ashari", "juang"

func main() {
	word = "Hello "
	// shorthand to create variable inside function
	fullname := fname + " " + lname
	age := 30

	fmt.Println(word + fullname)
	fmt.Println("your age " + strconv.Itoa(age))

	fmt.Println("1 + 1 =", tambah(1, 1))

	hasil, message := tambahWitMessage(1, 2)
	fmt.Println(message + strconv.Itoa(hasil))

	// for condition
	i := 0
	for i < 5 {
		println("maju", i)
		i++
	}

	//  atau bisa begini
	for a := 1; a < 5; a++ {
		println("serang", a)
	}
}

// sample function, return integer with 2 parameter integer
func tambah(angka1 int, angka2 int) int {
	return angka1 + angka2
}

func tambahWitMessage(angka1 int, angka2 int) (hasil int, message string) {
	// kita bisa return sesuai dengan nama yg sudah di tentukan sebelumnya tanpa harus define variable
	message = "hasil penjumlahan "
	hasil = angka1 + angka2

	return
}
