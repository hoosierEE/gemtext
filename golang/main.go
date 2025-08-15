package main

import (
	"gemtext_ref/g2h"
	"os"
)

func main() {
	g2h.GemToHtml(os.Stdin, os.Stdout)
}
