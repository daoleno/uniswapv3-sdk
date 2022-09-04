package periphery

import (
	_ "embed"
	"math/big"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const INCENTIVE_KEY_ABI = "tuple(address rewardToken, address pool, uint256 startTime, uint256 endTime, address refundee)"

//go:embed contracts/UniswapV3Staker.sol/UniswapV3Staker.json
var stakerABI []byte

type FullWithdrawOptions struct {
	ClaimOptions
	WithdrawOptions
}

// Represents a unique staking program.
type IncentiveKey struct {
	RewardToken *core.Token    // The token rewarded for participating in the staking program.
	Pool        *entities.Pool // The pool that the staked positions must provide in.
	StartTime   *big.Int       // The time when the incentive program begins.
	EndTime     *big.Int       // The time that the incentive program ends.
	Refundee    common.Address // The address which receives any remaining reward tokens at `endTime`.
}

type IncentiveKeyParams struct {
	RewardToken common.Address
	Pool        common.Address
	StartTime   *big.Int
	EndTime     *big.Int
	Refundee    common.Address
}

// Options to specify when claiming rewards.
type ClaimOptions struct {
	TokenID   *big.Int       // The id of the NFT
	Recipient common.Address // Address to send rewards to.
	Amount    *big.Int       // The amount of `rewardToken` to claim. 0 claims all.
}

// Options to specify when withdrawing a position.
type WithdrawOptions struct {
	Owner common.Address // Set when withdrawing. The position will be sent to `owner` on withdraw.
	Data  []byte         // Set when withdrawing. `data` is passed to `safeTransferFrom` when transferring the position from contract back to owner.
}

/**
*  To claim rewards, must unstake and then claim.
* @param incentiveKey The unique identifier of a staking program.
* @param options Options for producing the calldata to claim. Can't claim unless you unstake.
* @returns The calldatas for 'unstakeToken' and 'claimReward'.
 */
func EncodeClaim(incentiveKey *IncentiveKey, options *ClaimOptions) ([][]byte, error) {
	var calldatas [][]byte

	abi := GetABI(stakerABI)
	params, err := encodeIncentiveKey(incentiveKey)
	if err != nil {
		return nil, err
	}
	calldata, err := abi.Pack("unstakeToken", params, options.TokenID)
	if err != nil {
		return nil, err
	}
	calldatas = append(calldatas, calldata)

	amount := big.NewInt(0)
	if options.Amount != nil {
		amount = options.Amount
	}
	calldata, err = abi.Pack("claimReward", incentiveKey.RewardToken.Address, options.Recipient, amount)
	if err != nil {
		return nil, err
	}
	calldatas = append(calldatas, calldata)

	return calldatas, nil
}

/*
*

	*
	* Note:  A `tokenId` can be staked in many programs but to claim rewards and continue the program you must unstake, claim, and then restake.
	* @param incentiveKeys An IncentiveKey or array of IncentiveKeys that `tokenId` is staked in.
	* Input an array of IncentiveKeys to claim rewards for each program.
	* @param options ClaimOptions to specify tokenId, recipient, and amount wanting to collect.
	* Note that you can only specify one amount and one recipient across the various programs if you are collecting from multiple programs at once.
	* @returns
*/
func CollectRewards(incentiveKeys []*IncentiveKey, options *ClaimOptions) (*utils.MethodParameters, error) {
	abi := GetABI(stakerABI)
	var calldatas [][]byte
	for _, incentiveKey := range incentiveKeys {
		// unstakes and claims for the unique program
		datas, err := EncodeClaim(incentiveKey, options)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, datas...)

		// re-stakes the position for the unique program
		params, err := encodeIncentiveKey(incentiveKey)
		if err != nil {
			return nil, err
		}
		calldata, err := abi.Pack("stakeToken", params, options.TokenID)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, calldata)
	}
	multiCalldata, err := EncodeMulticall(calldatas)
	if err != nil {
		return nil, err
	}

	return &utils.MethodParameters{
		Calldata: multiCalldata,
		Value:    big.NewInt(0),
	}, nil
}

