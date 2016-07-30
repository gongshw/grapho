package main

import (
	"flag"
	lru "github.com/hashicorp/golang-lru"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var port = flag.Int("port", 8080, "set the server listening port")
var cacheSize = flag.Int("cache", 128, "set the cache size")

var supportGenerators = [...]GraphGenerator{
	&GraphvizGenerator{},
	&PlantUmlGenerator{},
}

var Cache *lru.Cache

var formatMineTypeMap = map[string]string{
	"svg": "image/svg+xml; charset=utf-8",
	"png": "image/png",
}

var installedGenerators []GraphGenerator

func main() {
	flag.Parse()
	InitCache()
	CheckGenerators()
	StartWeb()
}

func InitCache() {
	if *cacheSize > 0 {
		Cache, _ = lru.New(*cacheSize)
	}
}

func StartWeb() {
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

func GenerateGraphWithCache(w http.ResponseWriter, r *http.Request, outputType string) {
	mineType := formatMineTypeMap[outputType]
	w.Header().Set("Content-Type", mineType)

	if Cache == nil {
		GenerateGraphWithOutCache(w, r, outputType)
		return
	}

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

func GenerateGraphWithOutCache(w http.ResponseWriter, r *http.Request, outputType string) {
	var output []byte
	str := GetGraphString(r)
	output = GenerateGraph(str, outputType)
	w.Write(output)
}

func GenerateGraph(str string, outputType string) []byte {
	for _, g := range installedGenerators {
		log.Printf("Try %s", g)
		output, err := g.TryGenerateFromString(str, outputType)
		if err == nil {
			return output
		}
	}
	return ShowError("Error: Can't Parse Input Content!", outputType)
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
