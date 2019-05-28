package main

import (
	"bytes"

	"github.com/gokusenz/go-linenotify"
)

func main() {
	token := ""

	c := linenotify.New()
	c.Notify(token, "hello goku", "", "", nil)

}