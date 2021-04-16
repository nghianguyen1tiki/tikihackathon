package main

import (
	"github.com/nghiant3223/tikihackathon/internal"
	"github.com/nghiant3223/tikihackathon/internal/configs"
)

func main() {
	httpConfig := configs.GetHttpConfig()
	server := internal.NewServer(httpConfig)
	server.Start()
	defer server.Stop()
}
