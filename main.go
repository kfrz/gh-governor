package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/kfrz/gh-governor/cmd"
	_ "github.com/kfrz/gh-governor/config"
)

// init is called before main to initialize the logger
func init() {}

func main() {
	zap.L().Debug("main function executing")
	if err := cmd.Execute(); err != nil {
		zap.L().Fatal("Error executing command", zap.Error(err))
		os.Exit(1)
	}
}
