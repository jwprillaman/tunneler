package main

import(
	"flag"
	"fmt"
	"log"
	"os"
)

func main(){

	var config = flag.String("c", "", "Configuation file to mimic.")
	flag.Parse()
	if *config == "" {
        log.Fatal("Please provide a valid configuration file via the -c flag.")
	} else { //continue normally
        fmt.Printf("Using : %v\n", *config)
        //Open config
        configFile,err := os.Open(*config)
        if err != nil {
            log.Fatal(err)
        } else {
            //continue
            var configFileInfo,configFileInfoErr = configFile.Stat()
            if configFileInfoErr != nil {
                log.Fatal(configFileInfoErr)
            } else {
                configData := make([]byte, configFileInfo.Size())
                byteCount, err := configFile.Read(configData)
                if err != nil {
                        log.Fatal(err)
                } else {
                    configString := string(configData[:byteCount])
                    fmt.Printf("%v", configString)
                }
            }
        }
	}

}
