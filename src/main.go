package main

import (
  "log"
  "fmt"
  "strings"
  "net/http"
  "time"

  "github.com/spf13/viper"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)

func readConfig() *viper.Viper {
  log.Println("Reading config")
  v := viper.New()
  v.SetConfigName("bridge")
  v.AddConfigPath("config")

  err := v.ReadInConfig()
  if err != nil {
    log.Println("Could not read config file")
    log.Fatal(err)
  }

  return v
}

func makeRequest(token string, initBridge map[string]bool, twitterClient *twitter.Client) map[string]bool {
  var resp, err = getStatus(initBridge, token, twitterClient)
  if err != nil {
    log.Fatal(err)
  } else {
    bridges := resp
    log.Println(bridges)
    return bridges
  }
  return nil
}

func startServer(token string, twitterClient *twitter.Client) {
  log.Println("Starting gobridge server")
  log.Println("Getting bridge status")

  var interval = 30 * time.Second

  bridges := map[string]bool {
    "hawthorne": false,
    "morrison": false,
    "burnside": false,
    "broadway": false,
    "I-5": false,
  }

  for {
    bridges = makeRequest(token, bridges, twitterClient)
    time.Sleep(interval)
  }
}

func sendVersion(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()  // parse arguments, you have to call this by yourself
  fmt.Println(r.Form)  // print form information in server side
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
      fmt.Println("key:", k)
      fmt.Println("val:", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "Hello gobridge!") // send data to client side
}

func setupTwitterClient(config *viper.Viper) *twitter.Client {
  twitterConsumerKey := config.GetString("twitter.consumerkey")
  twitterConsumerSecret := config.GetString("twitter.consumersecret")
  twitterAccessToken := config.GetString("twitter.accesstoken")
  twitterAccessSecret := config.GetString("twitter.accesstokensecret")

  twitterConfig := oauth1.NewConfig(twitterConsumerKey, twitterConsumerSecret)
  twitterToken := oauth1.NewToken(twitterAccessToken, twitterAccessSecret)

  httpClient := twitterConfig.Client(oauth1.NoContext, twitterToken)
  client := twitter.NewClient(httpClient)

  return client
}

func main() {
  viperConfig := readConfig()

  token := viperConfig.GetString("apikey")
  twitterClient := setupTwitterClient(viperConfig)

  go startServer(token, twitterClient)

  http.HandleFunc("/", sendVersion)
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
