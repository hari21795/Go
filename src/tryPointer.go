// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import (
  "fmt"
  "sync"
  "log"
  "math"
  "math/big"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/googleapis/gax-go/tokens"
  "io/ioutil"
  "net/http"
  "os"
  "github.com/m7shapan/njson"
)
type TokenList struct {
    Tokens []string `njson:"tokens.#.address"`
}
var wg sync.WaitGroup
func main() {
  res := getTokens()
	tokens_balances := make(chan map[string]*big.Float, len(res)-1)
  walAdd := "0xfC43f5F9dd45258b3AFf31Bdbe6561D97e8B71de"
  address := common.HexToAddress(walAdd)
  var name = ""
  var nameptr *string = &name
	if err != nil {
	      log.Fatal(err)
	}
	
	for i := 1; i <= (len(res)-1); i++ {
   wg.Add(1)
	 go tokenBalance(tokens_balances, res[i], address, client, nameptr) 
  }
	
  wg.Wait()
  close(tokens_balances)
	
  // for item := range tokens_balances {
  //     fmt.Println(item)
  // }
}
func tokenBalance(c chan map[string]*big.Float, tokAdd string, address common.Address, client bind.ContractBackend, name *string ) {
    defer wg.Done()
    token_balance := make(map[string]*big.Float)
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
        *name, err = instance.Name(&bind.CallOpts{})
        fmt.Println(name)
        decimals, err := instance.Decimals(&bind.CallOpts{})
        if err != nil {
            log.Fatal(err)
        }
        fbal := new(big.Float)
        fbal.SetString(bal.String())
        value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
        if err != nil {
            log.Fatal(err)
        }
        token_balance[*name] = value
        c <- token_balance
    }
}
func getTokens()  []string {
    //response, err := http.Get("https://uniswap.mycryptoapi.com/")
    response, err := http.Get("https://wispy-bird-88a7.uniswap.workers.dev/?url=http://tokens.1inch.eth.link")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    result := TokenList{}
 
    // Unmarshal or Decode the JSON to the interface.
    njson.Unmarshal([]byte(responseData), &result)
    return result.Tokens
}