package main

import (
  "log"
  "os"

  "github.com/spf13/viper"
)

func readConfig() {
  v := viper.New()
  v.SetConfigName("bridge")
  v.SetConfigType("yaml")
  v.AddConfigPath("$HOME/.bridge")
  v.AddConfigPath("./config/bridge")

  err := viper.ReadInConfig()
  if err != nil {
    log.Println("Could not read config file")
    log.Fatal(err)
  }

  //token := v.Get("apitoken")
  //return token
}

func main() {
  // update this to read a config file
  if len(os.Args) != 2 {
    log.Println("Missing token string arugment")
    log.Fatal("Start server: bridge tokenstring")
  }

  //token := readConfig()
  token := os.Args[1]
  getStatus(token)
}
