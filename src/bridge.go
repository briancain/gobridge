package main

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
  "log"
  "net/http"
  "reflect"
  "github.com/dghubble/go-twitter/twitter"
)

const bridgeAPI string = "https://api.multco.us"

type Bridge struct {
  name string
  isUp bool
}

func toString(bridges map[string]bool) {
  for key,val := range bridges {
    log.Printf("%s status: %t", key, val)
  }
}

func readBridgeData(data []byte) map[string]bool {
  var bridges []interface{}
  if err := json.Unmarshal(data, &bridges); err != nil {
    log.Fatal(err)
  }

  var bridgeMap = make(map[string]bool)

  for _, bridge := range bridges {
    myBridge,_ := bridge.(map[string]interface{})

    name := myBridge["name"].(string)
    isUp := myBridge["isUp"].(bool)
    bridgeMap[name] = isUp
  }

  //toString(bridgeMap)
  return bridgeMap
}

func checkState(oldBridges map[string]bool, newBridges map[string]bool, twitterClient *twitter.Client) map[string]bool {
  if !reflect.DeepEqual(oldBridges, newBridges) {
    for bridge,status := range newBridges {
      if oldBridges[bridge] != status {
        log.Printf("Bridge %s has changed to %t", bridge, status)
        var returnString = ""
        if status == true {
          returnString = fmt.Sprintf("Bridge %s has been raised", bridge)
        } else if status == false {
          returnString = fmt.Sprintf("Bridge %s has been lowered", bridge)
        } else {
          log.Println("Could not determine bridge status")
          return oldBridges
        }

        log.Println("Sending tweet: ", returnString)
        tweet, resp, err := twitterClient.Statuses.Update(returnString, nil)

        if err != nil {
          log.Println("Something went wrong making a twitter request")
          log.Println(tweet)
          log.Println(resp)
          log.Println(err)
        }
      }
    }
  } else {
    log.Println("No change in bridge statuses")
  }

  return newBridges
}

func getStatus(bridges map[string]bool, token string, twitterClient *twitter.Client) (map[string]bool, error) {
  var allBridgeURL = bridgeAPI + "/bridges?access_token=" + token
  resp, err := http.Get(allBridgeURL)

  if err != nil {
    log.Printf("Failed to make API request for %s", allBridgeURL)
    return nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Println("Failed to parse response")
    return nil, err
  }

  var bridgesReturn = readBridgeData(body)
  var newState = checkState(bridges, bridgesReturn, twitterClient)

  return newState, nil
}
