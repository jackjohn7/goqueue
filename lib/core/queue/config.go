package queue

import (
	"crypto/tls"
	"net"
)

type Config struct {
	// Address on which Goqueue will listen for connections
	Addr net.Addr
	// TLS Certificate for encrypted connections
	Cert *tls.Certificate
	// Disables TLS certificate verification (used for local development and testing)
	InsecureSkipVerify bool
}

// Creates a net.Listener from the provided configurations
func (cfg *Config) CreateListener() (net.Listener, error) {
	if cfg.Cert != nil {
		listenerConfig := &tls.Config{Certificates: []tls.Certificate{*cfg.Cert}, InsecureSkipVerify: cfg.InsecureSkipVerify}
		listener, err := tls.Listen("tcp", cfg.Addr.String(), listenerConfig)
		return listener, err
	}

	listener, err := net.Listen("tcp", cfg.Addr.String())
	return listener, err
}
