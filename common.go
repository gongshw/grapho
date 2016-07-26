package main

type GraphGenerator interface {
	String() string
	CheckEnv() bool
	GenerateFromString(str string, outputType string) []byte
	IsCompatible(str string) bool
}
