package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	file, scanner := getInputFileContent()
	defer file.Close()

	var arr []int64
	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		arr = append(arr, x)
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 2020 {
				fmt.Printf("Multiplied value: %v", (arr[i] * arr[j]))
			}
		}
	}
}

func getInputFileContent() (*os.File, *bufio.Scanner) {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(fmt.Sprint("error getting executable path. err: ", err.Error()))
	}

	if strings.Contains(ex, "tmp/go-build") {
		_, currentFilePath, _, _ := runtime.Caller(0)
		ex = currentFilePath
	}

	dirpath := filepath.Dir(ex)
	file, err := os.Open(fmt.Sprint(dirpath, "/input.txt"))
	if err != nil {
		log.Fatal(fmt.Sprint("getting file. err: ", err.Error()))
	}

	return file, bufio.NewScanner(file)
}
