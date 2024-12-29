package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var a float64
	var b string
	var c float64

	fmt.Println("Iveskite pirmą skaičių")
	if scanner.Scan() {
		a, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	for {
		fmt.Println("Pasirinkite ženklą: + - * / stop")
		if scanner.Scan() {
			b = scanner.Text()
			if b == "stop" {
				break
			}
			fmt.Println("Iveskite papildomą skaičių")
			if scanner.Scan() {
				c, _ = strconv.ParseFloat(scanner.Text(), 64)
			}
			if b == "+" {
				a = a + c
				fmt.Println("Atsakymas", a)
			} else if b == "-" {
				a = a - c
				fmt.Println("Atsakymas", a)
			} else if b == "*" {
				a = a * c
				fmt.Println("Atsakymas", a)
			} else if b == "/" {
				a = a / c
				fmt.Println("Atsakymas", a)
			} else {
				fmt.Println("Klaidingas ženklas")
			}
		}
	}
}
