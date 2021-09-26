package uniswapv3

//go:generate docker run --rm -v ${PWD}:/root ethereum/client-go:alltools-latest abigen --abi root/QuoterABI.json --pkg uniswapv3  --type quoter --out root/quoter.go
