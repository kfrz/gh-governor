package main

import (
	"go.uber.org/zap"

	"github.com/kfrz/gh-governor/cmd"
	_ "github.com/kfrz/gh-governor/config"
)

// init is called before main to initialize the config including logger
func init() {}

func main() {
	zap.L().Debug("main function executing")
	if err := cmd.Execute(); err != nil {
		zap.L().Fatal("Error executing command", zap.Error(err))
	}
}
