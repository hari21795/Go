package main

import (
  "fmt"
  "sync"
  //"time"
  "log"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/googleapis/gax-go/tokens"
  //"github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
  "os"
  "encoding/json"
)


func main() {
  var wg *sync.WaitGroup = new(sync.WaitGroup)
  walAdd := "0xfC43f5F9dd45258b3AFf31Bdbe6561D97e8B71de"
  address := common.HexToAddress(walAdd)
  tokens_balance := make(map[string]string)
  //r := gin.Default()
  if err != nil {
      log.Fatal(err)
  }

  res := getTokens()
  for i := 1; i <= 100; i++ {
    wg.Add(1)
 
    i := i
    go func() {
      //defer wg.Done()
      tokenBalance(wg, address, res[i], tokens_balance, client)
    }()
  }


  wg.Wait()

}
func tokenBalance(w *sync.WaitGroup, address common.Address, tokAdd string, tokens_balance map[string]string, client bind.ContractBackend){
   defer w.Done()
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
    fmt.Printf("Worker end", tokens_balance)
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