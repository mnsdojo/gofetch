package main

import (
	"fmt"

	"github.com/BigJk/imeji"
)

func main() {

	text, _ := imeji.FileString("./image.png", imeji.WithTrueColor())
	fmt.Println(text)
}
