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

func main() {
  log.Println("Starting gobridge server")
  token := readConfig()

  log.Println("Getting bridge status")
  getStatus(token)
}
