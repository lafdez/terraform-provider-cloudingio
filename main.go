package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/lafdez/terraform-provider-cloudingio/provider"
)

var (
	version string = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to enable debug")
	flag.Parse()

	opts := providerserver.ServeOpts {
		Address: "github.com/lafdez/terraform-provider-cloudingio",
		Debug: debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
