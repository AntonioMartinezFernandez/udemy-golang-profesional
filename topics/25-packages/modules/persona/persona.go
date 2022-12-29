package persona

type Persona struct {
	name string
	age  int
}

func (p *Persona) Constructor(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Persona) SetName(name string) {
	p.name = name
}

func (p *Persona) SetAge(age int) {
	p.age = age
}

func (p *Persona) GetName() string {
	return p.name
}

func (p *Persona) GetAge() int {
	return p.age
}
