package main

import (
	"fmt"
	"github.com/ilhamtubagus/go_patterns/functional_options/server"
)

func main() {
	// Here we use functional options to create a server with default options
	s := server.NewServer(
		server.WithMaxConnections(200),
		server.WithID("ID"),
		server.WithTLS,
		// Add more options as needed...
	)

	fmt.Printf("%#v", s)
}
