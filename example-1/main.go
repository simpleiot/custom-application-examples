package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/simpleiot/simpleiot"
	"github.com/simpleiot/simpleiot/store"
)

func main() {
	log.Println("My app that embeds siot")

	siotOptions := simpleiot.Options{
		StoreType: store.StoreTypeBolt,
		DataDir:   "./",
		HTTPPort:  "8080",
	}

	// start simpleiot datastore. Start returns a nats conn
	// that can be used with helper functions in the simpleiot/nats
	// package.
	siot := simpleiot.NewSiot(siotOptions)

	nc, err := siot.Start()

	// this is not used yet
	_ = nc

	if err != nil {
		log.Fatal("Error starting SIOT store: ", err)
	}

	log.Println("SIOT started, open up UI at http://localhost:8080")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println()
		log.Println(sig)
		done <- true
	}()

	log.Println("running ...")
	<-done
	// cleanup
	log.Println("cleaning up ...")
	siot.Close()
	log.Println("exiting")
}
