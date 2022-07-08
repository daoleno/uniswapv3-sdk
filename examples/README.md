# Examples
[quoter](./quoter/main.go) - get the swapped amount from chain.   
[swap](./swap/main.go) - swap two tokens on chain.      
[liquidity](liquidity/main.go) - shows how to mint/add/remove/burn a liquidity position.   

## Usage 
If you want to see the code running in real environment, set you private key to environment variable. The variable `MY_PRIVATE_KEY` will be get by each `main` function.
```bash
MY_PRIVATE_KEY=""
```

Replace `helper.TryTx` to `helper.SendTx` in each example case. 
```go
//try go send a transaction, it try to estimate gas price.
tx, err := helper.TryTX(client, common.HexToAddress(helper.ContractV3SwapRouterV1),
		swapValue, params.Calldata, wallet)
    
//send a transaction to chain, it will cost your money.
tx, err := helper.SendTX(client, common.HexToAddress(helper.ContractV3SwapRouterV1),
		swapValue, params.Calldata, wallet)
```
If you just want to check the paramerters are passed correctly, we recommend you use `TryTx`.

We use Polygon to test our code(it so cheap), you can set for your own.
```
client, err := ethclient.Dial(helper.PolygonRPC)
```

## Run
```
go run examples/quoter/*.go
```