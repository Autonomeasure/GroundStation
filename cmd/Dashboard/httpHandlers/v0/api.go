package v0

import (
	"fmt"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "CanSat competition 2020-2021 Team Autonomeasure | Dashboard API v0")
}
