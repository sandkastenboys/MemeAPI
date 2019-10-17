package main

import (
	"fmt"
	"net/http"
)

func apiSetup(){
	apiSetupHandler();

	err := http.ListenAndServe(globalConf["appHost"] + ":" + globalConf["appPort"], nil);
	if err != nil{
		fmt.Errorf("Can not setup socket on that port and address");
	}

}

func apiSetupHandler(){
	http.HandleFunc("/meme/size/", apiGetSize);
	http.HandleFunc("/meme/get/", apiMemeDatabyID)
	http.HandleFunc("/meme/search/", apiSearchMemebyTag);
	http.HandleFunc("/social/get/", apiUserDatabyID);
	// nenne den nächsten /social/
}