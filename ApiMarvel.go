/*

API : MARVEL.COM API
author:EDWIN NDUTI
DATE : JULY 2020

*/


package main

import(
	"fmt"
	"crypto/md5"
	"net/http"
	"io/ioutil"
	"encoding/hex"
)

//credentials
var (
	URI = "https://gateway.marvel.com/v1/public/"
	PUBLIC_KEY="a3c8827aa0350dce00ef650661e225cc"
	PRIVATE_KEY="91e76ebd17a6af61726a598b1496531e006d82dc"
	TIMESTAMP = "1"
	COMICS = "comics"
)

func main(){
	var HASH string

	//Create HASH 
	data := []byte(TIMESTAMP+PRIVATE_KEY+PUBLIC_KEY)

	//convert [16]byte to string
	b_hash := md5.Sum(data)
	newHash := b_hash[:]
	HASH = hex.EncodeToString(newHash)


	//GET endpoint
	newURL := URI+COMICS+"?ts="+TIMESTAMP+"&apikey="+PUBLIC_KEY+"&hash="+HASH

	resp,err := http.Get(newURL)

	//show response code
	fmt.Println(resp.StatusCode)

	if err!= nil{
		fmt.Println(err)
	}

	//Close connection at the end
	defer resp.Body.Close()

	//read response
	site,er := ioutil.ReadAll(resp.Body)
	if er!= nil{
		fmt.Println(er)
	}


	//display 
	fmt.Printf(string(site))

}
