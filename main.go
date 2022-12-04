package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/LF-Certification/remote-desktop-ping-service/daemon"
)

func waitOsSignals(done chan bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs
	log.Printf("Receive %s OS signal", sig)

	done <- true
}

func main() {
	var url = flag.String("url", "0.0.0.0:3001", "Websocket URL")
	flag.Parse()

	if *url == "" {
		log.Fatal("Missing argument")
	}
	done := make(chan bool, 1)
	go waitOsSignals(done)
	fmt.Println(*url)

	srv := &http.Server{
		Addr:    *url,
		Handler: daemon.BuildHandler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Fatal(http.ListenAndServe(*url, nil))
	log.Println("BLABLABLA")
	<-done
	log.Println("The End.")
}
