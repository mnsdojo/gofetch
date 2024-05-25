package main

import (
	"fmt"

	"github.com/BigJk/imeji"
	// "github.com/mnsdojo/gofetch/internal/ascii"
	"github.com/mnsdojo/gofetch/internal/battery"
)

func main() {
	text, _ := imeji.FileString("./image.png", imeji.WithTrueColor())
	fmt.Println(text)

	battery.DisplayBatteryStatus()
}
