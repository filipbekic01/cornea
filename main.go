package main

import "github.com/filipbekic01/cornea/app"

import "os"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		_ = app.Run()
	} else {
		app.Manage(args)
	}
}
