package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type PlantUmlGenerator struct {
	javaCmd     string
	plantUmlJar string
}

func (g *PlantUmlGenerator) CheckEnv() bool {
	javaHome := os.Getenv("JAVA_HOME")
	if javaHome == "" {
		path, err := exec.LookPath("java")
		if err != nil {
			log.Println("CheckEnv: java not found! Please check the $JAVA_HOME or make sure a java command in your $PATH")
			return false
		}
		g.javaCmd = path
	} else {
		log.Printf("CheckEnv: Found $JAVA_HOME: %s\n", javaHome)
		g.javaCmd = javaHome + "/bin/java"
	}
	output, err := Exec(g.javaCmd, "", "-version")
	log.Printf("Java Version:\n%s", output)
	if err != nil {
		return false
	}
	plantUmlJar := os.Getenv("PLANT_UML_JAR")
	if plantUmlJar == "" {
		log.Printf("CheckEnv: please set the $PLANT_UML_JAR the a valid plant UML jar!")
		return false
	}
	log.Printf("CheckEnv: Found $PLANT_UML_JAR: %s\n", plantUmlJar)
	g.plantUmlJar = plantUmlJar
	output, err = Exec(g.javaCmd, "", "-jar", plantUmlJar, "-version")
	if err != nil {
		log.Printf("Plant UML is invalid:\n%s", output)
		return false
	}
	log.Printf("Plant UML Version:\n%s", output)
	return err == nil
}

func (g *PlantUmlGenerator) TryGenerateFromString(str string, outputType string) ([]byte, error) {
	str = strings.Replace(str, ";", "\n", -1)
	return Exec(g.javaCmd, str, "-jar", g.plantUmlJar, "-p", "-t"+outputType, "-charset", "UTF-8")
}

func (*PlantUmlGenerator) String() string {
	return "PlantUmlGenerator"
}
