package right

type Retired struct {
	*PersonBase
}

func (h *Retired) Eat() {
	h.hungry = false
}

func (h *Retired) Sleep() {
	h.hungry = true
}



