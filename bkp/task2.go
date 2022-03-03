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
    tokens_balance := make(map[string]string)
    if err != nil {
        log.Fatal(err)
    }

    r := gin.Default()

    res := getTokens()
    i := 0
    //fmt.Println(res)
    r.GET("/balance/:wallet_address", func(c *gin.Context){
      for _, token_address := range res {
        if i < 20 {
          tokAdd := token_address
          //fmt.Println(tokAdd)
          walAdd := c.Param("wallet_address")
          tokenAddress := common.HexToAddress(tokAdd)
          instance, err := token.NewToken(tokenAddress, client)
          if err != nil {
              log.Fatal(err)
          }

          address := common.HexToAddress(walAdd)
          bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
          if err != nil {
              log.Fatal(err)
          }
          name, err := instance.Name(&bind.CallOpts{})
          if err != nil {
              log.Fatal(err)
          }
          //fmt.Println(bal)
          
          value := bal.String()
          if value > "0" {
              tokens_balance[name] = value
              //fmt.Println(tokens_balance)
           } 
         i++
        }
      }
      jsonStr, _ := json.Marshal(tokens_balance)

      //fmt.Println(string(jsonStr))
      c.JSON(200, string(jsonStr))
    })
        

    r.Run()
    

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
            //fmt.Println(val.(string))
            a = append(a, val.(string))
            //a = val.(string)
        }
      }
    }
    return a
}
