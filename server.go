package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// go run server.go
func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/ya", func(r render.Render) {
		r.Redirect("http://ya.ru", 302)
	})
	m.Run()
}
