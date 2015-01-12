package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/nu7hatch/gouuid"
	"time"
)

var ch = make(chan string)

func saver(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		utm := <-c
		fmt.Println(utm)
	}
}

// go run server.go
func main() {
	go saver(ch)
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
	id, err := uuid.NewV4()
	if (err == nil) {
		ch <- params["utm"]
		redirect_url := fmt.Sprintf("http://%v/?id=%v", params["url"], id)
		r.Redirect(redirect_url, 302)
	}
}
