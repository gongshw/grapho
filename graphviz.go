package main

import (
	"log"
	"os/exec"
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

func (GraphvizGenerator) GenerateFromString(str string, outputType string) []byte {
	output, err := ExecGraphviz(str, outputType)
	if err != nil {
		return ShowError(err.Error(), outputType)
	}
	return output
}

func (GraphvizGenerator) IsCompatible(str string) bool {
	_, error := ExecGraphviz(str, "png")
	return error == nil
}

func (GraphvizGenerator) String() string {
	return "GraphvizGenerator"
}
