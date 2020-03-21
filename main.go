package main

import (
	"log"
	"net/http"
	"os"

	gorest "github.com/inquizarus/gorest"
)

func main() {
	tlsConfig := gorest.ServeTLSConfig{
		Enabled:          false,
		Strict:           false,
		Preload:          false,
		StrictSubDomains: false,
		CertPath:         "restexample.crt",
		KeyPath:          "restexample.key",
	}

	logger := log.New(os.Stdout, "restexample: ", log.LstdFlags)

	config := gorest.ServeConfig{
		Port:   "8080",
		TLS:    tlsConfig,
		Logger: logger,
		Middlewares: []gorest.Middleware{
			gorest.WithJSONContent(),
			gorest.WithStrictTransportSecurity(tlsConfig),
		},
		Handlers: []gorest.Handler{
			&gorest.BaseHandler{
				Path: "/",
				Get: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
					defer r.Body.Close()
					w.Write([]byte(`{"message": "Hello, World!"}`))
				},
			},
		},
	}
	gorest.Serve(config)
}
