package struct_test

import "testing"

type Employee struct {
	ID     int
	Name   string
	Salary int
}

func TestStruct(t *testing.T) {
	t.Log("TestStruct")
	e := Employee{1, "Bob", 1000}
	t.Log(e)
	e.Salary = 2000
	t.Log(e)
	e1 := &e
	e1.Salary = 3000
	t.Log(e1, e)
}

type Code string
type Programmer interface {
	WriteHelloPpm() Code
}
type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloPpm() Code {
	return "fmt.Println(\"Hello, World!\")"
}
func TestInterFace(t *testing.T) {
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloPpm())
}
