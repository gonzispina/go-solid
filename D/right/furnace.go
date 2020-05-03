package right

type Thermometer interface {
	Read() (int, error)
}

type Heater interface {
	Engage()    error
	Disengage() error
}

type ErrorHandler interface {
	SetError(err error)
}

func Regulate(maxTemp int, thermometer Thermometer, heater Heater, errHandler ErrorHandler) {
	for {
		temp, err := thermometer.Read()
		if err != nil {
			errHandler.SetError(err)
			continue
		}

		if temp < maxTemp {
			err = heater.Engage()
		} else {
			err = heater.Disengage()
		}

		if err != nil {
			errHandler.SetError(err)
		}
	}
}

