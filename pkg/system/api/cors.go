package api

import (
	"log"
	"net/http"
	"net/url"

	"github.com/DeviaVir/instio/pkg/system/db"
)

// sendPreflight is used to respond to a cross-origin "OPTIONS" request.
func sendPreflight(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(200)
	return
}

func responseWithCORS(res http.ResponseWriter, req *http.Request) (http.ResponseWriter, bool) {
	if db.ConfigCache("cors_disabled").(bool) == true {
		// Check origin matches config domain.
		domain := db.ConfigCache("domain").(string)
		origin := req.Header.Get("Origin")
		u, err := url.Parse(origin)
		if err != nil {
			log.Println("Error parsing URL from request Origin header:", origin)
			return res, false
		}

		// Hack to get dev environments to bypass cors since u.Host (below) will
		// be empty, based on Go's url.Parse function.
		if domain == "localhost" {
			domain = ""
		}
		origin = u.Host

		// Currently, this will check for exact match. will need feedback to
		// determine if subdomains should be allowed or allow multiple domains
		// in config.
		if origin == domain {
			// Apply limited CORS headers and return.
			res.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
			res.Header().Set("Access-Control-Allow-Origin", domain)
			return res, true
		}

		// Disallow request.
		res.WriteHeader(http.StatusForbidden)
		return res, false
	}

	// Apply full CORS headers and return.
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	return res, true
}

// CORS wraps a HandlerFunc to respond to OPTIONS requests properly.
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return db.CacheControl(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res, cors := responseWithCORS(res, req)
		if !cors {
			return
		}

		if req.Method == http.MethodOptions {
			sendPreflight(res)
			return
		}

		next.ServeHTTP(res, req)
	}))
}
