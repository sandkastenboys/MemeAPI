package main

import (
	"log"
	"os"
	"strings"
)

var globalConf map[string]string


func loadConfig(){

	globalConf = make(map[string]string)

	globalConf["appHost"] = "0.0.0.0"
	globalConf["appPort"] = "8080"
	globalConf["dbHost"] = "0.0.0.0"
	globalConf["dbPort"] = "3306"
	globalConf["dbUser"] = "MemeAPI"
	globalConf["dbPassword"] = "MemesAREGreat"
	globalConf["dbDB"] = "memes"

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if _, ok := globalConf[pair[0]]; ok{
			log.Println("Set Config: ", pair[0], "=", pair[1])
			globalConf[pair[0]] = pair[1]
		}
	}
}
