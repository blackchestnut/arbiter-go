package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"time"
)

type Data struct {
	id  int64
	url string
	utm string
}

func (d *Data) String() string {
	return fmt.Sprintf("%#v", d)
}

func (d *Data) RedirectUrl() string {
	return fmt.Sprintf("http://%v/?id=%v", d.url, d.id)
}

var ch = make(chan *Data)

func saver(c chan *Data) {
	for {
		data := <-c
		fmt.Println(data)
	}
}

func GenerateID() int64 {
	return time.Now().UnixNano()
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
	data := Data{GenerateID(), params["url"], params["utm"]}
	ch <- &data
	r.Redirect(data.RedirectUrl(), 302)
}

func GetHtml(r render.Render, params martini.Params) {
	data := Data{GenerateID(), params["url"], params["utm"]}
	ch <- &data
	r.HTML(200, "html", data.RedirectUrl())
}
