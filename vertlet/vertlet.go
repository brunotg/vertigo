package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/vmarmol/vertigo/gce"
	"github.com/vmarmol/vertigo/vertlet/migration"
)

var argPort = flag.Int("port", 8080, "port to listen")

func main() {
	flag.Parse()
	gceService, err := gce.NewCompute()
	if err != nil {
		log.Fatal(err)
	}

	migration.RegisterHandlers(gceService)

	log.Print("About to serve on port ", *argPort)
	addr := fmt.Sprintf(":%v", *argPort)
	log.Fatal(http.ListenAndServe(addr, nil))
}