package main

import (
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
		log.Fatal("CheckEnv: graphviz not found")
	}
	log.Printf("CheckEnv: graphviz is available at %s\n", path)
	output, err := Exec("dot", "", "-V")
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	log.Printf("CheckEnv: %s", output)
	return true
}

func (GraphvizGenerator) TryGenerateFromString(str string, outputType string) ([]byte, error) {
	if strings.HasPrefix(strings.TrimSpace(str), "@") {
		return nil, errors.New("Not Supoport")
	}
	return ExecGraphviz(str, outputType)
}

func (GraphvizGenerator) String() string {
	return "GraphvizGenerator"
}
