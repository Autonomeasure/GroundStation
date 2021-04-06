package main

import (
	"bufio"
	"fmt"
	"github.com/J0eppp/go-nmea"
	"github.com/tarm/serial"
	"log"
	"strconv"
	"strings"
)

func main() {
	c := &serial.Config{Name: "COM5", Baud: 19200}
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
			parity, err := s.CheckParity()
			if err != nil {
				fmt.Println(err)
				fmt.Println(strconv.FormatUint(uint64(parity), 16))
			} else {
				fmt.Println("Parity: ", strconv.FormatUint(uint64(parity), 16))
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