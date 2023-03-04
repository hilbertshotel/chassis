package main

import (
	"chassis/dep"
	"chassis/handlers"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// initialize dependencies
	d, err := dep.Init()
	if err != nil {
		panic(err)
	}

	// initialize service
	service := http.Server{
		Addr:         d.Cfg.HostAddr,
		Handler:      handlers.Mux(d),
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}

	// start service
	ech := make(chan error)
	go func() {
		d.Log.Ok("Service start @ " + d.Cfg.HostAddr)
		ech <- service.ListenAndServe()
	}()

	// shutdown service
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	select {
	case <-shutdown:
		service.Shutdown(context.Background())
		d.Log.Ok("Service stop")
	case err := <-ech:
		d.Log.Error(err)
	}
}
