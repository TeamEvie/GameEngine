package main

import (
	"eviecoin/client"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	color.Magenta("[eviecoin] Starting...")
	client.GetClient()
	color.Magenta("[eviecoin] Started!")

	fmt.Println("Started! - Press CTRL + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc
}
