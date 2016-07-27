package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var port = flag.Int("port", 8080, "set the server listening port")

var supportGenerators = [...]GraphGenerator{
	&GraphvizGenerator{},
	&PlantUmlGenerator{},
}

var formatMineTypeMap = map[string]string{
	"svg": "image/svg+xml; charset=utf-8",
	"png": "image/png",
}

var installedGenerators []GraphGenerator

func main() {
	flag.Parse()
	CheckGenerators()
	StartWeb()
}

func StartWeb() {
	http.HandleFunc("/test/graphviz", TestGraphviz)
	http.HandleFunc("/test/plantuml", TestPlantUml)
	http.HandleFunc("/g", GeneratrPng)
	http.HandleFunc("/svg", GeneratrSvg)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(*port), nil))
}

func CheckGenerators() {
	for _, generator := range supportGenerators {
		if generator.CheckEnv() {
			installedGenerators = append(installedGenerators, generator)
		}
	}
	if len(installedGenerators) == 0 {
		log.Fatal("No Installed Graph Generator found!")
	}
	log.Printf("Enabled graph generators: %s\n", installedGenerators)
}

func GetCompatibleGenerator(str string) GraphGenerator {
	for _, generator := range installedGenerators {
		if generator.IsCompatible(str) {
			return generator
		}
	}
	return nil
}

func GenerateGraphWithCache(w http.ResponseWriter, r *http.Request, outputType string) {
	mineType := formatMineTypeMap[outputType]
	w.Header().Set("Content-Type", mineType)

	outputFromCache, exist := Cache.Get(r.RequestURI)
	var output []byte
	if exist {
		output, _ = outputFromCache.([]byte)
	} else {
		str := GetGraphString(r)
		output = GenerateGraph(str, outputType)
		Cache.Add(r.RequestURI, output)
	}
	w.Write(output)
}

func GenerateGraph(str string, outputType string) []byte {
	generator := GetCompatibleGenerator(str)
	if generator != nil {
		return generator.GenerateFromString(str, outputType)
	} else {
		return ShowError("Error: Can't Parse Input Content!", outputType)
	}
}

func GeneratrPng(w http.ResponseWriter, r *http.Request) {
	GenerateGraphWithCache(w, r, "png")

}

func GeneratrSvg(w http.ResponseWriter, r *http.Request) {
	GenerateGraphWithCache(w, r, "svg")
}

func GetGraphString(r *http.Request) string {
	str, _ := url.QueryUnescape(r.URL.RawQuery)
	log.Printf("GetGraphString: %s\n", str)
	return str
}

func TestGraphviz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	str := "digraph G {T [label=\"Graphviz Works\"]}"
	generator := GetCompatibleGenerator(str)
	w.Write(generator.GenerateFromString(str, "svg"))
}

func TestPlantUml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	str := "digraph G {T [label=\"PlantUML Works\"]}"
	generator := GetCompatibleGenerator(str)
	w.Write(generator.GenerateFromString(str, "svg"))
}
