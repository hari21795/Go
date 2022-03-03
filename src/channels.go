// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import (
  "fmt"
  "sync"
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
var wg sync.WaitGroup
func main() {
	tokens_balances := make(chan string, 500)
	walAdd := "0xfC43f5F9dd45258b3AFf31Bdbe6561D97e8B71de"
	address := common.HexToAddress(walAdd)
	res := getTokens()
	if err != nil {
	      log.Fatal(err)
	}
	
	for i := 1; i <= 500; i++ {
   wg.Add(1)
	 go tokenBalance(tokens_balances, res[i], address, client) 
  }
	
  wg.Wait()
  close(tokens_balances)
	
  for item := range tokens_balances {
      fmt.Println(item)
  }
}
func tokenBalance(c chan string, tokAdd string, address common.Address, client bind.ContractBackend) {
    defer wg.Done()
    tokenAddress := common.HexToAddress(tokAdd)
    instance, err := token.NewToken(tokenAddress, client)
    if err != nil {
        log.Fatal(err)
    }
    bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
    if err != nil {
        log.Fatal(err)
    }
    value := bal.String()
    if value > "0" {
        if err != nil {
            log.Fatal(err)
        }
        c <- value
    }
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