package main

import (
	"eviecoin/eviecoin"
	"github.com/fatih/color"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	color.Magenta("[eviecoin] Starting...")
	eviecoin.GetClient()
	color.Magenta("[eviecoin] Started!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc
}
