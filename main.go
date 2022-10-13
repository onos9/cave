package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/cave/cmd/api"
	cfg "github.com/cave/config"
)

func main() {

	service := flag.String("s", "monolith", "service api, chat or stream")
	flag.Parse()

	cfg.LoadConfig()
	db := cfg.NewDBConnection()
	apiServer := api.Init(db)

	switch *service {
	case "monolith":
		apiServer.Run()
	case "api":
		apiServer.Run()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	apiServer.Shutdown()
}
