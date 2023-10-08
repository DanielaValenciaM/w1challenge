package main

import (
	"fmt"
	"os"
	"w1challenge/Bmi"
	"w1challenge/Evenodd"
	"w1challenge/Foobar"
	"w1challenge/Mario"
	"w1challenge/Ohm"
)

func main() {
	args := os.Args
	functionName := args[1]
	switch functionName {
	case "evenodd":
		var num int
		fmt.Println("Que número deseas evaluar?")
		fmt.Scan(&num)
		fmt.Println(Evenodd.Evenodd(num))
	case "ohm":
		var v float32
		var r float32
		var i float32
		fmt.Println("Cuales son los valores para v, r, i?")
		fmt.Scan(&v, &r, &i)
		fmt.Println(Ohm.Ohm(v, r, i))
	case "foobar":
		var limit int
		fmt.Println("Cual sera el limite de FooBar?")
		fmt.Scan(&limit)
		fmt.Println(Foobar.Foobar(limit))
	case "mario":
		var response int
		fmt.Print("What height would you like Mario's pyramid to be, within the range of 1 to 8?:")
		fmt.Scan(&response)
		fmt.Println(Mario.Pyramid(response))
	case "bmi":
		var weight float64
		var height float64
		fmt.Println("How much do you weigh? (don't lie)")
		fmt.Scanln(&weight)
		fmt.Println("How tall are you? (barefoot)")
		fmt.Scanln(&height)
		fmt.Println(Bmi.BmiCalculator(weight, height))
	}
}
