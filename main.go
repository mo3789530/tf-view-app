package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
)

//go:embed ui/build
var frontend embed.FS

func main() {
	fe, err := fs.Sub(frontend, "ui/build")
	if err != nil {
		log.Fatalln(err)
	}
	frontendFS := http.FileServer(http.FS(fe))

	err = startServer("0.0.0.0:9000", frontendFS)
	if err != nil {
		log.Fatalln(err)
	}
}

func startServer(ipPort string, frontendFS http.Handler) error {
	m := http.NewServeMux()
	server := http.Server{Addr: ipPort, Handler: m}

	m.Handle("/", frontendFS)
	m.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// simple healthcheck
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"alive": true}`)
	})

	listen, err := net.Listen("tcp", ipPort)

	log.Printf("tf-view-app is running on %s", ipPort)

	if err != nil {
		log.Fatalln(err)
	}

	return server.Serve(listen)
}