/*
*
	*
	* @param incentiveKeys A list of incentiveKeys to unstake from. Should include all incentiveKeys (unique staking programs) that `options.tokenId` is staked in.
	* @param withdrawOptions Options for producing claim calldata and withdraw calldata. Can't withdraw without unstaking all programs for `tokenId`.
	* @returns Calldata for unstaking, claiming, and withdrawing.
*/
func WithdrawToken(incentiveKeys []*IncentiveKey, withdrawOptions *FullWithdrawOptions) (*utils.MethodParameters, error) {
	abi := GetABI(stakerABI)
	var calldatas [][]byte

	claimOptions := &ClaimOptions{
		TokenID:   withdrawOptions.TokenID,
		Recipient: withdrawOptions.Recipient,
		Amount:    withdrawOptions.Amount,
	}
	for _, incentiveKey := range incentiveKeys {
		datas, err := EncodeClaim(incentiveKey, claimOptions)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, datas...)
	}

	withdrawCalldata, err := abi.Pack("withdrawToken", withdrawOptions.TokenID, withdrawOptions.Owner, withdrawOptions.Data)
	if err != nil {
		return nil, err
	}

	calldatas = append(calldatas, withdrawCalldata)
	multiCalldata, err := EncodeMulticall(calldatas)
	if err != nil {
		return nil, err
	}
	return &utils.MethodParameters{
		Calldata: multiCalldata,
		Value:    big.NewInt(0),
	}, nil
}

/*
*

	*
	* @param incentiveKeys A single IncentiveKey or array of IncentiveKeys to be encoded and used in the data parameter in `safeTransferFrom`
	* @returns An IncentiveKey as a string
	*
*/
func EncodeDeposit(incentiveKeys []*IncentiveKey) ([]byte, error) {
	var data []byte
	var err error
	if len(incentiveKeys) > 1 {
		var keys []IncentiveKeyParams
		for _, incentiveKey := range incentiveKeys {
			params, err := encodeIncentiveKey(incentiveKey)
			if err != nil {
				return nil, err
			}
			keys = append(keys, *params)
		}
		// tuple(address rewardToken, address pool, uint256 startTime, uint256 endTime, address refundee)[]
		tupleArrayTy, _ := abi.NewType("tuple[]", "IncentiveKeyParams[]", []abi.ArgumentMarshaling{
			{Name: "rewardToken", Type: "address"},
			{Name: "pool", Type: "address"},
			{Name: "startTime", Type: "uint256"},
			{Name: "endTime", Type: "uint256"},
			{Name: "refundee", Type: "address"},
		})
		args := abi.Arguments{
			{Name: "keys", Type: tupleArrayTy},
		}
		data, err = args.Pack(keys)
		if err != nil {
			return nil, err
		}
	} else {
		params, err := encodeIncentiveKey(incentiveKeys[0])
		if err != nil {
			return nil, err
		}
		// tuple(address rewardToken, address pool, uint256 startTime, uint256 endTime, address refundee)
		tupleTy, _ := abi.NewType("tuple", "tuple", []abi.ArgumentMarshaling{
			{Name: "rewardToken", Type: "address"},
			{Name: "pool", Type: "address"},
			{Name: "startTime", Type: "uint256"},
			{Name: "endTime", Type: "uint256"},
			{Name: "refundee", Type: "address"},
		})
		args := abi.Arguments{
			{Name: "key", Type: tupleTy},
		}
		data, err = args.Pack(params)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

/*
*

	*
	* @param incentiveKey An `IncentiveKey` which represents a unique staking program.
	* @returns An encoded IncentiveKey to be read by ethers
	*
*/
func encodeIncentiveKey(incentiveKey *IncentiveKey) (*IncentiveKeyParams, error) {
	pool := incentiveKey.Pool
	addr, err := entities.GetAddress(pool.Token0, pool.Token1, pool.Fee, "")
	if err != nil {
		return nil, err
	}

	return &IncentiveKeyParams{
		RewardToken: incentiveKey.RewardToken.Address,
		Pool:        addr,
		StartTime:   incentiveKey.StartTime,
		EndTime:     incentiveKey.EndTime,
		Refundee:    incentiveKey.Refundee,
	}, nil

}
