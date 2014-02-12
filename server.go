package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	// m.Get("true.com", func() string {
	// 	return m.Use(martini.Static("Dropbox/websites/true.com/www/htdocs/"))
	// })

	// m.Get("false.com", func() string {
	// 	return m.Use(martini.Static("Dropbox/websites/false.com/www/htdocs/"))
	// })

	http.ListenAndServe(":80", m)
}

// route incoming http requests
// if true.com
// serve /Dropbox/websites/true.com/www/htdocs/index.html
