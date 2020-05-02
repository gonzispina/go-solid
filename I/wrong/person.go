package wrong

type Person interface {
	Name() string
	Eat()
	Sleep()
	Work()
}

type PersonBase struct {
	name string
	hungry bool
	sleepy bool
}

func (p *PersonBase) Name() string { return p.name }
