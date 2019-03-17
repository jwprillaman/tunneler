package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	commentToken = "#-=tunneler=-#"
)

var config *string
var resolv *string

func isSetMode() (bool, error) {
	output := false
	var err error = nil
	foundUnset := false

	flag.Visit(func(f *flag.Flag) {
		//look for set flag lexicographically
		if f.Name == "s" {
			output = true
		}
		//will hit set before unset so we can check output now
		if f.Name == "u" {
			foundUnset = true
		}
	})
	//return error if neither s or u is set
	if !output && !foundUnset {
		err = errors.New("Flag s to set or u to unset is a require parameter")
	}

	return output, err
}

func getConfigFile() string{
	output := ""

	if (*config != ""){
		output = *config
	} else {
		//look for config in current directory
		files, err := ioutil.ReadDir("./")
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".conf") {
				fullPath, err :=filepath.Abs(file.Name())
				if err != nil {
					panic(err)
				}
				output = fullPath
			}
		}
	}

	if output == "" {
		panic(errors.New("Unable to find config file"))
	}

	return output
}

func set() error {
	resolveFile, err := os.OpenFile(*resolv, os.O_RDWR, 0775)
	if err != nil {
		panic(err)
	}
	defer resolveFile.Close()

	configFile, err := os.Open(getConfigFile())
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	builder := strings.Builder{}
	resolvScanner := bufio.NewScanner(resolveFile)
	//write token followed by each line for current resolv
	for resolvScanner.Scan() {
		builder.WriteString(commentToken)
		builder.WriteString(resolvScanner.Text())
		builder.WriteString("\n")
	}
	//write config file
	configScanner := bufio.NewScanner(configFile)
	for configScanner.Scan() {
		builder.WriteString(configScanner.Text())
		builder.WriteString("\n")
	}
	//build string and set seek to start of file
	resolvString := builder.String()
	resolveFile.Seek(0, 0)
	//write the new resolve to the file
	writer := bufio.NewWriter(resolveFile)
	_, err = writer.WriteString(resolvString)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	return nil
}

func unset() error {
	resolveFile, err := os.OpenFile(*resolv, os.O_RDWR, 0775)
	if err != nil {
		panic(err)
	}
	defer resolveFile.Close()

	builder := strings.Builder{}
	resolvScanner := bufio.NewScanner(resolveFile)
	for resolvScanner.Scan() {
		text := resolvScanner.Text()
		if strings.HasPrefix(text, commentToken) {
			builder.WriteString(text[len(commentToken):])
			builder.WriteString("\n")
		}
	}

	resolvString := builder.String()
	resolveFile.Truncate(0)
	resolveFile.Seek(0, 0)
	//write new file
	writer := bufio.NewWriter(resolveFile)
	_, err = writer.WriteString(resolvString)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	return nil
}

func main() {

	//get flags
	config = flag.String("c", "", "Configuation file to mimic.")
	resolv = flag.String("r", "/etc/resolv.conf", "resolv.conf")
	flag.Bool("s", true, "set resolv to config")
	flag.Bool("u", true, "unset config in resolv")
	flag.Parse()
	isSet, err := isSetMode()
	if err != nil {
		panic(err)
	}
	//report success
	if isSet {
		err = set()
		fmt.Println("resolv set")
	} else {
		err = unset()
		fmt.Println("resolv unset")
	}
}
