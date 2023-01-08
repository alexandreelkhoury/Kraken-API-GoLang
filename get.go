package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://api.kraken.com/0/public/"
)

// the function get takes a string as parameter which is the
//rest of the url and returns the body of the response.

func get(rest string) []byte {
	mylink := url + rest
	response, err := http.Get(mylink)

	if err != nil {
		fmt.Printf("Error : %s", err)
		panic(err)
	} else {
		fmt.Println("Response Status : ", response.Status)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error : %s", err)
		panic(err)
	}
	return body
}

func getTime() *Time {

	body := get("Time")
	var result Time
	json.Unmarshal(body, &result)
	return &result
}

func getStatus() *Status {

	body := get("SystemStatus")
	var result Status
	json.Unmarshal(body, &result)
	return &result
}

func getAssets() *Assets {

	body := get("Assets")
	var result Assets
	json.Unmarshal(body, &result)
	return &result
}

func getAssetPairs() *AssetPairs {

	body := get("AssetPairs")
	var result AssetPairs
	json.Unmarshal(body, &result)
	return &result
}
