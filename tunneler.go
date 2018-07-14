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
                data := make([]byte, configFileInfo.Size())
                count, err := configFile.Read(data)
                if err != nil {
                        log.Fatal(err)
                }
                fmt.Printf("read %d bytes: %q\nWith total size : %v\n", count, data[:count], configFileInfo.Size())
            }
        }
	}

}
