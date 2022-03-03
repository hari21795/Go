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
type Token struct {
    name string
    balance  *big.Float
}
var wg sync.WaitGroup
func main() {
  res := getTokens()
  c := make(chan Token, (len(res)/2))
  walAdd := "0xfC43f5F9dd45258b3AFf31Bdbe6561D97e8B71de"
  address := common.HexToAddress(walAdd)
 
  if err != nil {
        log.Fatal(err)
  }

  for i := 0; i <= (len(res)-5); i=i+6  {
   j := i + 5
   wg.Add(1)
   go tokenBalance(c, res[i:j], address, client)
  }

  wg.Wait()
  close(c)
  for item := range c {
      fmt.Println(item.name)
      fmt.Println(item.balance)
  }
}
func tokenBalance(c chan Token, tokAdd []string, address common.Address, client bind.ContractBackend) {
    defer wg.Done()
    for k := 0; k <= (len(tokAdd)-1); k++ {
      tokenAddress := common.HexToAddress(tokAdd[k])
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
          name, err := instance.Name(&bind.CallOpts{})
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
          //fmt.Println(tokAdd)
          c <- Token{name, value}
      }
    }  
}
func getTokens()  []string {
    //response, err := http.Get("https://wispy-bird-88a7.uniswap.workers.dev/?url=http://t2crtokens.eth.link")
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