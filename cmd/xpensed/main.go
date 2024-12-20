package main

import (
	"fmt"
	"os"

	"github.com/WlinkNET/xpense_chain/cmd/xpensed/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}