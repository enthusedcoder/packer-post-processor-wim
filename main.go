package main

import (
	"github.com/attacktac/packer-post-processor-wim/wim"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(wim.PostProcessor))
	server.Serve()
}
