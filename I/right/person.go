package right

type Person interface {
	Name() string
	Eat()
	Sleep()
}

type SystemSlave interface {
	Work()
}

type PersonBase struct {
	name string
	hungry bool
	sleepy bool
}

func (p *PersonBase) Name() string { return p.name }
