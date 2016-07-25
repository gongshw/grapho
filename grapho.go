package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var port = flag.Int("port", 8080, "set the server listening port")

func main() {
	flag.Parse()
	CheckGraphvizVersion()
	StartWeb()
}

func StartWeb() {
	http.HandleFunc("/test/graphviz", TestGraphviz)
	http.HandleFunc("/g", GeneratrPng)
	http.HandleFunc("/svg", GeneratrSvg)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(*port), nil))
}

func GeneratrPng(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(ExecGraphviz(GetGraphString(r), "png"))

}

func GeneratrSvg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Write(ExecGraphviz(GetGraphString(r), "svg"))
}

func GetGraphString(r *http.Request) string {
	str, _ := url.QueryUnescape(r.URL.RawQuery)
	log.Printf("GetGraphString: %s\n", str)
	return str
}

func TestGraphviz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Write(ExecGraphviz("digraph G {T [label=\"Graphviz Works\"]}", "svg"))
}
