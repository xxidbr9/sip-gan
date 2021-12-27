package main

import (
	"runtime"

	"github.com/xxidbr9/sip-gan/backend/enhancer/config"
)

func main() {
	if runtime.GOOS == "windows" {
		panic("Can't Execute this on a windows machine")
	}
	
	baseConfig, err := config.NewBaseConfig()
	if err != nil {
		panic(err)
	}

	if baseConfig.EnvMode == "production" {
		execute()
	}
	if baseConfig.EnvMode == "development" {
		runFaker()
	}
}
