package main

import (
	"fmt"

	"github.com/mbndr/figlet4go"
	"github.com/tanigroup/jorah/config"
)

func main() {
	ascii := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{figlet4go.ColorRed}

	renderStr, _ := ascii.RenderOpts(config.Key.AppName, options)
	fmt.Print(renderStr)
	fmt.Print(config.Key.AppQuote)
	fmt.Printf("\nv%s", config.Key.Version)
}
