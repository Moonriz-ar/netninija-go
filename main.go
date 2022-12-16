package main

import "fmt"

func main() {
	myBill := newBill("mario's bill", map[string]float64{"pie": 5.99, "cake": 3.99})

	fmt.Println(myBill.format())
}