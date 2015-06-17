package main

import (
	"github.com/franksong/logger"
)

func main() {
	logger.Debug("test debug!!!")
	logger.Info("test info!!!")
	logger.Warning("test warning!!!")
	logger.Error("test error!!!")
	// logger.Fatal("test fatal!!!")
}
