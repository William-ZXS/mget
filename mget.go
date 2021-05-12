package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/gookit/color"
)

func getData(sourceFilePath string, keyWord string) {
	inputFile, inputError := os.Open(sourceFilePath)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	hasKeyWord := false
	paragraph := ""

	for {
		inputString, readerError := inputReader.ReadString('\n')

		// 判断是否有keyWord
		validKeyWord := regexp.MustCompile(keyWord)
		res := validKeyWord.MatchString(inputString)
		if res == true {
			hasKeyWord = true
			inputString = validKeyWord.ReplaceAllString(inputString, color.Red.Sprint(keyWord))

		}

		if inputString == "\n" {
			if hasKeyWord == true {
				paragraph += "file path: " + sourceFilePath + "\n"
				fmt.Println(paragraph)
				hasKeyWord = false
			}
			paragraph = ""

		} else {
			paragraph = paragraph + inputString
		}
		if readerError == io.EOF {
			return
		}
	}
}

func main() {

	//conf
	confPath := "mget.conf"

	//解析参数
	args := os.Args
	if args == nil || len(args) < 2 {
		fmt.Println("your key word?")
		return
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	absConfPath := path.Join(dir, confPath)

	keyWord := args[1]

	// confFileList := make([]string, 10)
	var confFileSlice []string
	confFile, confErr := os.Open(absConfPath)
	if confErr != nil {
		// fmt.Errorf("mget.conf not found")
		fmt.Println("mget.conf not found!")
		return
	}
	defer confFile.Close()
	confReader := bufio.NewReader(confFile)
	for {
		confString, readerErr := confReader.ReadString('\n')
		confFileSlice = append(confFileSlice, confString)

		if readerErr == io.EOF {
			break
		}
	}

	fmt.Println("数据源", confFileSlice)

	for _, sourceFilePath := range confFileSlice {
		getData(sourceFilePath, keyWord)
	}

}
