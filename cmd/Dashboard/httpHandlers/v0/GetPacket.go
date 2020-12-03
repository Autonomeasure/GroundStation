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
	product, ok := r.URL.Query()["packet"]
	if !ok || len(product[0]) < 1 {
		// Return HTTP error to the user
		fmt.Fprintf(w, "{ 'error': true }")
		return
	}

	pId, _ := strconv.ParseUint(product[0], 10, 32)

	packet, err := Memory.Database.GetRadioPacket(uint32(pId))
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(packet)
}
