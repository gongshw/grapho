package main

import (
	"bytes"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func CheckGraphvizVersion() string {
	path, err := exec.LookPath("dot")
	if err != nil {
		log.Fatal("CheckGraphvizVersion: graphviz not found")
	}
	log.Printf("CheckGraphvizVersion: graphviz is available at %s\n", path)
	output, err := exec.Command("dot", "-V").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CheckGraphvizVersion: %s\n", output)
	return ""
}
func TestGraphviz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Write(ExecGraphviz("digraph G {T [label=\"Graphviz Works\"]}", "svg"))
}

func ExecGraphviz(dotString string, outputType string) []byte {
	dotCmd := exec.Command("dot", "-T", outputType)
	dotCmd.Stdin = strings.NewReader(dotString)
	var out bytes.Buffer
	dotCmd.Stdout = &out
	dotCmd.Stderr = &out
	err := dotCmd.Run()
	if err != nil {
		errMsg := out.String()
		return ExecGraphviz("digraph G {T [label=\""+errMsg+"\", shape=box]}", outputType)
	}
	return out.Bytes()
}
