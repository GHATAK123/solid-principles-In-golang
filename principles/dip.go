package principles

import "fmt"

// DIP - Dependency Inversion Principle
// - High-level modules should not depend on low-level modules. Both should depend on abstractions.
// - Abstractions should not depend on details. Details should depend on abstractions.

type Faculty interface {
	Work()
}

type Teacher struct{}

func (t Teacher) Work() {
	fmt.Println("Teaching Students!!!")
}

type Assistant struct{}

func (a Assistant) Work() {
	fmt.Println("Assisting Teachers!!!")
}

type Helper struct{}

func (h Helper) Work() {
	fmt.Println("Helping Casually!!!")
}

type Secretary struct{}

func (s Secretary) Work() {
	fmt.Println("Managing Office!!!")
}

type HeadMaster struct {
	faculty []Faculty
}

func (hm HeadMaster) Manage() {
	for _, f := range hm.faculty {
		f.Work()
	}
}

func DIP() {
	fmt.Println("From DIP Package!!!")
	teacher := Teacher{}
	assistant := Assistant{}
	helper := Helper{}
	secretary := Secretary{}
	Faculty := []Faculty{teacher, assistant, helper, secretary}
	headMaster := HeadMaster{faculty: Faculty}
	headMaster.Manage()
}


// ✔ Headmaster depends on the Faculty interface, not specific faculty types
// ✔ Adding new faculty members (like Secretary) requires no changes to Headmaster
// ✔ System is now loosely coupled, flexible, and easily extendable
// Before DIP: Headmaster directly depended on concrete faculty types → rigid and hard to modify.
// After DIP: Headmaster now depends on the Faculty interface → flexible and extensible.

