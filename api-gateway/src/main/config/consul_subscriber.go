package config

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
)

// RegisterService registers service loaded from yml def
func RegisterService(service Service) {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Create an instance representing this service.
	// name of _this_ service. The service should be cleaned up via Close.
	svc, _ := connect.NewService(service.Name, client)
	defer svc.Close()

	// Creating an HTTP server that serves via Connect
	server := &http.Server{
		Addr:      fmt.Sprintf(":%v", strconv.Itoa(service.Port)),
		TLSConfig: svc.ServerTLSConfig(),
	}

	// Serve!
	server.ListenAndServeTLS("", "")
}
