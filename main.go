package main

import (
	"log"
	"net/http"
	"text/template"
	"sync"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(t.filename))
	})
	t.templ.Execute(w, nil)
}

func newGroup() *group {
	return &group{
		bridge: make(chan []byte),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
	}
}

func main() {

	g := newGroup()
	api := new(api)
	api.group = g
	http.Handle("/", &templateHandler{filename : "base.html"})
	http.Handle("/ws", g)
	http.Handle("/api", api)
	go g.run()
	
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("LisntenAndServe: ", err)
	}
}
