package tls

import (
	"log"
	"net/http"
	"path/filepath"
  "fmt"

  "github.com/DeviaVir/instio/pkg/system/db"
	"github.com/DeviaVir/instio/pkg/system/cfg"
)

// EnableDev generates self-signed SSL certificates to use HTTPS & HTTP/2 while
// working in a development environment. The certs are saved in a different
// directory than the production certs (from Let's Encrypt), so that the
// acme/autocert package doesn't mistake them for it's own.
// Additionally, a TLS server is started using the default http mux.
func EnableDev() {
	setupDev()

	vendorPath := cfg.TlsDir()
	cert := filepath.Join(vendorPath, "devcerts", "cert.pem")
	key := filepath.Join(vendorPath, "devcerts", "key.pem")

	log.Fatalln(http.ListenAndServeTLS(fmt.Sprintf(":%s", db.ConfigCache("https_port").(string)), cert, key, nil))
}