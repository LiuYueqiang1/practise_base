package main

import "fmt"

func main() {
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	println(USD, symbol[USD])

}
