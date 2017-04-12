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
  actualCount int
  avgUpTime int
  id int
  name string
  isUp bool
  scheduledCount int
  totalUpTime int
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
    //if b, ok := bridge.(map[string]interface{}); ok {
    //  for k,v := range b {
    //    fmt.Println(k, " == ", v)
    //  }
    //  fmt.Println()
    //} else {
    //  fmt.Println("Bridge not a map[string]interface{}: ", bridge)
    //}
  }
}

func getStatus(token string){
  var allBridgeURL = bridgeAPI + "/bridges?access_token=" + token
  resp, err := http.Get(allBridgeURL)

  if err != nil {
    fmt.Println("We h*cked up")
    log.Fatal(err)
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("We h*cked up")
    log.Fatal(err)
  }

  readBridgeData(body)
}
