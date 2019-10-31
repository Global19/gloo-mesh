package main

import (
	"log"

	//check "github.com/solo-io/go-checkpoint"
	"github.com/solo-io/mesh-discovery/pkg/setup"
	//"github.com/solo-io/mesh-discovery/pkg/version"
)

func main() {

	if err := run(); err != nil {
		log.Fatalf("err in main: %v", err.Error())
	}
}

func run() error {
	errs := make(chan error)
	//start := time.Now()
	//check.CallCheck("meshdiscovery", version.Version, start)
	go func() {
		errs <- setup.Main(nil, nil)
	}()
	return <-errs
}
