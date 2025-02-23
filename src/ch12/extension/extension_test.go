package extension_test

import "testing"

type Pet struct {
}

func (p *Pet) Speak() {
	print("Pet Speak")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	print(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	print("Dog Speak")
	// d.p.Speak()
}
func (d *Dog) SpeakTo(host string) {
	// d.p.SpeakTo(host)
	d.Speak()
	print(" ", host)
}

func TestExtension(t *testing.T) {
	d := new(Dog)
	d.SpeakTo("pp")
}
