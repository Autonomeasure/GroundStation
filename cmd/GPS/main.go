package main

import (
	"bufio"
	"github.com/J0eppp/go-nmea"
	"github.com/tarm/serial"
	"log"
	"strings"
)

func main() {
	c := &serial.Config{Name: "COM5", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	//n, err := s.Write([]byte("test"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		msg := scanner.Text()
		if strings.HasPrefix(msg, "$GPGGA") {
			s := go_nmea.GPGGASentence{
				RawSentence: msg,
				Time: []uint16{0, 0, 0, 0},
			}
			s.Parse(1)
			//fmt.Println(msg) // Println will add back the final '\n'
			//fmt.Printf("%+v\n", s)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}