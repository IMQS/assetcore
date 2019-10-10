package main

import (
	"github.com/IMQS/assetcore/core"
	"github.com/IMQS/nf"
)

func main() {
	svc := core.NewService()
	svc.LoadConfig()
	nf.RunService(svc.ListenAndServe)
}
