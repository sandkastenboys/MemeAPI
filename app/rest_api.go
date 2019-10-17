package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createError(message string)[]byte{
  p := map[string]string{"error": message}
  jsonData, _ := json.Marshal(p)
  return jsonData
}

func apiGetSize(w http.ResponseWriter, _ * http.Request){
	w.Header().Add("Content-Type", "application/json")

	data, err := querySize()

	if err == nil{
		jsonData, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(createError("Internal Server Error"))

}

func apiSearchMemebyTag(w http.ResponseWriter, r * http.Request){
	w.Header().Add("Content-Type", "application/json")

	request := getMemeSearch{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write(createError("invalid request"))
		return
	}

	data, err := queryByTag(request.tags);

	if err != nil{
		w.Write(createError("No Memes with this Tag"))
		return
	}

	jsonData, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func apiMemeDatabyID(w http.ResponseWriter, r * http.Request){
	w.Header().Add("Content-Type", "application/json")

	request := getByID{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(createError("invalid request"))
		return
	}

	data, err := queryMeme(request.ID);

	if err != nil{
		fmt.Print("nothing")
		w.Write(createError("No Meme with this ID"))
		return
	}
  	jsonData, _ := json.Marshal(data)
  	w.WriteHeader(http.StatusOK)
  	w.Write(jsonData)
}

func apiUserDatabyID(w http.ResponseWriter, r * http.Request){

}
