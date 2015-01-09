package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// go run server.go
func main() {
	m := martini.Classic()

	// Setup routes
	r := martini.NewRouter()
	r.Get("/", GetLanding)
	r.Get("/test/:foo", GetTest)
	r.Get("/arbiter/:utm/redirect/:url", GetArbiter)
	m.Action(r.Handle)

	m.Use(render.Renderer())
	m.Run()
}

func GetLanding() string {
	return "Hellow, I`m Arbiter \nTry: /arbiter/UTM_LABEL/redirect/URL \nExample: /arbiter/foo/redirect/ya.ru"
}

func GetTest(params martini.Params) string {
	return fmt.Sprintf("Test route params: %v, params['foo']: %v", params, params["foo"])
}

func GetArbiter(r render.Render, params martini.Params) {
	r.Redirect("http://"+params["url"], 302)
}
