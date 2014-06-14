package web

import (
	"log"
	"net/http"
	"os"
)

type Web struct {
	ip   string
	port string
}

var (
	name string = "Web"
)

func NewWeb() Web {
	return Web{ip: "localhost", port: ":8080"}
}

func (w *Web) Run() {
	LOG := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	LOG.Printf("%s Listening on %s%s", name, w.ip, w.port)
	http.HandleFunc("/ws", serveWs)
	http.Handle("/", http.FileServer(http.Dir("./web/static/")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
