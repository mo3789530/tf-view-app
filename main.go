package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()
	server := http.Server{Addr: ipPort, Handler: router}

	// for ui handler
	// https://github.com/go-chi/chi/issues/403#issuecomment-468911337
	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("ui/build" + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, frontendFS).ServeHTTP(w, r)
		} else {
			frontendFS.ServeHTTP(w, r)
		}
	})
	// for health check api
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
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
