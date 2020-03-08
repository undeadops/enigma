package server

import "net/http"

// Instance - Instance of the http server package
type Instance struct {
	httpServer *http.Server
}
