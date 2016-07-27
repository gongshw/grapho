package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type GraphGenerator interface {
	String() string
	CheckEnv() bool
	GenerateFromString(str string, outputType string) []byte
	IsCompatible(str string) bool
}

func ExecGraphviz(dotString string, outputType string) ([]byte, error) {
	return Exec("dot", dotString, "-T", outputType)
}

func Exec(cmd string, inputStr string, args ...string) ([]byte, error) {
	dotCmd := exec.Command(cmd, args...)
	if inputStr != "" {
		dotCmd.Stdin = strings.NewReader(inputStr)
	}
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

func ShowError(errMsg string, outputType string) []byte {
	errMsg = strings.Replace(errMsg, "\n", "\\n", -1)
	errMsg = strings.Replace(errMsg, "\"", "\\\"", -1)
	output, _ := ExecGraphviz("digraph G {T [label=\""+errMsg+"\", shape=box]}", outputType)
	return output
}
