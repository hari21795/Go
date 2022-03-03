
   r.GET("/balance/:token_address/:wallet_address", func(c *gin.Context){
    tokAdd := c.Param("token_address")
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

    symbol, err := instance.Symbol(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    decimals, err := instance.Decimals(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

   

    fbal := new(big.Float)
    fbal.SetString(bal.String())
    value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

    fmt.Printf("balance: %f", value) 
        c.JSON(200, gin.H{
            "name" : name,
            "symbol" : symbol,
            "decimals" : decimals,
            "wei" : bal,
            "balance": value,
        })
    })
