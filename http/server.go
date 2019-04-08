package http

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"greetings-service/configuration"
	"log"
	"net/http"
	"time"
)

type Server struct {
	*log.Logger
	*mux.Router
	ServerPort string
	CertFile   string
	KeyFile    string
}

func NewServer(configurationProvider configuration.Provider, logger *log.Logger, router *mux.Router) *Server {
	return &Server{
		logger,
		router,
		configurationProvider.GetStringOr("SERVER_PORT", "8080"),
		configurationProvider.GetString("CERT_FILE"),
		configurationProvider.GetString("KEY_FILE"),
	}
}

func (s *Server) Start() {
	server := New(s.Router, s.ServerPort)
	if s.CertFile == "" || s.KeyFile == "" {
		s.Logger.Printf("Starting HTTP server on %v", s.ServerPort)
		s.Logger.Fatal(
			"Startup of server failed",
			server.ListenAndServe())
	} else {
		s.Logger.Printf("Starting HTTPS server on %v", s.ServerPort)
		s.Logger.Fatal(
			"Startup of server failed",
			server.ListenAndServeTLS(s.CertFile, s.KeyFile))
	}
}

func New(mux *mux.Router, serverPort string) *http.Server {
	// See https://blog.cloudflare.com/exposing-go-on-the-internet/ for details
	// about these settings
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default cipher suite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
	srv := &http.Server{
		Addr:         ":" + serverPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      mux,
	}
	return srv
}
