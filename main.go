package main

import (
	"gemtext/g2h"
	"os"
)

func main() {
	g2h.GemToHtml(os.Stdin, os.Stdout)
}
