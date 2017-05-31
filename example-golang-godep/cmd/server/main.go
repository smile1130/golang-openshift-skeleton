package main

import (
	"os"

	"github.com/amsokol/openshift-golang-template/example-golang-godep/pkg/cmd/cli"
	"github.com/amsokol/openshift-golang-template/example-golang-godep/pkg/cmd/server"
)

func main() {
	if err := server.Start(cli.GetPort()); err != nil {
		os.Exit(1)
	}
}
