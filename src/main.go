package main

import (
    "fmt"
    "log"
    //"math"
    //"math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    //"github.com/gin-gonic/gin"
    "github.com/googleapis/gax-go/tokens"
    //token "./erc20" // for demo

    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
    "sync"
    //"time"
)


func main() {
  //channels
    tokens_balance := make(map[string]string)
    
    if err != nil {
        log.Fatal(err)
    }
    
    //r := gin.Default()

    var wg sync.WaitGroup
    //r.GET("/balance/:wallet_address", func(c *gin.Context){
      walAdd := "0xfC43f5F9dd45258b3AFf31Bdbe6561D97e8B71de"
      address := common.HexToAddress(walAdd)
      res := getTokens()
      for i := 0; i < 100; i++ {
        //if i < 20 {
          wg.Add(1)
          go func(){
            defer wg.Done()
            tokAdd := res[i]
            //fmt.Println(tokAdd)
             tokenBalance(address, tokAdd, tokens_balance, client)
             fmt.Printf("Worker %d done\n", res[i])
          }()
          wg.Wait()  
        // i++
        //}
      }
      //fmt.Println(tokens_balance)
      //jsonStr, _ := json.Marshal(tokens_balance)
      //c.JSON(200, string(jsonStr))
    //})
    //r.Run()
}
 

func tokenBalance(address common.Address, tokAdd string, tokens_balance map[string]string, client bind.ContractBackend){
    fmt.Printf("Worker %d starting\n", tokAdd)
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
    fmt.Printf("Worker end", tokAdd)
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
