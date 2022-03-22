package periphery

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestEncodeRouteToPath(t *testing.T) {
	// packs them for exact input single hop
	p, _ := EncodeRouteToPath(route_0_1, false)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb80000000000000000000000000000000000000002", hexutil.Encode(p))

	// packs them correctly for exact output single hop
	p, _ = EncodeRouteToPath(route_0_1, true)
	assert.Equal(t, "0x0000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001", hexutil.Encode(p))

	// packs them correctly for multihop exact input
	p, _ = EncodeRouteToPath(route_0_1_2, false)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb800000000000000000000000000000000000000020001f40000000000000000000000000000000000000003", hexutil.Encode(p))

	// packs them correctly for multihop exact output
	p, _ = EncodeRouteToPath(route_0_1_2, true)
	assert.Equal(t, "0x00000000000000000000000000000000000000030001f40000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001", hexutil.Encode(p))

	// wraps ether input for exact input single hop
	p, _ = EncodeRouteToPath(route_weth_0, false)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000001", hexutil.Encode(p))

	// wraps ether input for exact output single hop
	p, _ = EncodeRouteToPath(route_weth_0, true)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", hexutil.Encode(p))

	// wraps ether input for exact input multihop
	p, _ = EncodeRouteToPath(route_weth_0_1, false)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000001000bb80000000000000000000000000000000000000002", hexutil.Encode(p))

	// wraps ether input for exact output multihop
	p, _ = EncodeRouteToPath(route_weth_0_1, true)
	assert.Equal(t, "0x0000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", hexutil.Encode(p))

	// wraps ether output for exact input single hop
	p, _ = EncodeRouteToPath(route_0_weth, false)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", hexutil.Encode(p))

	// wraps ether output for exact output single hop
	p, _ = EncodeRouteToPath(route_0_weth, true)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000001", hexutil.Encode(p))

	// wraps ether output for exact input multihop
	p, _ = EncodeRouteToPath(route_0_1_weth, false)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb80000000000000000000000000000000000000002000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", hexutil.Encode(p))

	// wraps ether output for exact output multihop
	p, _ = EncodeRouteToPath(route_0_1_weth, true)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001", hexutil.Encode(p))
}
