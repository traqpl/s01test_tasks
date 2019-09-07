package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)
  
func main() {
	// elasticsearchurl := os.Args[1]
	// searchstring := os.Args[2]
	elasticsearchurl := "localhost:9200/index/message/_count?q=error:Handbill+not+printed&pretty=true"
	searchstring := ""

	var httpCustomClient = &http.Client{
		Timeout: time.Second * 60,
	}
	//http requests
	req, err := http.NewRequest("GET", elasticsearchurl, bytes.NewBufferString(searchstring))
	if err != nil {
		log.Printf("Building http request error: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)
	resp, err := httpCustomClient.Do(req)
	if err != nil {
		log.Printf("Http response error: %v", err)
		return
	}
	//fmt.Sprintf("%v", resp), err
}


//curl -XGET 'http://localhost:9200/index/message/_count?q=error:Handbill+not+printed&pretty=true'

// {
// 	"count" : 3,
// 	"_shards" : {
// 	  "total" : 1,
// 	  "successful" : 1,
// 	  "skipped" : 0,
// 	  "failed" : 0
// 	}
//   }