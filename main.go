package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	dn := os.Getenv("DOMAIN_NAME")
	if dn == "" {
		panic("domain name cannot be empty")
	}
	tlsCrt := fmt.Sprintf("/certs/%s.crt", dn) // cert and key for HTTPS
	tlsKey := fmt.Sprintf("/certs/%s.key", dn)
	log.Printf("tls certificate path: %s tls private key path: %s\n", tlsCrt, tlsKey)

	crt := "/certs/token.crt" // use server cert to sign the token
	key := "/certs/token.key"
	log.Printf("token certificate path: %s token private key path: %s\n", crt, key)

	opt := &Option{
		TLSCertFile:     tlsCrt,
		TLSKeyFile:      tlsKey,
		Certfile:        crt,
		Keyfile:         key,
		TokenExpiration: time.Now().Add(24 * time.Hour).Unix(), // 24hrs
		TokenIssuer:     os.Getenv("AUTH_TOKEN_ISSUER"),        // must be aligned with what passed to registry in compose file
		Authenticator:   &DefaultAuthenticator{},               // could be nil, meaning all users would be authenticated by default
	}
	srv, err := NewAuthServer(opt)
	if err != nil {
		log.Fatal(err)
	}
	addr := ":" + os.Getenv("PORT")
	http.Handle("/auth", srv)
	log.Println("Server running at ", addr)
	if err := srv.Run(addr); err != nil {
		log.Fatal(err)
	}
}
