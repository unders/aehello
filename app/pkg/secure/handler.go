package secure

import "net/http"

const protocol = "X-FORWARDED-PROTO"

type Handler struct {
	Mux http.Handler

	// https://example.com:443
	// https://project-id.appspot.com:443
	RedirectTo string
	//Whitelist  []string
}

// https://www.owasp.org/index.php/HTTP_Strict_Transport_Security_Cheat_Sheet
//  Strict-Transport-Security: max-age=31536000; includeSubDomains
// Strict-Transport-Security: max-age=31536000; includeSubDomains

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p := r.Header.Get(protocol); p == "http" {
		http.Redirect(w, r, h.RedirectTo+r.RequestURI, http.StatusMovedPermanently)
		return
	}

	h.Mux.ServeHTTP(w, r)
}
