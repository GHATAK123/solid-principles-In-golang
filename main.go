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

// ✔ Headmaster depends on the Faculty interface, not specific faculty types
// ✔ Adding new faculty members (like Secretary) requires no changes to Headmaster
// ✔ System is now loosely coupled, flexible, and easily extendable
// Before DIP: Headmaster directly depended on concrete faculty types → rigid and hard to modify.
// After DIP: Headmaster now depends on the Faculty interface → flexible and extensible.
