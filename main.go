package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	gorest "github.com/inquizarus/gorest"
)

// Response is a basic response structure that can
// be Marshaled into bytes for writing in a ResponseWriter
type Response struct {
	Message string `json:"message"`
}

func main() {
	tlsConfig := gorest.ServeTLSConfig{
		Enabled:          true,
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
				Get:  makeGetHandler("Hello, World!"),
			},
		},
	}
	gorest.Serve(config)
}

func makeGetHandler(reply string) gorest.VerbHandler {
	return func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		var response Response
		defer r.Body.Close()
		response.Message = reply
		bodyBytes, _ := json.Marshal(response)
		w.Write([]byte(bodyBytes))
	}
}
