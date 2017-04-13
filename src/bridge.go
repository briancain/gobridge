package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

const bridgeAPI string = "https://api.multco.us"

type Bridge struct {
  name string
  isUp bool
}

func readBridgeData(data []byte) {
  var bridges []interface{}
  if err := json.Unmarshal(data, &bridges); err != nil {
    log.Fatal(err)
  }
  fmt.Println(bridges)
  for _, bridge := range bridges {
    //fmt.Println("bridge:", bridge)
    myBridge,_ := bridge.(map[string]interface{})
    fmt.Println(myBridge["name"], ":", myBridge["isUp"])
  }
}

func getStatus(token string){
  var allBridgeURL = bridgeAPI + "/bridges?access_token=" + token
  resp, err := http.Get(allBridgeURL)

  if err != nil {
    log.Printf("Failed to make API request for %s", allBridgeURL)
    log.Println(err)
    return
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Println("Failed to parse response")
    log.Println(err)
  }

  readBridgeData(body)
}
