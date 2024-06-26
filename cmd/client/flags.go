package main

import (
	"flag"
	"strconv"
)

type Flags struct {
	Server struct {
		Host *string
		Port *int
	}
}

func parseFlags() Flags {
	var flags Flags

	flags.Server.Host = flag.String("host", "localhost", "Server host")
	flags.Server.Port = flag.Int("port", 8081, "Server port")

	return flags
}

func (f Flags) GetServerURL() string {
	return *f.Server.Host + ":" + strconv.Itoa(*f.Server.Port)
}
