package endpoints

import (
	// "fcserver/netguard"
	"fmt"
	"net/http"
)

type Gex struct {
}

func HellNo(w http.ResponseWriter, r *http.Request) error {
	if true {
		http.Error(w, "hell no", 569)
		return fmt.Errorf("hell naw")
	}
	return nil
}

func (g Gex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hai")
}
