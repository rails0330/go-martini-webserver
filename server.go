package main

import "github.com/codegangsta/martini"

func main() { // test
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
