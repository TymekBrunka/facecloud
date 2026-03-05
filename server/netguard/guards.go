package netguard

import (
	"fmt"
	"net/http"
)

// for best use, use SameSite cookies to prevent unauthorized actions in case the request header gets added manualy
func G_CSRF_simple(w http.ResponseWriter, r *http.Request) error {
	origin := r.Header.Get("sec-fetch-site")
	if origin == "same-origin" || origin == "same-site" {
		return nil
	}

	http.Error(w, "CSRF detected", http.StatusForbidden)
	return fmt.Errorf("CSRF detected")
}
