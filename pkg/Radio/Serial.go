package Radio

import (
	"bufio"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

func OpenSerial(port string, baudrate uint) {
	options := serial.OpenOptions{
		PortName: port,
		BaudRate: baudrate,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
	}

	serialPort, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer serialPort.Close()
	reader := bufio.NewReader(serialPort)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
