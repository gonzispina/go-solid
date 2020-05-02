package wrong

type McDonaldsEmployee struct {
	*PersonBase
}


func (m *McDonaldsEmployee) Eat() {
	m.hungry = false
}

func (m *McDonaldsEmployee) Sleep() {
	m.sleepy = false
}

func (m *McDonaldsEmployee) Work() {
	m.sleepy = true
	m.hungry = true
}
