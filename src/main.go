package main

import (
  "log"
  "github.com/spf13/viper"
)

func readConfig() string{
  log.Println("Reading config")
  v := viper.New()
  v.SetConfigName("bridge")
  v.AddConfigPath("config")

  err := v.ReadInConfig()
  if err != nil {
    log.Println("Could not read config file")
    log.Fatal(err)
  }

  token := v.GetString("apikey")
  return token
}

func startServer(token string) {
  log.Println("Starting gobridge server")
  log.Println("Getting bridge status")

  // init data
  bridges := map[string]bool {
    "hawthorne": false,
    "morrison": false,
    "burnside": false,
    "broadway": false,
    "I-5": false,
  }

  var resp, err = getStatus(bridges, token)
  if err != nil {
    log.Fatal(err)
  } else {
    bridges = resp
    log.Println(bridges)
  }
}

func main() {
  log.Println("Starting gobridge server")
  token := readConfig()

  startServer(token)
}
