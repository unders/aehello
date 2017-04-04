package secure

import "net/http"

const (
	hKey     = "Strict-Transport-Security"
	hValue   = "max-age=31536000; includeSubDomains"
	protocol = "X-FORWARDED-PROTO"
)

type Handler struct {
	Mux http.Handler

	// Only set this header when you have a custom domain
	// https://www.owasp.org/index.php/HTTP_Strict_Transport_Security_Cheat_Sheet
	SetHeader bool

	// https://example.com:443
	// https://project-id.appspot.com:443
	RedirectTo string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p := r.Header.Get(protocol); p == "http" {
		if h.SetHeader {
			// Once a supported browser receives this header that browser
			// will prevent any communications from being sent over HTTP to
			// the specified domain and will instead send all communications
			// over HTTPS. It also prevents HTTPS click through prompts on browsers.
			w.Header().Set(hKey, hValue)
		}
		http.Redirect(w, r, h.RedirectTo+r.RequestURI, http.StatusMovedPermanently)
		return
	}

	h.Mux.ServeHTTP(w, r)
}
