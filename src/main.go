package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

// start bridge package

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
  var bridges interface{}
  if err := json.Unmarshal(data, &bridges); err != nil {
    log.Fatal(err)
  }
  fmt.Println(bridges)
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

// end bridge package

func main() {
  // update this to read a config file
  if len(os.Args) != 2 {
    log.Fatal("Missing token string arugment")
    log.Fatal("Usage: go run src/main.go tokenstring")
    os.Exit(1)
  }

  var token = os.Args[1]
  getStatus(token)
}
