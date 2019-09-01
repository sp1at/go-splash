package main

import (
	"github.com/sp1at/go-splash/internal/args"
	"github.com/sp1at/go-splash/internal/config"
	"github.com/sp1at/go-splash/internal/output"
	"github.com/sp1at/go-splash/internal/request"
)

func main () {
	c := config.InitConfig()
	a := args.ProcessArgs()
	r := request.InitRequests(*c, *a)

	output.SaveOutput(r, *a)
}
