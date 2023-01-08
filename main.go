package main

import (
	"log"
	"flag"
	"net/http"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
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
		bridge:  make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func main() {

	var addr = flag.String("addr", ":8090", "server address:port")
	flag.Parse()

	g := newGroup()
	api := new(api)
	api.group = g
	http.Handle("/", &templateHandler{filename: "base.html"})
	http.Handle("/ws", g)
	http.Handle("/api", api)
	go g.run()

	log.Println("bind address: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
