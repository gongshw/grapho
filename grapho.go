package main

import (
	"flag"
	"fmt"
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
	http.HandleFunc("/g", GeneratrGraph)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(*port), nil))
}

func GeneratrGraph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	fmt.Fprintf(w, ExecGraphviz(GetGraphString(r)))
}

func GetGraphString(r *http.Request) string {
	str, _ := url.QueryUnescape(r.URL.RawQuery)
	log.Printf("GetGraphString: %s\n", str)
	return str
}
