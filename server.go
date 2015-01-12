package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"time"
)

type Data struct {
	id  int64
	utm string
}

var ch = make(chan *Data)

func saver(c chan *Data) {
	for {
		time.Sleep(1 * time.Second)
		data := <-c
		fmt.Printf("%#v\n", data)
	}
}

func FullRedirectUrl(utm string, url string) string {
	id := time.Now().UnixNano()
	ch <- &Data{id, utm}
	return fmt.Sprintf("http://%v/?id=%v", url, id)
}

// go run server.go
func main() {
	go saver(ch)
	m := martini.Classic()

	// Setup routes
	r := martini.NewRouter()
	r.Get("/", GetLanding)
	r.Get("/test/:foo", GetTest)
	r.Get("/arbiter/:utm/redirect/:url", GetRedirect)
	r.Get("/arbiter/:utm/html/:url", GetHtml)
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

func GetRedirect(r render.Render, params martini.Params) {
	r.Redirect(FullRedirectUrl(params["utm"], params["url"]), 302)
}

func GetHtml(r render.Render, params martini.Params) {
	r.HTML(200, "html", FullRedirectUrl(params["utm"], params["url"]))
}
