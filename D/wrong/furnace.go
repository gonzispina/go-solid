package wrong

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Disengage byte = 0x00
	Engage    byte = 0xFF
)

func Regulate(maxTemp int) {
	thermometerReader := bufio.NewReader(os.Stdin)
	furnaceWriter := bufio.NewWriter(os.Stdout)

	writeErr := func(err error) {
		_, _ = fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
	}

	for {
		byteTemp, err := thermometerReader.ReadByte()
		if err != nil{
			writeErr(err)
		}

		temp := int(byteTemp)
		var action byte
		if temp < maxTemp {
			action = Engage
		} else {
			action = Disengage
		}

		_ = furnaceWriter.WriteByte(action)
	}
}