package src

import (
	"net/http"
)

type Home struct {
}

func (h *Home) Home(w http.ResponseWriter, r *http.Request) {
	// helpers.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
