package main

import (
    "fmt"
    "log"
    //"math"
    //"math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/gin-gonic/gin"
    "github.com/googleapis/gax-go/tokens"
    //token "./erc20" // for demo

    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
)


func main() {
  //channels
    tokens_balance := make(map[string]string)
    tokens_list := make(chan []string)
    go SendTokens(tokens_list)
    if err != nil {
        log.Fatal(err)
    }
    
    r := gin.Default()

    
    i := 0
    r.GET("/balance/:wallet_address", func(c *gin.Context){
      walAdd := c.Param("wallet_address")
      address := common.HexToAddress(walAdd)
      res := <-tokens_list
      for _, token_address := range res {
        if i < 20 {
          tokAdd := token_address
          //fmt.Println(tokAdd)
          tokenAddress := common.HexToAddress(tokAdd)
          instance, err := token.NewToken(tokenAddress, client)
          if err != nil {
              log.Fatal(err)
          }
          bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
          if err != nil {
              log.Fatal(err)
          }
          //tokens_balance := <-tokens_balance
          value := bal.String()
          if value > "0" {
             name, err := instance.Name(&bind.CallOpts{})
              if err != nil {
                  log.Fatal(err)
              }
              tokens_balance[name] = value
           } 
         i++
        }
      }
      //fmt.Println(tokens_balance)
      jsonStr, _ := json.Marshal(tokens_balance)
      c.JSON(200, string(jsonStr))
    })
    r.Run()
}


func SendTokens(c chan []string){
  c <- getTokens()
}

func getTokens() []string  {
    response, err := http.Get("https://wispy-bird-88a7.uniswap.workers.dev/?url=http://tokens.1inch.eth.link")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    result := make(map[string]interface{})
 
    // Unmarshal or Decode the JSON to the interface.
    json.Unmarshal([]byte(responseData), &result)
    tokens := result["tokens"].([]interface{})
    a := []string{}

    for _, value := range tokens {
      // Each value is an interface{} type, that is type asserted as a string
      //
      for k, val := range value.(map[string]interface{}){
        if k == "address" {
            //fmt.Println(val.(string))
            a = append(a, val.(string))
            //a = val.(string)
        }
      }
    }
    return a
}
