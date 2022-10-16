package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	inputFile *string
	inputDir  *string
	outputDir *string
)

func init() {
	inputFile = flag.String("inputFile", "NoInputFile", "The pdbqt file which contains many models.")
	inputDir = flag.String("inputDir", "NoInputDir", "The Directory of pdbqt file which contains many models.")
	outputDir = flag.String("outputDir", "./output/", "The output directory to save output models.")
}

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	content, err := io.ReadAll(file)
	return string(content)
}

func splitModels(content string) []string {
	var models []string
	models = strings.Split(content, "ENDMDL")
	return models[:len(models)-1]
}

func getZINC(models []string) map[string]string {
	var modelsMap map[string]string
	modelsMap = make(map[string]string)
	findZINC, _ := regexp.Compile("ZINC[0-9]*")
	for _, model := range models {
		modelsMap[findZINC.FindString(model)] = strings.TrimPrefix(model, "\n") + "ENDMDL\n"
	}

	return modelsMap
}

func writeModel(zinc string, filePath string, content string) {
	file, err := os.Create(filePath + zinc + ".pdbqt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	start := time.Now()

	var inputFiles []string

	if *inputFile == "NoInputFile" && *inputDir == "NoInputDir" {
		panic("No inputFile or inputDir!")
	}

	if *inputDir != "NoInputDir" {
		inputDir := *inputDir
		inputFiles, _ = filepath.Glob(inputDir + "/*.pdbqt")
	}

	if *inputFile != "NoInputFile" {
		inputFiles = append(inputFiles, *inputFile)
	}

	var num int

	for _, fileName := range inputFiles {
		fmt.Printf("Loading %s\n", fileName)
		modelsMap := getZINC(splitModels(readFile(fileName)))
		fmt.Printf("Get %d models from %s\n", len(modelsMap), fileName)
		for zinc, content := range modelsMap {
			writeModel(zinc, *outputDir, content)
			num += 1
		}
	}
	elapsed := time.Now().Sub(start)
	fmt.Printf("Get %d models in total. Finished in %s.\n", num, elapsed)
}
