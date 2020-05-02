package wrong

type Retired struct {
	*PersonBase
}

func (h *Retired) Eat() {
	h.hungry = false
}

func (h *Retired) Sleep() {
	h.hungry = true
}

// This method is polluting our jubilado implementation
func (h *Retired) Work() {
	panic("Retired people shouldn't work unless they are argentinian")
}



