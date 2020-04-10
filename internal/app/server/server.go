package server

import "net/http"

func NewApiServer(config *Configuration) error {
	NewHandle()
	err := http.ListenAndServe(config.Port, nil)
	return err
}
