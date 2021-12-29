package utils

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

func token(sortOrder, decimals, chainID uint) *entities.Token {
	if sortOrder > 9 || sortOrder%1 != 0 {
		panic("invalid sort order")
	}
	address := common.HexToAddress("0x" + strings.Repeat(fmt.Sprint(sortOrder), 40))
	return entities.NewToken(chainID, address, decimals, fmt.Sprintf("T%d", sortOrder), fmt.Sprintf("token%d", sortOrder))
}

var (
	token0           = token(0, 18, 1)
	token1           = token(1, 18, 1)
	token2_6decimals = token(2, 6, 1)
)

func TestTickToPrice(t *testing.T) {
	type args struct {
		baseToken  *entities.Token
		quoteToken *entities.Token
		tick       int
	}
	tests := []struct {
		name            string
		args            args
		wantSignificant string
	}{
		{"1800 t0/1 t1", args{token1, token0, -74959}, "1800"},
		{"1 t1/1800 t0", args{token0, token1, -74959}, "0.00055556"},
		{"1800 t1/1 t0", args{token0, token1, 74959}, "1800"},
		{"1 t0/1800 t1", args{token1, token0, 74959}, "0.00055556"},

		// 12 decimal difference
		{"1.01 t2/1 t0", args{token1, token2_6decimals, -276225}, "1.01"},
		{"1 t0/1.01 t2", args{token2_6decimals, token0, -276225}, "0.99015"},
		{"1 t2/1.01 t0", args{token0, token2_6decimals, -276423}, "0.99015"},
		{"1.01 t0/1 t2", args{token2_6decimals, token0, -276423}, "1.0099"},
		{"1.01 t2/1 t0", args{token0, token2_6decimals, -276225}, "1.01"},
		{"1 t0/1.01 t2", args{token2_6decimals, token0, -276225}, "0.99015"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TickToPrice(tt.args.baseToken, tt.args.quoteToken, tt.args.tick)
			if err != nil {
				t.Errorf("TickToPrice() error = %v", err)
				return
			}
			if got.ToSignificant(5) != tt.wantSignificant {
				t.Errorf("TickToPrice() = %v, want %v", got, tt.wantSignificant)
			}
		})
	}
}

func TestPriceToClosestTick(t *testing.T) {
	tickToPriceNoError := func(baseToken *entities.Token, quoteToken *entities.Token, tick int) *entities.Price {
		p, err := TickToPrice(baseToken, quoteToken, tick)
		if err != nil {
			panic(err)
		}
		return p
	}
	type args struct {
		price      *entities.Price
		baseToken  *entities.Token
		quoteToken *entities.Token
	}
	B100e18 := decimal.NewFromBigInt(big.NewInt(100), 18).BigInt()

	tests := []struct {
		name     string
		args     args
		wantTick int
	}{
		{"1800 t0/1 t1", args{entities.NewPrice(token1.Currency, token0.Currency, big.NewInt(1), big.NewInt(1800)), token1, token0}, -74960},
		{"1 t1/1800 t0", args{entities.NewPrice(token0.Currency, token1.Currency, big.NewInt(1800), big.NewInt(1)), token0, token1}, -74960},
		{"1.01 t2/1 t0", args{entities.NewPrice(token0.Currency, token2_6decimals.Currency, B100e18, big.NewInt(101e6)), token0, token2_6decimals}, -276225},
		{"1 t0/1.01 t2", args{entities.NewPrice(token2_6decimals.Currency, token0.Currency, big.NewInt(101e6), B100e18), token2_6decimals, token0}, -276225},

		// reciprocal with tickToPrice
		{"1800 t0/1 t1", args{tickToPriceNoError(token1, token0, -74960), token1, token0}, -74960},
		{"1 t0/1800 t1", args{tickToPriceNoError(token1, token0, 74960), token1, token0}, 74960},
		{"1 t1/1800 t0", args{tickToPriceNoError(token0, token1, -74960), token0, token1}, -74960},
		{"1800 t1/1 t0", args{tickToPriceNoError(token0, token1, 74960), token0, token1}, 74960},
		{"1.01 t2/1 t0", args{tickToPriceNoError(token0, token2_6decimals, -276225), token0, token2_6decimals}, -276225},
		{"1 t0/1.01 t2", args{tickToPriceNoError(token2_6decimals, token0, -276225), token2_6decimals, token0}, -276225},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PriceToClosestTick(tt.args.price, tt.args.baseToken, tt.args.quoteToken)
			if err != nil {
				t.Errorf("PriceToClosestTick() error = %v", err)
				return
			}
			if got != tt.wantTick {
				t.Errorf("PriceToClosestTick() = %v, want %v", got, tt.wantTick)
			}
		})
	}
}
