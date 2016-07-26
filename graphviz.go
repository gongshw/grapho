package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

type GraphvizGenerator struct {
}

func (GraphvizGenerator) CheckEnv() bool {
	path, err := exec.LookPath("dot")
	if err != nil {
		log.Fatal("CheckGraphvizVersion: graphviz not found")
	}
	log.Printf("CheckGraphvizVersion: graphviz is available at %s\n", path)
	output, err := exec.Command("dot", "-V").CombinedOutput()
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	log.Printf("CheckGraphvizVersion: %s\n", output)
	return true
}

func (GraphvizGenerator) GenerateFromString(str string, outputType string) []byte {
	output, err := ExecGraphviz(str, outputType)
	if err != nil {
		errMsg := err.Error()
		errMsg = strings.Replace(errMsg, "\n", "\\n", -1)
		errMsg = strings.Replace(errMsg, "\"", "\\\"", -1)
		output, _ = ExecGraphviz("digraph G {T [label=\""+errMsg+"\", shape=box]}", outputType)
		return output
	}
	return output
}

func (GraphvizGenerator) IsCompatible(str string) bool {
	return true
}

func (GraphvizGenerator) String() string {
	return "GraphvizGenerator"
}

func ExecGraphviz(dotString string, outputType string) ([]byte, error) {
	dotCmd := exec.Command("dot", "-T", outputType)
	dotCmd.Stdin = strings.NewReader(dotString)
	var out bytes.Buffer
	dotCmd.Stdout = &out
	dotCmd.Stderr = &out
	err := dotCmd.Run()
	if err != nil {
		errMsg := out.String()
		return nil, errors.New(errMsg)
	}
	return out.Bytes(), nil
}
