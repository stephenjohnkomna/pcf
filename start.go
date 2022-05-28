package main

import (
	"pfc/api"
)

func main() {
	route := api.SetupRouter()
	route.Run()
}
