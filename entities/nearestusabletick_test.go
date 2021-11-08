package entities

import (
	"testing"

	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestNearestUsableTick(t *testing.T) {
	assert.Panics(t, func() { NearestUsableTick(1, 0) }, "panics if tickSpacing is 0")
	assert.Panics(t, func() { NearestUsableTick(1, -5) }, "panics if tickSpacing is negative")
	assert.Panics(t, func() { NearestUsableTick(utils.MaxTick+1, 1) }, "panics if tick is greater than MaxTick")
	assert.Panics(t, func() { NearestUsableTick(utils.MinTick-1, 1) }, "panics if tick is smaller than MinTick")

	type args struct {
		ticks       int
		tickSpacing int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "rounds at positive half", args: args{ticks: 5, tickSpacing: 10}, want: 10},
		{name: "rounds down below positive half", args: args{ticks: 4, tickSpacing: 10}, want: 0},
		{name: "rounds up for negative half 0", args: args{ticks: -5, tickSpacing: 10}, want: 0},
		{name: "rounds up for negative half 1", args: args{ticks: -6, tickSpacing: 10}, want: -10},
		{name: "cannot round past MinTick", args: args{ticks: utils.MinTick, tickSpacing: utils.MaxTick/2 + 100}, want: -(utils.MaxTick/2 + 100)},
		{name: "cannot round past MaxTick", args: args{ticks: utils.MaxTick, tickSpacing: utils.MaxTick/2 + 100}, want: utils.MaxTick/2 + 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NearestUsableTick(tt.args.ticks, tt.args.tickSpacing); got != tt.want {
				t.Errorf("NearestUsableTick() = %v, want %v", got, tt.want)
			}
		})
	}
}
