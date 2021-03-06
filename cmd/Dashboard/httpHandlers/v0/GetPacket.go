package v0

import (
	"encoding/json"
	"fmt"
	"github.com/Autonomeasure/GroundStation/pkg/Memory"
	"net/http"
	"strconv"
)

func GetPacket(w http.ResponseWriter, r *http.Request) {
	// Get the 'packet' URI parameter
	p, ok1 := r.URL.Query()["packet"]
	l, ok2 := r.URL.Query()["last"]
	if !(ok1 || ok2) {
		// Return HTTP error to the user
		fmt.Fprintf(w, "{ \"error\": true }")
		return
	}

	if p != nil {
		pId, _ := strconv.ParseUint(p[0], 10, 32)

		packet, err := Memory.Database.GetRadioPacket(uint32(pId))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(packet)
		return
	} else if l != nil {
		lastID, _ := strconv.ParseUint(l[0], 10, 32)

		packets, err := Memory.Database.GetRadioPacketsFrom(uint32(lastID))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(packets)
		return
	} else {
		fmt.Fprintf(w, "{ \"error\": true }")
		return
	}

}

func GetBMPTemperature(w http.ResponseWriter, r *http.Request) {
	// Get the 'packet' URI parameter
	p, ok1 := r.URL.Query()["packet"]
	l, ok2 := r.URL.Query()["last"]
	if !(ok1 || ok2) {
		// Return HTTP error to the user
		fmt.Fprintf(w, "{ \"error\": true }")
		return
	}

	if p != nil {

	} else if l != nil {
		pId, _ := strconv.ParseUint(l[0], 10, 32)

		bmpTemps, ids, err := Memory.Database.GetBMPTemperatureFrom(uint32(pId))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		type Message struct {
			BMPTemps 	[]float32 	`json:"bmpTemps"`
			IDs			[]int 		`json:"IDs"`
		}

		var msg Message
		msg.BMPTemps = bmpTemps
		msg.IDs = ids
		json.NewEncoder(w).Encode(msg)
	}
}

func GetMPUTemperature(w http.ResponseWriter, r *http.Request) {
	// Get the 'packet' URI parameter
	p, ok1 := r.URL.Query()["packet"]
	l, ok2 := r.URL.Query()["last"]
	if !(ok1 || ok2) {
		// Return HTTP error to the user
		fmt.Fprintf(w, "{ \"error\": true }")
		return
	}

	if p != nil {

	} else if l != nil {
		pId, _ := strconv.ParseUint(l[0], 10, 32)

		bmpTemps, ids, err := Memory.Database.GetMPUTemperatureFrom(uint32(pId))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		type Message struct {
			BMPTemps 	[]float32 	`json:"mpuTemps"`
			IDs			[]int 		`json:"IDs"`
		}

		var msg Message
		msg.BMPTemps = bmpTemps
		msg.IDs = ids
		json.NewEncoder(w).Encode(msg)
	}
}

func GetPressure(w http.ResponseWriter, r *http.Request) {
	// Get the 'packet' URI parameter
	p, ok1 := r.URL.Query()["packet"]
	l, ok2 := r.URL.Query()["last"]
	if !(ok1 || ok2) {
		// Return HTTP error to the user
		fmt.Fprintf(w, "{ \"error\": true }")
		return
	}

	if p != nil {

	} else if l != nil {
		pId, _ := strconv.ParseUint(l[0], 10, 32)

		pressures, ids, err := Memory.Database.GetPressureFrom(uint32(pId))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		type Message struct {
			Pressures 	[]float32 	`json:"pressures"`
			IDs			[]int 		`json:"IDs"`
		}

		var msg Message
		msg.Pressures = pressures
		msg.IDs = ids
		json.NewEncoder(w).Encode(msg)
	}
}