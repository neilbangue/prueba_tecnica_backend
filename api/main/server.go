package main

import (
	"flag"
	"server/mappings"
)

func main() {
	server := flag.String("addr", "", "Init Server")

	flag.Parse()

	if server != nil && len(*server) > 0 {
		mappings.CreateUrlMappings()
		mappings.Router.Run(*server)
	} else {
		println("Incorrect params")
	}

}
