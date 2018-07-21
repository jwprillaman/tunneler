package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	var config = flag.String("c", "", "Configuation file to mimic.")
	var resolv = flag.String("r", "/etc/resolv.conf", "resolv.conf")
	flag.Parse()
	if *config == "" {
		log.Fatal("Please provide a valid configuration file via the -c flag.")
	} else { //continue normally
		fmt.Printf("Using : %v\n", *config)
		//Open config
		configFile, err := os.Open(*config)
		if err != nil {
			log.Fatal(err)
		} else {
			//Open config file and read it
			var configFileInfo, configFileInfoErr = configFile.Stat()
			if configFileInfoErr != nil {
				log.Fatal(configFileInfoErr)
			} else {
				configData := make([]byte, configFileInfo.Size())
				configByteCount, err := configFile.Read(configData)
				if err != nil {
					log.Fatal(err)
				} else {
					configString := string(configData[:configByteCount])
					fmt.Printf("%v", configString)
					//Open /etc/resolv.conf
					resolvFile, err := os.OpenFile(*resolv, os.O_RDWR, 0755)
					if err != nil {
						log.Fatal(err)
					} else {
						var resolvFileInfo, resolvFileInfoErr = resolvFile.Stat()
						if resolvFileInfoErr != nil {
							log.Fatal(resolvFileInfoErr)
						} else {
							fmt.Printf("%v\n", resolvFileInfo.Size())
							resolvData := make([]byte, resolvFileInfo.Size())
							resolvByteCount, err := resolvFile.Read(resolvData)
							if err != nil {
								log.Fatal(err)
							} else {
								resolvString := string(resolvData[:resolvByteCount])
								fmt.Printf("%v", resolvString)
								resolvNewlineCount := strings.Count(resolvString, "\n")
								var resolvInsertIndices = make([]int, resolvNewlineCount+1)

								for i := 0; i < resolvNewlineCount; i++ {
									var testString string
									if i != 0 {
										testString = resolvString[0:resolvInsertIndices[i-1]]
									} else {
										testString = resolvString
									}
									lastIndex := strings.LastIndex(testString, "\n")
									resolvInsertIndices[i] = lastIndex
								}
								resolvInsertIndices[len(resolvInsertIndices)-1] = 0 //Add to the extra index added
								sort.Ints(resolvInsertIndices)
								fmt.Printf("%v\n", resolvInsertIndices)

								//time to write
								writeBytes := []byte("#-=Begone=-:")
								for i := 0; i < len(resolvInsertIndices); i++ {
									fmt.Printf("For index : %v\n", resolvInsertIndices[i])
									preChunkBytes := []byte(resolvString[0:resolvInsertIndices[i]])
									postChunkByteStartIndex := resolvInsertIndices[i]
									if resolvInsertIndices[i] != 0 {
										postChunkByteStartIndex++
									}
									postChunkBytes := []byte(resolvString[postChunkByteStartIndex:len(resolvString)])
									newPostChunkBytes := append(writeBytes, postChunkBytes...)
									fmt.Printf("writting : %v\n", string(newPostChunkBytes[:len(newPostChunkBytes)]))
									fmt.Printf("at the index of : %v\n", int64((len(preChunkBytes))+(i*len(writeBytes))))
									startWriteIndex := int64(len(preChunkBytes)) + (int64(i) * (int64(len(writeBytes))))

									written, err := resolvFile.WriteAt(newPostChunkBytes, startWriteIndex)
									if err != nil {
										log.Fatal(err)
									} else {
										fmt.Printf("Wrote : %v\n", written)
									}
								}
							}
						}
					}
				}
			}
		}
	}

}
