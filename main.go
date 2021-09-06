package main

import "go-gearbox/api"

var (
	endpoints = api.GearboxEndpoint{}
)

func main() {
	endpoints.ALL()
	endpoints.Api().Start(":8080")
}
