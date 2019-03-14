package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
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

func set() error {
	resolveFile, err := os.Open(*resolv)
	if err != nil {
		panic(err)
	}
	defer resolveFile.Close()

	configFile, err := os.Open(*config)
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

	configString := builder.String()

	fmt.Println(configString)
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
	if isSet {
		fmt.Println("Mode : Set")
	} else {
		fmt.Println("Mode : Unset")
	}

	if isSet && *config == "" {
		panic("Please provide a valid configuration file via the -c flag.")
	}
	fmt.Printf("Using : %v\n", config)
	set()
}
