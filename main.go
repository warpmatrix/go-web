package main

import (
	flag "github.com/spf13/pflag"
	"github.com/warpmatrix/cloud-go/service"
)

const (
	defPort string = "8080"
)

func main() {
	pPort := flag.StringP("port", "p", defPort, "PORT for httpd listening")
	flag.Parse()
	port := *pPort

	server := service.NewServer()
	server.Run(":" + port)
}
