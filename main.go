package main

import (
	"fmt"
	"solidPrinciple/principles"
)

func main() {
	fmt.Println("Solid Principle Demo In Go!")
	principles.SRP()
	fmt.Println("=====================================")
	principles.OCP()
	fmt.Println("=====================================")
	principles.LSP()
	fmt.Println("=====================================")
	principles.ISP()
	fmt.Println("=====================================")
	principles.DIP()
}
