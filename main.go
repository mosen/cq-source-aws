package main

import (
	"context"
	"log"

	"github.com/cloudquery/plugin-sdk/v4/serve"
	plugin "github.com/mosen/cq-source-aws/resources/plugin"
)

func main() {
	if err := serve.Plugin(plugin.AWS()).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
