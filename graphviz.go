package main

import (
	"bytes"
	"log"
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

func ExecGraphviz(dotString string, outputType string) []byte {
	dotCmd := exec.Command("dot", "-T", outputType)
	dotCmd.Stdin = strings.NewReader(dotString)
	var out bytes.Buffer
	dotCmd.Stdout = &out
	dotCmd.Stderr = &out
	err := dotCmd.Run()
	if err != nil {
		errMsg := out.String()
		return ErrorGraph(errMsg, outputType)
	}
	return out.Bytes()
}

func ErrorGraph(errMsg string, outputType string) []byte {
	errMsg = strings.Replace(errMsg, "\n", "\\n", -1)
	errMsg = strings.Replace(errMsg, "\"", "\\\"", -1)
	return ExecGraphviz("digraph G {T [label=\""+errMsg+"\", shape=box]}", outputType)
}
