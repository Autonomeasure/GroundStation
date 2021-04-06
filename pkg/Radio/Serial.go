package Radio

import (
	"github.com/jacobsa/go-serial/serial"
	"github.com/warthog618/gpio"
	"io"
	"time"
)

func OpenSerial(port string, baudrate uint) (io.ReadWriteCloser, error) {
	options := serial.OpenOptions{
		PortName: port,
		BaudRate: baudrate,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
	}

	serialPort, err := serial.Open(options)
	// if err != nil {
	// 	log.Fatalf("serial.Open: %v", err)
	// }

	return serialPort, err

	// defer serialPort.Close()
	// reader := bufio.NewReader(serialPort)
	// scanner := bufio.NewScanner(reader)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
}

func ConfigureRadioModule(serialPort io.ReadWriteCloser) error {
	// Access GPIO headers
	err := gpio.Open()
	if err != nil {
		return err
	}

	defer gpio.Close()

	setPin := gpio.NewPin(gpio.GPIO8)
	defer setPin.Input()
	setPin.Output()
	setPin.Low()
	time.Sleep(50 / 1000)

	// Write to the radio module to set the right configuration
	_, err = serialPort.Write([]byte("WR 434000 1 4 3 0"))
	if err != nil {
		return err
	}
	_, err = serialPort.Write([]byte{0x0D, 0x0A})
	if err != nil {
		return err
	}

	time.Sleep(10 / 1000)
	setPin.High()

	return nil
}