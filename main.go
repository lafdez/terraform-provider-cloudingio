package main

import (
	"context"
	"flag"
	"log"

	"github.com/lafdez/terraform-provider-cloudingio"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string = "dev"
)

func main() {
	opts := providerserver.ServerOpts {
		Address: "github.com/lafdez/terraform-provider-cloudingio",
		Debug: debug,
	}

	err := providerserver.Serve(context,.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
