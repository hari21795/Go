package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
)

// type TokenList struct {
//     Token Token `json:"Token"`
// }
// type Token struct {
//   Address  string `json:"address"`
// }
func main() {


    response, err := http.Get("https://wispy-bird-88a7.uniswap.workers.dev/?url=http://tokens.1inch.eth.link")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(string(responseData))

    result := make(map[string]interface{})
 
    // Unmarshal or Decode the JSON to the interface.
    json.Unmarshal([]byte(responseData), &result)
    tokens := result["tokens"].([]interface{})
    //fmt.Println(tokens)
    a := []string{}

    for _, value := range tokens {
      // Each value is an interface{} type, that is type asserted as a string
      //
      for k, val := range value.(map[string]interface{}){
        if k == "address" {
            fmt.Println(val.(string))
            a = append(a, val.(string))
            //a = val.(string)
        }
      }
    }
    for _, address := range a{
        fmt.Println(address)
    }
   // fmt.Println(a)

}