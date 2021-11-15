// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IStakingCloseAllocationRequest is an auto generated low-level Go binding around an user-defined struct.
type IStakingCloseAllocationRequest struct {
	AllocationID common.Address
	Poi          [32]byte
}

// IStakingDataAllocation is an auto generated low-level Go binding around an user-defined struct.
type IStakingDataAllocation struct {
	Indexer                     common.Address
	SubgraphDeploymentID        [32]byte
	Tokens                      *big.Int
	CreatedAtEpoch              *big.Int
	ClosedAtEpoch               *big.Int
	CollectedFees               *big.Int
	EffectiveAllocation         *big.Int
	AccRewardsPerAllocatedToken *big.Int
}

// IStakingDataDelegation is an auto generated low-level Go binding around an user-defined struct.
type IStakingDataDelegation struct {
	Shares            *big.Int
	TokensLocked      *big.Int
	TokensLockedUntil *big.Int
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveAllocation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"poi\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDelegator\",\"type\":\"bool\"}],\"name\":\"AllocationClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curationFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rebateFees\",\"type\":\"uint256\"}],\"name\":\"AllocationCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"AllocationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetHolder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AssetHolderUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"indexingRewardCut\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"queryFeeCut\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"cooldownBlocks\",\"type\":\"uint32\"}],\"name\":\"DelegationParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"param\",\"type\":\"string\"}],\"name\":\"ParameterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unclaimedAllocationsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegationFees\",\"type\":\"uint256\"}],\"name\":\"RebateClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"SetController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"SetRewardsDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SlasherUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StakeDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"StakeDelegatedLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"StakeDelegatedWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"StakeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIGraphProxy\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"acceptProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGraphProxy\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"acceptProxyAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addressCache\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"allocateFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allocations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closedAtEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardsPerAllocatedToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alphaDenominator\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alphaNumerator\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetHolders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"channelDisputeEpochs\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_restake\",\"type\":\"bool\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_allocationID\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_restake\",\"type\":\"bool\"}],\"name\":\"claimMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_poi\",\"type\":\"bytes32\"}],\"name\":\"closeAllocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"poi\",\"type\":\"bytes32\"}],\"internalType\":\"structIStaking.CloseAllocationRequest[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"}],\"name\":\"closeAllocationMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_closingAllocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_poi\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"closeAndAllocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curationPercentage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationParametersCooldown\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegationPools\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"cooldownBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"indexingRewardCut\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"queryFeeCut\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"updatedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationRatio\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationTaxPercentage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationUnbondingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"getAllocation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closedAtEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardsPerAllocatedToken\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingData.Allocation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"getAllocationState\",\"outputs\":[{\"internalType\":\"enumIStaking.AllocationState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"getDelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingData.Delegation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"}],\"name\":\"getIndexerCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"}],\"name\":\"getIndexerStakedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"}],\"name\":\"getSubgraphAllocatedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingData.Delegation\",\"name\":\"_delegation\",\"type\":\"tuple\"}],\"name\":\"getWithdraweableDelegatedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"}],\"name\":\"hasStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumIndexerStake\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_thawingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_protocolPercentage\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_curationPercentage\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_channelDisputeEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxAllocationEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_delegationUnbondingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_delegationRatio\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_rebateAlphaNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_rebateAlphaDenominator\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"isAllocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"isDelegator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAllocationEpochs\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumIndexerStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolPercentage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rebates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveAllocatedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"unclaimedAllocationsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"alphaNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"alphaDenominator\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetHolder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setAssetHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_channelDisputeEpochs\",\"type\":\"uint32\"}],\"name\":\"setChannelDisputeEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_percentage\",\"type\":\"uint32\"}],\"name\":\"setCurationPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_indexingRewardCut\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_queryFeeCut\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_cooldownBlocks\",\"type\":\"uint32\"}],\"name\":\"setDelegationParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_blocks\",\"type\":\"uint32\"}],\"name\":\"setDelegationParametersCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_delegationRatio\",\"type\":\"uint32\"}],\"name\":\"setDelegationRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_percentage\",\"type\":\"uint32\"}],\"name\":\"setDelegationTaxPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_delegationUnbondingPeriod\",\"type\":\"uint32\"}],\"name\":\"setDelegationUnbondingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maxAllocationEpochs\",\"type\":\"uint32\"}],\"name\":\"setMaxAllocationEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumIndexerStake\",\"type\":\"uint256\"}],\"name\":\"setMinimumIndexerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_percentage\",\"type\":\"uint32\"}],\"name\":\"setProtocolPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_alphaNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_alphaDenominator\",\"type\":\"uint32\"}],\"name\":\"setRebateRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"setRewardsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setSlasher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_thawingPeriod\",\"type\":\"uint32\"}],\"name\":\"setThawingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slashers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"stakeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokensStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensAllocated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLockedUntil\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subgraphAllocations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thawingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegateToIndexer\",\"type\":\"address\"}],\"name\":\"withdrawDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// AddressCache is a free data retrieval call binding the contract method 0xdc675a65.
//
// Solidity: function addressCache(bytes32 ) view returns(address)
func (_Staking *StakingCaller) AddressCache(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "addressCache", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressCache is a free data retrieval call binding the contract method 0xdc675a65.
//
// Solidity: function addressCache(bytes32 ) view returns(address)
func (_Staking *StakingSession) AddressCache(arg0 [32]byte) (common.Address, error) {
	return _Staking.Contract.AddressCache(&_Staking.CallOpts, arg0)
}

// AddressCache is a free data retrieval call binding the contract method 0xdc675a65.
//
// Solidity: function addressCache(bytes32 ) view returns(address)
func (_Staking *StakingCallerSession) AddressCache(arg0 [32]byte) (common.Address, error) {
	return _Staking.Contract.AddressCache(&_Staking.CallOpts, arg0)
}

// Allocations is a free data retrieval call binding the contract method 0x52a9039c.
//
// Solidity: function allocations(address ) view returns(address indexer, bytes32 subgraphDeploymentID, uint256 tokens, uint256 createdAtEpoch, uint256 closedAtEpoch, uint256 collectedFees, uint256 effectiveAllocation, uint256 accRewardsPerAllocatedToken)
func (_Staking *StakingCaller) Allocations(opts *bind.CallOpts, arg0 common.Address) (struct {
	Indexer                     common.Address
	SubgraphDeploymentID        [32]byte
	Tokens                      *big.Int
	CreatedAtEpoch              *big.Int
	ClosedAtEpoch               *big.Int
	CollectedFees               *big.Int
	EffectiveAllocation         *big.Int
	AccRewardsPerAllocatedToken *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "allocations", arg0)

	outstruct := new(struct {
		Indexer                     common.Address
		SubgraphDeploymentID        [32]byte
		Tokens                      *big.Int
		CreatedAtEpoch              *big.Int
		ClosedAtEpoch               *big.Int
		CollectedFees               *big.Int
		EffectiveAllocation         *big.Int
		AccRewardsPerAllocatedToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Indexer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SubgraphDeploymentID = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Tokens = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CreatedAtEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ClosedAtEpoch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CollectedFees = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EffectiveAllocation = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AccRewardsPerAllocatedToken = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Allocations is a free data retrieval call binding the contract method 0x52a9039c.
//
// Solidity: function allocations(address ) view returns(address indexer, bytes32 subgraphDeploymentID, uint256 tokens, uint256 createdAtEpoch, uint256 closedAtEpoch, uint256 collectedFees, uint256 effectiveAllocation, uint256 accRewardsPerAllocatedToken)
func (_Staking *StakingSession) Allocations(arg0 common.Address) (struct {
	Indexer                     common.Address
	SubgraphDeploymentID        [32]byte
	Tokens                      *big.Int
	CreatedAtEpoch              *big.Int
	ClosedAtEpoch               *big.Int
	CollectedFees               *big.Int
	EffectiveAllocation         *big.Int
	AccRewardsPerAllocatedToken *big.Int
}, error) {
	return _Staking.Contract.Allocations(&_Staking.CallOpts, arg0)
}

// Allocations is a free data retrieval call binding the contract method 0x52a9039c.
//
// Solidity: function allocations(address ) view returns(address indexer, bytes32 subgraphDeploymentID, uint256 tokens, uint256 createdAtEpoch, uint256 closedAtEpoch, uint256 collectedFees, uint256 effectiveAllocation, uint256 accRewardsPerAllocatedToken)
func (_Staking *StakingCallerSession) Allocations(arg0 common.Address) (struct {
	Indexer                     common.Address
	SubgraphDeploymentID        [32]byte
	Tokens                      *big.Int
	CreatedAtEpoch              *big.Int
	ClosedAtEpoch               *big.Int
	CollectedFees               *big.Int
	EffectiveAllocation         *big.Int
	AccRewardsPerAllocatedToken *big.Int
}, error) {
	return _Staking.Contract.Allocations(&_Staking.CallOpts, arg0)
}

// AlphaDenominator is a free data retrieval call binding the contract method 0xce853613.
//
// Solidity: function alphaDenominator() view returns(uint32)
func (_Staking *StakingCaller) AlphaDenominator(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "alphaDenominator")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AlphaDenominator is a free data retrieval call binding the contract method 0xce853613.
//
// Solidity: function alphaDenominator() view returns(uint32)
func (_Staking *StakingSession) AlphaDenominator() (uint32, error) {
	return _Staking.Contract.AlphaDenominator(&_Staking.CallOpts)
}

// AlphaDenominator is a free data retrieval call binding the contract method 0xce853613.
//
// Solidity: function alphaDenominator() view returns(uint32)
func (_Staking *StakingCallerSession) AlphaDenominator() (uint32, error) {
	return _Staking.Contract.AlphaDenominator(&_Staking.CallOpts)
}

// AlphaNumerator is a free data retrieval call binding the contract method 0x7ef82070.
//
// Solidity: function alphaNumerator() view returns(uint32)
func (_Staking *StakingCaller) AlphaNumerator(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "alphaNumerator")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AlphaNumerator is a free data retrieval call binding the contract method 0x7ef82070.
//
// Solidity: function alphaNumerator() view returns(uint32)
func (_Staking *StakingSession) AlphaNumerator() (uint32, error) {
	return _Staking.Contract.AlphaNumerator(&_Staking.CallOpts)
}

// AlphaNumerator is a free data retrieval call binding the contract method 0x7ef82070.
//
// Solidity: function alphaNumerator() view returns(uint32)
func (_Staking *StakingCallerSession) AlphaNumerator() (uint32, error) {
	return _Staking.Contract.AlphaNumerator(&_Staking.CallOpts)
}

// AssetHolders is a free data retrieval call binding the contract method 0x6b535d7e.
//
// Solidity: function assetHolders(address ) view returns(bool)
func (_Staking *StakingCaller) AssetHolders(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "assetHolders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetHolders is a free data retrieval call binding the contract method 0x6b535d7e.
//
// Solidity: function assetHolders(address ) view returns(bool)
func (_Staking *StakingSession) AssetHolders(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AssetHolders(&_Staking.CallOpts, arg0)
}

// AssetHolders is a free data retrieval call binding the contract method 0x6b535d7e.
//
// Solidity: function assetHolders(address ) view returns(bool)
func (_Staking *StakingCallerSession) AssetHolders(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AssetHolders(&_Staking.CallOpts, arg0)
}

// ChannelDisputeEpochs is a free data retrieval call binding the contract method 0xba8c8193.
//
// Solidity: function channelDisputeEpochs() view returns(uint32)
func (_Staking *StakingCaller) ChannelDisputeEpochs(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "channelDisputeEpochs")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChannelDisputeEpochs is a free data retrieval call binding the contract method 0xba8c8193.
//
// Solidity: function channelDisputeEpochs() view returns(uint32)
func (_Staking *StakingSession) ChannelDisputeEpochs() (uint32, error) {
	return _Staking.Contract.ChannelDisputeEpochs(&_Staking.CallOpts)
}

// ChannelDisputeEpochs is a free data retrieval call binding the contract method 0xba8c8193.
//
// Solidity: function channelDisputeEpochs() view returns(uint32)
func (_Staking *StakingCallerSession) ChannelDisputeEpochs() (uint32, error) {
	return _Staking.Contract.ChannelDisputeEpochs(&_Staking.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Staking *StakingCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Staking *StakingSession) Controller() (common.Address, error) {
	return _Staking.Contract.Controller(&_Staking.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Staking *StakingCallerSession) Controller() (common.Address, error) {
	return _Staking.Contract.Controller(&_Staking.CallOpts)
}

// CurationPercentage is a free data retrieval call binding the contract method 0x85b52ad0.
//
// Solidity: function curationPercentage() view returns(uint32)
func (_Staking *StakingCaller) CurationPercentage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "curationPercentage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurationPercentage is a free data retrieval call binding the contract method 0x85b52ad0.
//
// Solidity: function curationPercentage() view returns(uint32)
func (_Staking *StakingSession) CurationPercentage() (uint32, error) {
	return _Staking.Contract.CurationPercentage(&_Staking.CallOpts)
}

// CurationPercentage is a free data retrieval call binding the contract method 0x85b52ad0.
//
// Solidity: function curationPercentage() view returns(uint32)
func (_Staking *StakingCallerSession) CurationPercentage() (uint32, error) {
	return _Staking.Contract.CurationPercentage(&_Staking.CallOpts)
}

// DelegationParametersCooldown is a free data retrieval call binding the contract method 0x8a7ff87d.
//
// Solidity: function delegationParametersCooldown() view returns(uint32)
func (_Staking *StakingCaller) DelegationParametersCooldown(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegationParametersCooldown")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DelegationParametersCooldown is a free data retrieval call binding the contract method 0x8a7ff87d.
//
// Solidity: function delegationParametersCooldown() view returns(uint32)
func (_Staking *StakingSession) DelegationParametersCooldown() (uint32, error) {
	return _Staking.Contract.DelegationParametersCooldown(&_Staking.CallOpts)
}

// DelegationParametersCooldown is a free data retrieval call binding the contract method 0x8a7ff87d.
//
// Solidity: function delegationParametersCooldown() view returns(uint32)
func (_Staking *StakingCallerSession) DelegationParametersCooldown() (uint32, error) {
	return _Staking.Contract.DelegationParametersCooldown(&_Staking.CallOpts)
}

// DelegationPools is a free data retrieval call binding the contract method 0x92511c8f.
//
// Solidity: function delegationPools(address ) view returns(uint32 cooldownBlocks, uint32 indexingRewardCut, uint32 queryFeeCut, uint256 updatedAtBlock, uint256 tokens, uint256 shares)
func (_Staking *StakingCaller) DelegationPools(opts *bind.CallOpts, arg0 common.Address) (struct {
	CooldownBlocks    uint32
	IndexingRewardCut uint32
	QueryFeeCut       uint32
	UpdatedAtBlock    *big.Int
	Tokens            *big.Int
	Shares            *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegationPools", arg0)

	outstruct := new(struct {
		CooldownBlocks    uint32
		IndexingRewardCut uint32
		QueryFeeCut       uint32
		UpdatedAtBlock    *big.Int
		Tokens            *big.Int
		Shares            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CooldownBlocks = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.IndexingRewardCut = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.QueryFeeCut = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.UpdatedAtBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Tokens = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Shares = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelegationPools is a free data retrieval call binding the contract method 0x92511c8f.
//
// Solidity: function delegationPools(address ) view returns(uint32 cooldownBlocks, uint32 indexingRewardCut, uint32 queryFeeCut, uint256 updatedAtBlock, uint256 tokens, uint256 shares)
func (_Staking *StakingSession) DelegationPools(arg0 common.Address) (struct {
	CooldownBlocks    uint32
	IndexingRewardCut uint32
	QueryFeeCut       uint32
	UpdatedAtBlock    *big.Int
	Tokens            *big.Int
	Shares            *big.Int
}, error) {
	return _Staking.Contract.DelegationPools(&_Staking.CallOpts, arg0)
}

// DelegationPools is a free data retrieval call binding the contract method 0x92511c8f.
//
// Solidity: function delegationPools(address ) view returns(uint32 cooldownBlocks, uint32 indexingRewardCut, uint32 queryFeeCut, uint256 updatedAtBlock, uint256 tokens, uint256 shares)
func (_Staking *StakingCallerSession) DelegationPools(arg0 common.Address) (struct {
	CooldownBlocks    uint32
	IndexingRewardCut uint32
	QueryFeeCut       uint32
	UpdatedAtBlock    *big.Int
	Tokens            *big.Int
	Shares            *big.Int
}, error) {
	return _Staking.Contract.DelegationPools(&_Staking.CallOpts, arg0)
}

// DelegationRatio is a free data retrieval call binding the contract method 0xbfdfa7af.
//
// Solidity: function delegationRatio() view returns(uint32)
func (_Staking *StakingCaller) DelegationRatio(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegationRatio")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DelegationRatio is a free data retrieval call binding the contract method 0xbfdfa7af.
//
// Solidity: function delegationRatio() view returns(uint32)
func (_Staking *StakingSession) DelegationRatio() (uint32, error) {
	return _Staking.Contract.DelegationRatio(&_Staking.CallOpts)
}

// DelegationRatio is a free data retrieval call binding the contract method 0xbfdfa7af.
//
// Solidity: function delegationRatio() view returns(uint32)
func (_Staking *StakingCallerSession) DelegationRatio() (uint32, error) {
	return _Staking.Contract.DelegationRatio(&_Staking.CallOpts)
}

// DelegationTaxPercentage is a free data retrieval call binding the contract method 0xe6aeb796.
//
// Solidity: function delegationTaxPercentage() view returns(uint32)
func (_Staking *StakingCaller) DelegationTaxPercentage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegationTaxPercentage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DelegationTaxPercentage is a free data retrieval call binding the contract method 0xe6aeb796.
//
// Solidity: function delegationTaxPercentage() view returns(uint32)
func (_Staking *StakingSession) DelegationTaxPercentage() (uint32, error) {
	return _Staking.Contract.DelegationTaxPercentage(&_Staking.CallOpts)
}

// DelegationTaxPercentage is a free data retrieval call binding the contract method 0xe6aeb796.
//
// Solidity: function delegationTaxPercentage() view returns(uint32)
func (_Staking *StakingCallerSession) DelegationTaxPercentage() (uint32, error) {
	return _Staking.Contract.DelegationTaxPercentage(&_Staking.CallOpts)
}

// DelegationUnbondingPeriod is a free data retrieval call binding the contract method 0xb6846e47.
//
// Solidity: function delegationUnbondingPeriod() view returns(uint32)
func (_Staking *StakingCaller) DelegationUnbondingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegationUnbondingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DelegationUnbondingPeriod is a free data retrieval call binding the contract method 0xb6846e47.
//
// Solidity: function delegationUnbondingPeriod() view returns(uint32)
func (_Staking *StakingSession) DelegationUnbondingPeriod() (uint32, error) {
	return _Staking.Contract.DelegationUnbondingPeriod(&_Staking.CallOpts)
}

// DelegationUnbondingPeriod is a free data retrieval call binding the contract method 0xb6846e47.
//
// Solidity: function delegationUnbondingPeriod() view returns(uint32)
func (_Staking *StakingCallerSession) DelegationUnbondingPeriod() (uint32, error) {
	return _Staking.Contract.DelegationUnbondingPeriod(&_Staking.CallOpts)
}

// GetAllocation is a free data retrieval call binding the contract method 0x0e022923.
//
// Solidity: function getAllocation(address _allocationID) view returns((address,bytes32,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Staking *StakingCaller) GetAllocation(opts *bind.CallOpts, _allocationID common.Address) (IStakingDataAllocation, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getAllocation", _allocationID)

	if err != nil {
		return *new(IStakingDataAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingDataAllocation)).(*IStakingDataAllocation)

	return out0, err

}

// GetAllocation is a free data retrieval call binding the contract method 0x0e022923.
//
// Solidity: function getAllocation(address _allocationID) view returns((address,bytes32,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Staking *StakingSession) GetAllocation(_allocationID common.Address) (IStakingDataAllocation, error) {
	return _Staking.Contract.GetAllocation(&_Staking.CallOpts, _allocationID)
}

// GetAllocation is a free data retrieval call binding the contract method 0x0e022923.
//
// Solidity: function getAllocation(address _allocationID) view returns((address,bytes32,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Staking *StakingCallerSession) GetAllocation(_allocationID common.Address) (IStakingDataAllocation, error) {
	return _Staking.Contract.GetAllocation(&_Staking.CallOpts, _allocationID)
}

// GetAllocationState is a free data retrieval call binding the contract method 0x98c657dc.
//
// Solidity: function getAllocationState(address _allocationID) view returns(uint8)
func (_Staking *StakingCaller) GetAllocationState(opts *bind.CallOpts, _allocationID common.Address) (uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getAllocationState", _allocationID)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAllocationState is a free data retrieval call binding the contract method 0x98c657dc.
//
// Solidity: function getAllocationState(address _allocationID) view returns(uint8)
func (_Staking *StakingSession) GetAllocationState(_allocationID common.Address) (uint8, error) {
	return _Staking.Contract.GetAllocationState(&_Staking.CallOpts, _allocationID)
}

// GetAllocationState is a free data retrieval call binding the contract method 0x98c657dc.
//
// Solidity: function getAllocationState(address _allocationID) view returns(uint8)
func (_Staking *StakingCallerSession) GetAllocationState(_allocationID common.Address) (uint8, error) {
	return _Staking.Contract.GetAllocationState(&_Staking.CallOpts, _allocationID)
}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address _indexer, address _delegator) view returns((uint256,uint256,uint256))
func (_Staking *StakingCaller) GetDelegation(opts *bind.CallOpts, _indexer common.Address, _delegator common.Address) (IStakingDataDelegation, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDelegation", _indexer, _delegator)

	if err != nil {
		return *new(IStakingDataDelegation), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingDataDelegation)).(*IStakingDataDelegation)

	return out0, err

}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address _indexer, address _delegator) view returns((uint256,uint256,uint256))
func (_Staking *StakingSession) GetDelegation(_indexer common.Address, _delegator common.Address) (IStakingDataDelegation, error) {
	return _Staking.Contract.GetDelegation(&_Staking.CallOpts, _indexer, _delegator)
}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address _indexer, address _delegator) view returns((uint256,uint256,uint256))
func (_Staking *StakingCallerSession) GetDelegation(_indexer common.Address, _delegator common.Address) (IStakingDataDelegation, error) {
	return _Staking.Contract.GetDelegation(&_Staking.CallOpts, _indexer, _delegator)
}

// GetIndexerCapacity is a free data retrieval call binding the contract method 0xa510be20.
//
// Solidity: function getIndexerCapacity(address _indexer) view returns(uint256)
func (_Staking *StakingCaller) GetIndexerCapacity(opts *bind.CallOpts, _indexer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getIndexerCapacity", _indexer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexerCapacity is a free data retrieval call binding the contract method 0xa510be20.
//
// Solidity: function getIndexerCapacity(address _indexer) view returns(uint256)
func (_Staking *StakingSession) GetIndexerCapacity(_indexer common.Address) (*big.Int, error) {
	return _Staking.Contract.GetIndexerCapacity(&_Staking.CallOpts, _indexer)
}

// GetIndexerCapacity is a free data retrieval call binding the contract method 0xa510be20.
//
// Solidity: function getIndexerCapacity(address _indexer) view returns(uint256)
func (_Staking *StakingCallerSession) GetIndexerCapacity(_indexer common.Address) (*big.Int, error) {
	return _Staking.Contract.GetIndexerCapacity(&_Staking.CallOpts, _indexer)
}

// GetIndexerStakedTokens is a free data retrieval call binding the contract method 0x1787e69f.
//
// Solidity: function getIndexerStakedTokens(address _indexer) view returns(uint256)
func (_Staking *StakingCaller) GetIndexerStakedTokens(opts *bind.CallOpts, _indexer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getIndexerStakedTokens", _indexer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexerStakedTokens is a free data retrieval call binding the contract method 0x1787e69f.
//
// Solidity: function getIndexerStakedTokens(address _indexer) view returns(uint256)
func (_Staking *StakingSession) GetIndexerStakedTokens(_indexer common.Address) (*big.Int, error) {
	return _Staking.Contract.GetIndexerStakedTokens(&_Staking.CallOpts, _indexer)
}

// GetIndexerStakedTokens is a free data retrieval call binding the contract method 0x1787e69f.
//
// Solidity: function getIndexerStakedTokens(address _indexer) view returns(uint256)
func (_Staking *StakingCallerSession) GetIndexerStakedTokens(_indexer common.Address) (*big.Int, error) {
	return _Staking.Contract.GetIndexerStakedTokens(&_Staking.CallOpts, _indexer)
}

// GetSubgraphAllocatedTokens is a free data retrieval call binding the contract method 0xe2e1e8e9.
//
// Solidity: function getSubgraphAllocatedTokens(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_Staking *StakingCaller) GetSubgraphAllocatedTokens(opts *bind.CallOpts, _subgraphDeploymentID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getSubgraphAllocatedTokens", _subgraphDeploymentID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubgraphAllocatedTokens is a free data retrieval call binding the contract method 0xe2e1e8e9.
//
// Solidity: function getSubgraphAllocatedTokens(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_Staking *StakingSession) GetSubgraphAllocatedTokens(_subgraphDeploymentID [32]byte) (*big.Int, error) {
	return _Staking.Contract.GetSubgraphAllocatedTokens(&_Staking.CallOpts, _subgraphDeploymentID)
}

// GetSubgraphAllocatedTokens is a free data retrieval call binding the contract method 0xe2e1e8e9.
//
// Solidity: function getSubgraphAllocatedTokens(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_Staking *StakingCallerSession) GetSubgraphAllocatedTokens(_subgraphDeploymentID [32]byte) (*big.Int, error) {
	return _Staking.Contract.GetSubgraphAllocatedTokens(&_Staking.CallOpts, _subgraphDeploymentID)
}

// GetWithdraweableDelegatedTokens is a free data retrieval call binding the contract method 0x130bea57.
//
// Solidity: function getWithdraweableDelegatedTokens((uint256,uint256,uint256) _delegation) view returns(uint256)
func (_Staking *StakingCaller) GetWithdraweableDelegatedTokens(opts *bind.CallOpts, _delegation IStakingDataDelegation) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getWithdraweableDelegatedTokens", _delegation)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdraweableDelegatedTokens is a free data retrieval call binding the contract method 0x130bea57.
//
// Solidity: function getWithdraweableDelegatedTokens((uint256,uint256,uint256) _delegation) view returns(uint256)
func (_Staking *StakingSession) GetWithdraweableDelegatedTokens(_delegation IStakingDataDelegation) (*big.Int, error) {
	return _Staking.Contract.GetWithdraweableDelegatedTokens(&_Staking.CallOpts, _delegation)
}

// GetWithdraweableDelegatedTokens is a free data retrieval call binding the contract method 0x130bea57.
//
// Solidity: function getWithdraweableDelegatedTokens((uint256,uint256,uint256) _delegation) view returns(uint256)
func (_Staking *StakingCallerSession) GetWithdraweableDelegatedTokens(_delegation IStakingDataDelegation) (*big.Int, error) {
	return _Staking.Contract.GetWithdraweableDelegatedTokens(&_Staking.CallOpts, _delegation)
}

// HasStake is a free data retrieval call binding the contract method 0xe73e14bf.
//
// Solidity: function hasStake(address _indexer) view returns(bool)
func (_Staking *StakingCaller) HasStake(opts *bind.CallOpts, _indexer common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "hasStake", _indexer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasStake is a free data retrieval call binding the contract method 0xe73e14bf.
//
// Solidity: function hasStake(address _indexer) view returns(bool)
func (_Staking *StakingSession) HasStake(_indexer common.Address) (bool, error) {
	return _Staking.Contract.HasStake(&_Staking.CallOpts, _indexer)
}

// HasStake is a free data retrieval call binding the contract method 0xe73e14bf.
//
// Solidity: function hasStake(address _indexer) view returns(bool)
func (_Staking *StakingCallerSession) HasStake(_indexer common.Address) (bool, error) {
	return _Staking.Contract.HasStake(&_Staking.CallOpts, _indexer)
}

// IsAllocation is a free data retrieval call binding the contract method 0xf1d60d66.
//
// Solidity: function isAllocation(address _allocationID) view returns(bool)
func (_Staking *StakingCaller) IsAllocation(opts *bind.CallOpts, _allocationID common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isAllocation", _allocationID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllocation is a free data retrieval call binding the contract method 0xf1d60d66.
//
// Solidity: function isAllocation(address _allocationID) view returns(bool)
func (_Staking *StakingSession) IsAllocation(_allocationID common.Address) (bool, error) {
	return _Staking.Contract.IsAllocation(&_Staking.CallOpts, _allocationID)
}

// IsAllocation is a free data retrieval call binding the contract method 0xf1d60d66.
//
// Solidity: function isAllocation(address _allocationID) view returns(bool)
func (_Staking *StakingCallerSession) IsAllocation(_allocationID common.Address) (bool, error) {
	return _Staking.Contract.IsAllocation(&_Staking.CallOpts, _allocationID)
}

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address _indexer, address _delegator) view returns(bool)
func (_Staking *StakingCaller) IsDelegator(opts *bind.CallOpts, _indexer common.Address, _delegator common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isDelegator", _indexer, _delegator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address _indexer, address _delegator) view returns(bool)
func (_Staking *StakingSession) IsDelegator(_indexer common.Address, _delegator common.Address) (bool, error) {
	return _Staking.Contract.IsDelegator(&_Staking.CallOpts, _indexer, _delegator)
}

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address _indexer, address _delegator) view returns(bool)
func (_Staking *StakingCallerSession) IsDelegator(_indexer common.Address, _delegator common.Address) (bool, error) {
	return _Staking.Contract.IsDelegator(&_Staking.CallOpts, _indexer, _delegator)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _operator, address _indexer) view returns(bool)
func (_Staking *StakingCaller) IsOperator(opts *bind.CallOpts, _operator common.Address, _indexer common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isOperator", _operator, _indexer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _operator, address _indexer) view returns(bool)
func (_Staking *StakingSession) IsOperator(_operator common.Address, _indexer common.Address) (bool, error) {
	return _Staking.Contract.IsOperator(&_Staking.CallOpts, _operator, _indexer)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _operator, address _indexer) view returns(bool)
func (_Staking *StakingCallerSession) IsOperator(_operator common.Address, _indexer common.Address) (bool, error) {
	return _Staking.Contract.IsOperator(&_Staking.CallOpts, _operator, _indexer)
}

// MaxAllocationEpochs is a free data retrieval call binding the contract method 0xfb765938.
//
// Solidity: function maxAllocationEpochs() view returns(uint32)
func (_Staking *StakingCaller) MaxAllocationEpochs(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "maxAllocationEpochs")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxAllocationEpochs is a free data retrieval call binding the contract method 0xfb765938.
//
// Solidity: function maxAllocationEpochs() view returns(uint32)
func (_Staking *StakingSession) MaxAllocationEpochs() (uint32, error) {
	return _Staking.Contract.MaxAllocationEpochs(&_Staking.CallOpts)
}

// MaxAllocationEpochs is a free data retrieval call binding the contract method 0xfb765938.
//
// Solidity: function maxAllocationEpochs() view returns(uint32)
func (_Staking *StakingCallerSession) MaxAllocationEpochs() (uint32, error) {
	return _Staking.Contract.MaxAllocationEpochs(&_Staking.CallOpts)
}

// MinimumIndexerStake is a free data retrieval call binding the contract method 0xf2485bf2.
//
// Solidity: function minimumIndexerStake() view returns(uint256)
func (_Staking *StakingCaller) MinimumIndexerStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "minimumIndexerStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumIndexerStake is a free data retrieval call binding the contract method 0xf2485bf2.
//
// Solidity: function minimumIndexerStake() view returns(uint256)
func (_Staking *StakingSession) MinimumIndexerStake() (*big.Int, error) {
	return _Staking.Contract.MinimumIndexerStake(&_Staking.CallOpts)
}

// MinimumIndexerStake is a free data retrieval call binding the contract method 0xf2485bf2.
//
// Solidity: function minimumIndexerStake() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumIndexerStake() (*big.Int, error) {
	return _Staking.Contract.MinimumIndexerStake(&_Staking.CallOpts)
}

// OperatorAuth is a free data retrieval call binding the contract method 0xe2e94656.
//
// Solidity: function operatorAuth(address , address ) view returns(bool)
func (_Staking *StakingCaller) OperatorAuth(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "operatorAuth", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorAuth is a free data retrieval call binding the contract method 0xe2e94656.
//
// Solidity: function operatorAuth(address , address ) view returns(bool)
func (_Staking *StakingSession) OperatorAuth(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Staking.Contract.OperatorAuth(&_Staking.CallOpts, arg0, arg1)
}

// OperatorAuth is a free data retrieval call binding the contract method 0xe2e94656.
//
// Solidity: function operatorAuth(address , address ) view returns(bool)
func (_Staking *StakingCallerSession) OperatorAuth(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Staking.Contract.OperatorAuth(&_Staking.CallOpts, arg0, arg1)
}

// ProtocolPercentage is a free data retrieval call binding the contract method 0xa26b90f2.
//
// Solidity: function protocolPercentage() view returns(uint32)
func (_Staking *StakingCaller) ProtocolPercentage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "protocolPercentage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ProtocolPercentage is a free data retrieval call binding the contract method 0xa26b90f2.
//
// Solidity: function protocolPercentage() view returns(uint32)
func (_Staking *StakingSession) ProtocolPercentage() (uint32, error) {
	return _Staking.Contract.ProtocolPercentage(&_Staking.CallOpts)
}

// ProtocolPercentage is a free data retrieval call binding the contract method 0xa26b90f2.
//
// Solidity: function protocolPercentage() view returns(uint32)
func (_Staking *StakingCallerSession) ProtocolPercentage() (uint32, error) {
	return _Staking.Contract.ProtocolPercentage(&_Staking.CallOpts)
}

// Rebates is a free data retrieval call binding the contract method 0xd3cb644c.
//
// Solidity: function rebates(uint256 ) view returns(uint256 fees, uint256 effectiveAllocatedStake, uint256 claimedRewards, uint32 unclaimedAllocationsCount, uint32 alphaNumerator, uint32 alphaDenominator)
func (_Staking *StakingCaller) Rebates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Fees                      *big.Int
	EffectiveAllocatedStake   *big.Int
	ClaimedRewards            *big.Int
	UnclaimedAllocationsCount uint32
	AlphaNumerator            uint32
	AlphaDenominator          uint32
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rebates", arg0)

	outstruct := new(struct {
		Fees                      *big.Int
		EffectiveAllocatedStake   *big.Int
		ClaimedRewards            *big.Int
		UnclaimedAllocationsCount uint32
		AlphaNumerator            uint32
		AlphaDenominator          uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fees = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EffectiveAllocatedStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ClaimedRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnclaimedAllocationsCount = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AlphaNumerator = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.AlphaDenominator = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// Rebates is a free data retrieval call binding the contract method 0xd3cb644c.
//
// Solidity: function rebates(uint256 ) view returns(uint256 fees, uint256 effectiveAllocatedStake, uint256 claimedRewards, uint32 unclaimedAllocationsCount, uint32 alphaNumerator, uint32 alphaDenominator)
func (_Staking *StakingSession) Rebates(arg0 *big.Int) (struct {
	Fees                      *big.Int
	EffectiveAllocatedStake   *big.Int
	ClaimedRewards            *big.Int
	UnclaimedAllocationsCount uint32
	AlphaNumerator            uint32
	AlphaDenominator          uint32
}, error) {
	return _Staking.Contract.Rebates(&_Staking.CallOpts, arg0)
}

// Rebates is a free data retrieval call binding the contract method 0xd3cb644c.
//
// Solidity: function rebates(uint256 ) view returns(uint256 fees, uint256 effectiveAllocatedStake, uint256 claimedRewards, uint32 unclaimedAllocationsCount, uint32 alphaNumerator, uint32 alphaDenominator)
func (_Staking *StakingCallerSession) Rebates(arg0 *big.Int) (struct {
	Fees                      *big.Int
	EffectiveAllocatedStake   *big.Int
	ClaimedRewards            *big.Int
	UnclaimedAllocationsCount uint32
	AlphaNumerator            uint32
	AlphaDenominator          uint32
}, error) {
	return _Staking.Contract.Rebates(&_Staking.CallOpts, arg0)
}

// RewardsDestination is a free data retrieval call binding the contract method 0x7203ca78.
//
// Solidity: function rewardsDestination(address ) view returns(address)
func (_Staking *StakingCaller) RewardsDestination(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardsDestination", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDestination is a free data retrieval call binding the contract method 0x7203ca78.
//
// Solidity: function rewardsDestination(address ) view returns(address)
func (_Staking *StakingSession) RewardsDestination(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.RewardsDestination(&_Staking.CallOpts, arg0)
}

// RewardsDestination is a free data retrieval call binding the contract method 0x7203ca78.
//
// Solidity: function rewardsDestination(address ) view returns(address)
func (_Staking *StakingCallerSession) RewardsDestination(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.RewardsDestination(&_Staking.CallOpts, arg0)
}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_Staking *StakingCaller) Slashers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "slashers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_Staking *StakingSession) Slashers(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Slashers(&_Staking.CallOpts, arg0)
}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_Staking *StakingCallerSession) Slashers(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Slashers(&_Staking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 tokensStaked, uint256 tokensAllocated, uint256 tokensLocked, uint256 tokensLockedUntil)
func (_Staking *StakingCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	TokensStaked      *big.Int
	TokensAllocated   *big.Int
	TokensLocked      *big.Int
	TokensLockedUntil *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		TokensStaked      *big.Int
		TokensAllocated   *big.Int
		TokensLocked      *big.Int
		TokensLockedUntil *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokensStaked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokensAllocated = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokensLocked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokensLockedUntil = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 tokensStaked, uint256 tokensAllocated, uint256 tokensLocked, uint256 tokensLockedUntil)
func (_Staking *StakingSession) Stakes(arg0 common.Address) (struct {
	TokensStaked      *big.Int
	TokensAllocated   *big.Int
	TokensLocked      *big.Int
	TokensLockedUntil *big.Int
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 tokensStaked, uint256 tokensAllocated, uint256 tokensLocked, uint256 tokensLockedUntil)
func (_Staking *StakingCallerSession) Stakes(arg0 common.Address) (struct {
	TokensStaked      *big.Int
	TokensAllocated   *big.Int
	TokensLocked      *big.Int
	TokensLockedUntil *big.Int
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// SubgraphAllocations is a free data retrieval call binding the contract method 0xb1468f52.
//
// Solidity: function subgraphAllocations(bytes32 ) view returns(uint256)
func (_Staking *StakingCaller) SubgraphAllocations(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "subgraphAllocations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubgraphAllocations is a free data retrieval call binding the contract method 0xb1468f52.
//
// Solidity: function subgraphAllocations(bytes32 ) view returns(uint256)
func (_Staking *StakingSession) SubgraphAllocations(arg0 [32]byte) (*big.Int, error) {
	return _Staking.Contract.SubgraphAllocations(&_Staking.CallOpts, arg0)
}

// SubgraphAllocations is a free data retrieval call binding the contract method 0xb1468f52.
//
// Solidity: function subgraphAllocations(bytes32 ) view returns(uint256)
func (_Staking *StakingCallerSession) SubgraphAllocations(arg0 [32]byte) (*big.Int, error) {
	return _Staking.Contract.SubgraphAllocations(&_Staking.CallOpts, arg0)
}

// ThawingPeriod is a free data retrieval call binding the contract method 0xcdc747dd.
//
// Solidity: function thawingPeriod() view returns(uint32)
func (_Staking *StakingCaller) ThawingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "thawingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ThawingPeriod is a free data retrieval call binding the contract method 0xcdc747dd.
//
// Solidity: function thawingPeriod() view returns(uint32)
func (_Staking *StakingSession) ThawingPeriod() (uint32, error) {
	return _Staking.Contract.ThawingPeriod(&_Staking.CallOpts)
}

// ThawingPeriod is a free data retrieval call binding the contract method 0xcdc747dd.
//
// Solidity: function thawingPeriod() view returns(uint32)
func (_Staking *StakingCallerSession) ThawingPeriod() (uint32, error) {
	return _Staking.Contract.ThawingPeriod(&_Staking.CallOpts)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_Staking *StakingTransactor) AcceptProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "acceptProxy", _proxy)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_Staking *StakingSession) AcceptProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AcceptProxy(&_Staking.TransactOpts, _proxy)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_Staking *StakingTransactorSession) AcceptProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AcceptProxy(&_Staking.TransactOpts, _proxy)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_Staking *StakingTransactor) AcceptProxyAndCall(opts *bind.TransactOpts, _proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "acceptProxyAndCall", _proxy, _data)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_Staking *StakingSession) AcceptProxyAndCall(_proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _Staking.Contract.AcceptProxyAndCall(&_Staking.TransactOpts, _proxy, _data)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_Staking *StakingTransactorSession) AcceptProxyAndCall(_proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _Staking.Contract.AcceptProxyAndCall(&_Staking.TransactOpts, _proxy, _data)
}

// Allocate is a paid mutator transaction binding the contract method 0xa6fe292b.
//
// Solidity: function allocate(bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingTransactor) Allocate(opts *bind.TransactOpts, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "allocate", _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// Allocate is a paid mutator transaction binding the contract method 0xa6fe292b.
//
// Solidity: function allocate(bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingSession) Allocate(_subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.Contract.Allocate(&_Staking.TransactOpts, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// Allocate is a paid mutator transaction binding the contract method 0xa6fe292b.
//
// Solidity: function allocate(bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingTransactorSession) Allocate(_subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.Contract.Allocate(&_Staking.TransactOpts, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// AllocateFrom is a paid mutator transaction binding the contract method 0x23477e48.
//
// Solidity: function allocateFrom(address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingTransactor) AllocateFrom(opts *bind.TransactOpts, _indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "allocateFrom", _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// AllocateFrom is a paid mutator transaction binding the contract method 0x23477e48.
//
// Solidity: function allocateFrom(address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingSession) AllocateFrom(_indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.Contract.AllocateFrom(&_Staking.TransactOpts, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// AllocateFrom is a paid mutator transaction binding the contract method 0x23477e48.
//
// Solidity: function allocateFrom(address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingTransactorSession) AllocateFrom(_indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.Contract.AllocateFrom(&_Staking.TransactOpts, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _allocationID, bool _restake) returns()
func (_Staking *StakingTransactor) Claim(opts *bind.TransactOpts, _allocationID common.Address, _restake bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claim", _allocationID, _restake)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _allocationID, bool _restake) returns()
func (_Staking *StakingSession) Claim(_allocationID common.Address, _restake bool) (*types.Transaction, error) {
	return _Staking.Contract.Claim(&_Staking.TransactOpts, _allocationID, _restake)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _allocationID, bool _restake) returns()
func (_Staking *StakingTransactorSession) Claim(_allocationID common.Address, _restake bool) (*types.Transaction, error) {
	return _Staking.Contract.Claim(&_Staking.TransactOpts, _allocationID, _restake)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x36a4fbd6.
//
// Solidity: function claimMany(address[] _allocationID, bool _restake) returns()
func (_Staking *StakingTransactor) ClaimMany(opts *bind.TransactOpts, _allocationID []common.Address, _restake bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claimMany", _allocationID, _restake)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x36a4fbd6.
//
// Solidity: function claimMany(address[] _allocationID, bool _restake) returns()
func (_Staking *StakingSession) ClaimMany(_allocationID []common.Address, _restake bool) (*types.Transaction, error) {
	return _Staking.Contract.ClaimMany(&_Staking.TransactOpts, _allocationID, _restake)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x36a4fbd6.
//
// Solidity: function claimMany(address[] _allocationID, bool _restake) returns()
func (_Staking *StakingTransactorSession) ClaimMany(_allocationID []common.Address, _restake bool) (*types.Transaction, error) {
	return _Staking.Contract.ClaimMany(&_Staking.TransactOpts, _allocationID, _restake)
}

// CloseAllocation is a paid mutator transaction binding the contract method 0x44c32a61.
//
// Solidity: function closeAllocation(address _allocationID, bytes32 _poi) returns()
func (_Staking *StakingTransactor) CloseAllocation(opts *bind.TransactOpts, _allocationID common.Address, _poi [32]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "closeAllocation", _allocationID, _poi)
}

// CloseAllocation is a paid mutator transaction binding the contract method 0x44c32a61.
//
// Solidity: function closeAllocation(address _allocationID, bytes32 _poi) returns()
func (_Staking *StakingSession) CloseAllocation(_allocationID common.Address, _poi [32]byte) (*types.Transaction, error) {
	return _Staking.Contract.CloseAllocation(&_Staking.TransactOpts, _allocationID, _poi)
}

// CloseAllocation is a paid mutator transaction binding the contract method 0x44c32a61.
//
// Solidity: function closeAllocation(address _allocationID, bytes32 _poi) returns()
func (_Staking *StakingTransactorSession) CloseAllocation(_allocationID common.Address, _poi [32]byte) (*types.Transaction, error) {
	return _Staking.Contract.CloseAllocation(&_Staking.TransactOpts, _allocationID, _poi)
}

// CloseAllocationMany is a paid mutator transaction binding the contract method 0x0d851d97.
//
// Solidity: function closeAllocationMany((address,bytes32)[] _requests) returns()
func (_Staking *StakingTransactor) CloseAllocationMany(opts *bind.TransactOpts, _requests []IStakingCloseAllocationRequest) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "closeAllocationMany", _requests)
}

// CloseAllocationMany is a paid mutator transaction binding the contract method 0x0d851d97.
//
// Solidity: function closeAllocationMany((address,bytes32)[] _requests) returns()
func (_Staking *StakingSession) CloseAllocationMany(_requests []IStakingCloseAllocationRequest) (*types.Transaction, error) {
	return _Staking.Contract.CloseAllocationMany(&_Staking.TransactOpts, _requests)
}

// CloseAllocationMany is a paid mutator transaction binding the contract method 0x0d851d97.
//
// Solidity: function closeAllocationMany((address,bytes32)[] _requests) returns()
func (_Staking *StakingTransactorSession) CloseAllocationMany(_requests []IStakingCloseAllocationRequest) (*types.Transaction, error) {
	return _Staking.Contract.CloseAllocationMany(&_Staking.TransactOpts, _requests)
}

// CloseAndAllocate is a paid mutator transaction binding the contract method 0xc2b6df37.
//
// Solidity: function closeAndAllocate(address _closingAllocationID, bytes32 _poi, address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingTransactor) CloseAndAllocate(opts *bind.TransactOpts, _closingAllocationID common.Address, _poi [32]byte, _indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "closeAndAllocate", _closingAllocationID, _poi, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// CloseAndAllocate is a paid mutator transaction binding the contract method 0xc2b6df37.
//
// Solidity: function closeAndAllocate(address _closingAllocationID, bytes32 _poi, address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingSession) CloseAndAllocate(_closingAllocationID common.Address, _poi [32]byte, _indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.Contract.CloseAndAllocate(&_Staking.TransactOpts, _closingAllocationID, _poi, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// CloseAndAllocate is a paid mutator transaction binding the contract method 0xc2b6df37.
//
// Solidity: function closeAndAllocate(address _closingAllocationID, bytes32 _poi, address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_Staking *StakingTransactorSession) CloseAndAllocate(_closingAllocationID common.Address, _poi [32]byte, _indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _Staking.Contract.CloseAndAllocate(&_Staking.TransactOpts, _closingAllocationID, _poi, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// Collect is a paid mutator transaction binding the contract method 0x8d3c100a.
//
// Solidity: function collect(uint256 _tokens, address _allocationID) returns()
func (_Staking *StakingTransactor) Collect(opts *bind.TransactOpts, _tokens *big.Int, _allocationID common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "collect", _tokens, _allocationID)
}

// Collect is a paid mutator transaction binding the contract method 0x8d3c100a.
//
// Solidity: function collect(uint256 _tokens, address _allocationID) returns()
func (_Staking *StakingSession) Collect(_tokens *big.Int, _allocationID common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Collect(&_Staking.TransactOpts, _tokens, _allocationID)
}

// Collect is a paid mutator transaction binding the contract method 0x8d3c100a.
//
// Solidity: function collect(uint256 _tokens, address _allocationID) returns()
func (_Staking *StakingTransactorSession) Collect(_tokens *big.Int, _allocationID common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Collect(&_Staking.TransactOpts, _tokens, _allocationID)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _indexer, uint256 _tokens) returns(uint256)
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, _indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", _indexer, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _indexer, uint256 _tokens) returns(uint256)
func (_Staking *StakingSession) Delegate(_indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, _indexer, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _indexer, uint256 _tokens) returns(uint256)
func (_Staking *StakingTransactorSession) Delegate(_indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, _indexer, _tokens)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fc48624.
//
// Solidity: function initialize(address _controller, uint256 _minimumIndexerStake, uint32 _thawingPeriod, uint32 _protocolPercentage, uint32 _curationPercentage, uint32 _channelDisputeEpochs, uint32 _maxAllocationEpochs, uint32 _delegationUnbondingPeriod, uint32 _delegationRatio, uint32 _rebateAlphaNumerator, uint32 _rebateAlphaDenominator) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _controller common.Address, _minimumIndexerStake *big.Int, _thawingPeriod uint32, _protocolPercentage uint32, _curationPercentage uint32, _channelDisputeEpochs uint32, _maxAllocationEpochs uint32, _delegationUnbondingPeriod uint32, _delegationRatio uint32, _rebateAlphaNumerator uint32, _rebateAlphaDenominator uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _controller, _minimumIndexerStake, _thawingPeriod, _protocolPercentage, _curationPercentage, _channelDisputeEpochs, _maxAllocationEpochs, _delegationUnbondingPeriod, _delegationRatio, _rebateAlphaNumerator, _rebateAlphaDenominator)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fc48624.
//
// Solidity: function initialize(address _controller, uint256 _minimumIndexerStake, uint32 _thawingPeriod, uint32 _protocolPercentage, uint32 _curationPercentage, uint32 _channelDisputeEpochs, uint32 _maxAllocationEpochs, uint32 _delegationUnbondingPeriod, uint32 _delegationRatio, uint32 _rebateAlphaNumerator, uint32 _rebateAlphaDenominator) returns()
func (_Staking *StakingSession) Initialize(_controller common.Address, _minimumIndexerStake *big.Int, _thawingPeriod uint32, _protocolPercentage uint32, _curationPercentage uint32, _channelDisputeEpochs uint32, _maxAllocationEpochs uint32, _delegationUnbondingPeriod uint32, _delegationRatio uint32, _rebateAlphaNumerator uint32, _rebateAlphaDenominator uint32) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _controller, _minimumIndexerStake, _thawingPeriod, _protocolPercentage, _curationPercentage, _channelDisputeEpochs, _maxAllocationEpochs, _delegationUnbondingPeriod, _delegationRatio, _rebateAlphaNumerator, _rebateAlphaDenominator)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fc48624.
//
// Solidity: function initialize(address _controller, uint256 _minimumIndexerStake, uint32 _thawingPeriod, uint32 _protocolPercentage, uint32 _curationPercentage, uint32 _channelDisputeEpochs, uint32 _maxAllocationEpochs, uint32 _delegationUnbondingPeriod, uint32 _delegationRatio, uint32 _rebateAlphaNumerator, uint32 _rebateAlphaDenominator) returns()
func (_Staking *StakingTransactorSession) Initialize(_controller common.Address, _minimumIndexerStake *big.Int, _thawingPeriod uint32, _protocolPercentage uint32, _curationPercentage uint32, _channelDisputeEpochs uint32, _maxAllocationEpochs uint32, _delegationUnbondingPeriod uint32, _delegationRatio uint32, _rebateAlphaNumerator uint32, _rebateAlphaDenominator uint32) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _controller, _minimumIndexerStake, _thawingPeriod, _protocolPercentage, _curationPercentage, _channelDisputeEpochs, _maxAllocationEpochs, _delegationUnbondingPeriod, _delegationRatio, _rebateAlphaNumerator, _rebateAlphaDenominator)
}

// SetAssetHolder is a paid mutator transaction binding the contract method 0x58d7cf00.
//
// Solidity: function setAssetHolder(address _assetHolder, bool _allowed) returns()
func (_Staking *StakingTransactor) SetAssetHolder(opts *bind.TransactOpts, _assetHolder common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setAssetHolder", _assetHolder, _allowed)
}

// SetAssetHolder is a paid mutator transaction binding the contract method 0x58d7cf00.
//
// Solidity: function setAssetHolder(address _assetHolder, bool _allowed) returns()
func (_Staking *StakingSession) SetAssetHolder(_assetHolder common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.Contract.SetAssetHolder(&_Staking.TransactOpts, _assetHolder, _allowed)
}

// SetAssetHolder is a paid mutator transaction binding the contract method 0x58d7cf00.
//
// Solidity: function setAssetHolder(address _assetHolder, bool _allowed) returns()
func (_Staking *StakingTransactorSession) SetAssetHolder(_assetHolder common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.Contract.SetAssetHolder(&_Staking.TransactOpts, _assetHolder, _allowed)
}

// SetChannelDisputeEpochs is a paid mutator transaction binding the contract method 0x1d72e623.
//
// Solidity: function setChannelDisputeEpochs(uint32 _channelDisputeEpochs) returns()
func (_Staking *StakingTransactor) SetChannelDisputeEpochs(opts *bind.TransactOpts, _channelDisputeEpochs uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setChannelDisputeEpochs", _channelDisputeEpochs)
}

// SetChannelDisputeEpochs is a paid mutator transaction binding the contract method 0x1d72e623.
//
// Solidity: function setChannelDisputeEpochs(uint32 _channelDisputeEpochs) returns()
func (_Staking *StakingSession) SetChannelDisputeEpochs(_channelDisputeEpochs uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetChannelDisputeEpochs(&_Staking.TransactOpts, _channelDisputeEpochs)
}

// SetChannelDisputeEpochs is a paid mutator transaction binding the contract method 0x1d72e623.
//
// Solidity: function setChannelDisputeEpochs(uint32 _channelDisputeEpochs) returns()
func (_Staking *StakingTransactorSession) SetChannelDisputeEpochs(_channelDisputeEpochs uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetChannelDisputeEpochs(&_Staking.TransactOpts, _channelDisputeEpochs)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Staking *StakingTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Staking *StakingSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetController(&_Staking.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Staking *StakingTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetController(&_Staking.TransactOpts, _controller)
}

// SetCurationPercentage is a paid mutator transaction binding the contract method 0x39dcf476.
//
// Solidity: function setCurationPercentage(uint32 _percentage) returns()
func (_Staking *StakingTransactor) SetCurationPercentage(opts *bind.TransactOpts, _percentage uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setCurationPercentage", _percentage)
}

// SetCurationPercentage is a paid mutator transaction binding the contract method 0x39dcf476.
//
// Solidity: function setCurationPercentage(uint32 _percentage) returns()
func (_Staking *StakingSession) SetCurationPercentage(_percentage uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetCurationPercentage(&_Staking.TransactOpts, _percentage)
}

// SetCurationPercentage is a paid mutator transaction binding the contract method 0x39dcf476.
//
// Solidity: function setCurationPercentage(uint32 _percentage) returns()
func (_Staking *StakingTransactorSession) SetCurationPercentage(_percentage uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetCurationPercentage(&_Staking.TransactOpts, _percentage)
}

// SetDelegationParameters is a paid mutator transaction binding the contract method 0x9dcaa6c9.
//
// Solidity: function setDelegationParameters(uint32 _indexingRewardCut, uint32 _queryFeeCut, uint32 _cooldownBlocks) returns()
func (_Staking *StakingTransactor) SetDelegationParameters(opts *bind.TransactOpts, _indexingRewardCut uint32, _queryFeeCut uint32, _cooldownBlocks uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setDelegationParameters", _indexingRewardCut, _queryFeeCut, _cooldownBlocks)
}

// SetDelegationParameters is a paid mutator transaction binding the contract method 0x9dcaa6c9.
//
// Solidity: function setDelegationParameters(uint32 _indexingRewardCut, uint32 _queryFeeCut, uint32 _cooldownBlocks) returns()
func (_Staking *StakingSession) SetDelegationParameters(_indexingRewardCut uint32, _queryFeeCut uint32, _cooldownBlocks uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationParameters(&_Staking.TransactOpts, _indexingRewardCut, _queryFeeCut, _cooldownBlocks)
}

// SetDelegationParameters is a paid mutator transaction binding the contract method 0x9dcaa6c9.
//
// Solidity: function setDelegationParameters(uint32 _indexingRewardCut, uint32 _queryFeeCut, uint32 _cooldownBlocks) returns()
func (_Staking *StakingTransactorSession) SetDelegationParameters(_indexingRewardCut uint32, _queryFeeCut uint32, _cooldownBlocks uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationParameters(&_Staking.TransactOpts, _indexingRewardCut, _queryFeeCut, _cooldownBlocks)
}

// SetDelegationParametersCooldown is a paid mutator transaction binding the contract method 0xf8b80a92.
//
// Solidity: function setDelegationParametersCooldown(uint32 _blocks) returns()
func (_Staking *StakingTransactor) SetDelegationParametersCooldown(opts *bind.TransactOpts, _blocks uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setDelegationParametersCooldown", _blocks)
}

// SetDelegationParametersCooldown is a paid mutator transaction binding the contract method 0xf8b80a92.
//
// Solidity: function setDelegationParametersCooldown(uint32 _blocks) returns()
func (_Staking *StakingSession) SetDelegationParametersCooldown(_blocks uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationParametersCooldown(&_Staking.TransactOpts, _blocks)
}

// SetDelegationParametersCooldown is a paid mutator transaction binding the contract method 0xf8b80a92.
//
// Solidity: function setDelegationParametersCooldown(uint32 _blocks) returns()
func (_Staking *StakingTransactorSession) SetDelegationParametersCooldown(_blocks uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationParametersCooldown(&_Staking.TransactOpts, _blocks)
}

// SetDelegationRatio is a paid mutator transaction binding the contract method 0x1dd42f60.
//
// Solidity: function setDelegationRatio(uint32 _delegationRatio) returns()
func (_Staking *StakingTransactor) SetDelegationRatio(opts *bind.TransactOpts, _delegationRatio uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setDelegationRatio", _delegationRatio)
}

// SetDelegationRatio is a paid mutator transaction binding the contract method 0x1dd42f60.
//
// Solidity: function setDelegationRatio(uint32 _delegationRatio) returns()
func (_Staking *StakingSession) SetDelegationRatio(_delegationRatio uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationRatio(&_Staking.TransactOpts, _delegationRatio)
}

// SetDelegationRatio is a paid mutator transaction binding the contract method 0x1dd42f60.
//
// Solidity: function setDelegationRatio(uint32 _delegationRatio) returns()
func (_Staking *StakingTransactorSession) SetDelegationRatio(_delegationRatio uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationRatio(&_Staking.TransactOpts, _delegationRatio)
}

// SetDelegationTaxPercentage is a paid mutator transaction binding the contract method 0xe6dc5a1c.
//
// Solidity: function setDelegationTaxPercentage(uint32 _percentage) returns()
func (_Staking *StakingTransactor) SetDelegationTaxPercentage(opts *bind.TransactOpts, _percentage uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setDelegationTaxPercentage", _percentage)
}

// SetDelegationTaxPercentage is a paid mutator transaction binding the contract method 0xe6dc5a1c.
//
// Solidity: function setDelegationTaxPercentage(uint32 _percentage) returns()
func (_Staking *StakingSession) SetDelegationTaxPercentage(_percentage uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationTaxPercentage(&_Staking.TransactOpts, _percentage)
}

// SetDelegationTaxPercentage is a paid mutator transaction binding the contract method 0xe6dc5a1c.
//
// Solidity: function setDelegationTaxPercentage(uint32 _percentage) returns()
func (_Staking *StakingTransactorSession) SetDelegationTaxPercentage(_percentage uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationTaxPercentage(&_Staking.TransactOpts, _percentage)
}

// SetDelegationUnbondingPeriod is a paid mutator transaction binding the contract method 0x5e9a6392.
//
// Solidity: function setDelegationUnbondingPeriod(uint32 _delegationUnbondingPeriod) returns()
func (_Staking *StakingTransactor) SetDelegationUnbondingPeriod(opts *bind.TransactOpts, _delegationUnbondingPeriod uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setDelegationUnbondingPeriod", _delegationUnbondingPeriod)
}

// SetDelegationUnbondingPeriod is a paid mutator transaction binding the contract method 0x5e9a6392.
//
// Solidity: function setDelegationUnbondingPeriod(uint32 _delegationUnbondingPeriod) returns()
func (_Staking *StakingSession) SetDelegationUnbondingPeriod(_delegationUnbondingPeriod uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationUnbondingPeriod(&_Staking.TransactOpts, _delegationUnbondingPeriod)
}

// SetDelegationUnbondingPeriod is a paid mutator transaction binding the contract method 0x5e9a6392.
//
// Solidity: function setDelegationUnbondingPeriod(uint32 _delegationUnbondingPeriod) returns()
func (_Staking *StakingTransactorSession) SetDelegationUnbondingPeriod(_delegationUnbondingPeriod uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetDelegationUnbondingPeriod(&_Staking.TransactOpts, _delegationUnbondingPeriod)
}

// SetMaxAllocationEpochs is a paid mutator transaction binding the contract method 0x2652d75e.
//
// Solidity: function setMaxAllocationEpochs(uint32 _maxAllocationEpochs) returns()
func (_Staking *StakingTransactor) SetMaxAllocationEpochs(opts *bind.TransactOpts, _maxAllocationEpochs uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMaxAllocationEpochs", _maxAllocationEpochs)
}

// SetMaxAllocationEpochs is a paid mutator transaction binding the contract method 0x2652d75e.
//
// Solidity: function setMaxAllocationEpochs(uint32 _maxAllocationEpochs) returns()
func (_Staking *StakingSession) SetMaxAllocationEpochs(_maxAllocationEpochs uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxAllocationEpochs(&_Staking.TransactOpts, _maxAllocationEpochs)
}

// SetMaxAllocationEpochs is a paid mutator transaction binding the contract method 0x2652d75e.
//
// Solidity: function setMaxAllocationEpochs(uint32 _maxAllocationEpochs) returns()
func (_Staking *StakingTransactorSession) SetMaxAllocationEpochs(_maxAllocationEpochs uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxAllocationEpochs(&_Staking.TransactOpts, _maxAllocationEpochs)
}

// SetMinimumIndexerStake is a paid mutator transaction binding the contract method 0xddb8b131.
//
// Solidity: function setMinimumIndexerStake(uint256 _minimumIndexerStake) returns()
func (_Staking *StakingTransactor) SetMinimumIndexerStake(opts *bind.TransactOpts, _minimumIndexerStake *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMinimumIndexerStake", _minimumIndexerStake)
}

// SetMinimumIndexerStake is a paid mutator transaction binding the contract method 0xddb8b131.
//
// Solidity: function setMinimumIndexerStake(uint256 _minimumIndexerStake) returns()
func (_Staking *StakingSession) SetMinimumIndexerStake(_minimumIndexerStake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinimumIndexerStake(&_Staking.TransactOpts, _minimumIndexerStake)
}

// SetMinimumIndexerStake is a paid mutator transaction binding the contract method 0xddb8b131.
//
// Solidity: function setMinimumIndexerStake(uint256 _minimumIndexerStake) returns()
func (_Staking *StakingTransactorSession) SetMinimumIndexerStake(_minimumIndexerStake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinimumIndexerStake(&_Staking.TransactOpts, _minimumIndexerStake)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_Staking *StakingTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setOperator", _operator, _allowed)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_Staking *StakingSession) SetOperator(_operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.Contract.SetOperator(&_Staking.TransactOpts, _operator, _allowed)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_Staking *StakingTransactorSession) SetOperator(_operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.Contract.SetOperator(&_Staking.TransactOpts, _operator, _allowed)
}

// SetProtocolPercentage is a paid mutator transaction binding the contract method 0x9a48bf83.
//
// Solidity: function setProtocolPercentage(uint32 _percentage) returns()
func (_Staking *StakingTransactor) SetProtocolPercentage(opts *bind.TransactOpts, _percentage uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setProtocolPercentage", _percentage)
}

// SetProtocolPercentage is a paid mutator transaction binding the contract method 0x9a48bf83.
//
// Solidity: function setProtocolPercentage(uint32 _percentage) returns()
func (_Staking *StakingSession) SetProtocolPercentage(_percentage uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetProtocolPercentage(&_Staking.TransactOpts, _percentage)
}

// SetProtocolPercentage is a paid mutator transaction binding the contract method 0x9a48bf83.
//
// Solidity: function setProtocolPercentage(uint32 _percentage) returns()
func (_Staking *StakingTransactorSession) SetProtocolPercentage(_percentage uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetProtocolPercentage(&_Staking.TransactOpts, _percentage)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x979f9f5c.
//
// Solidity: function setRebateRatio(uint32 _alphaNumerator, uint32 _alphaDenominator) returns()
func (_Staking *StakingTransactor) SetRebateRatio(opts *bind.TransactOpts, _alphaNumerator uint32, _alphaDenominator uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setRebateRatio", _alphaNumerator, _alphaDenominator)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x979f9f5c.
//
// Solidity: function setRebateRatio(uint32 _alphaNumerator, uint32 _alphaDenominator) returns()
func (_Staking *StakingSession) SetRebateRatio(_alphaNumerator uint32, _alphaDenominator uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetRebateRatio(&_Staking.TransactOpts, _alphaNumerator, _alphaDenominator)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x979f9f5c.
//
// Solidity: function setRebateRatio(uint32 _alphaNumerator, uint32 _alphaDenominator) returns()
func (_Staking *StakingTransactorSession) SetRebateRatio(_alphaNumerator uint32, _alphaDenominator uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetRebateRatio(&_Staking.TransactOpts, _alphaNumerator, _alphaDenominator)
}

// SetRewardsDestination is a paid mutator transaction binding the contract method 0x772495c3.
//
// Solidity: function setRewardsDestination(address _destination) returns()
func (_Staking *StakingTransactor) SetRewardsDestination(opts *bind.TransactOpts, _destination common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setRewardsDestination", _destination)
}

// SetRewardsDestination is a paid mutator transaction binding the contract method 0x772495c3.
//
// Solidity: function setRewardsDestination(address _destination) returns()
func (_Staking *StakingSession) SetRewardsDestination(_destination common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetRewardsDestination(&_Staking.TransactOpts, _destination)
}

// SetRewardsDestination is a paid mutator transaction binding the contract method 0x772495c3.
//
// Solidity: function setRewardsDestination(address _destination) returns()
func (_Staking *StakingTransactorSession) SetRewardsDestination(_destination common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetRewardsDestination(&_Staking.TransactOpts, _destination)
}

// SetSlasher is a paid mutator transaction binding the contract method 0x52348080.
//
// Solidity: function setSlasher(address _slasher, bool _allowed) returns()
func (_Staking *StakingTransactor) SetSlasher(opts *bind.TransactOpts, _slasher common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setSlasher", _slasher, _allowed)
}

// SetSlasher is a paid mutator transaction binding the contract method 0x52348080.
//
// Solidity: function setSlasher(address _slasher, bool _allowed) returns()
func (_Staking *StakingSession) SetSlasher(_slasher common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.Contract.SetSlasher(&_Staking.TransactOpts, _slasher, _allowed)
}

// SetSlasher is a paid mutator transaction binding the contract method 0x52348080.
//
// Solidity: function setSlasher(address _slasher, bool _allowed) returns()
func (_Staking *StakingTransactorSession) SetSlasher(_slasher common.Address, _allowed bool) (*types.Transaction, error) {
	return _Staking.Contract.SetSlasher(&_Staking.TransactOpts, _slasher, _allowed)
}

// SetThawingPeriod is a paid mutator transaction binding the contract method 0x32bc9108.
//
// Solidity: function setThawingPeriod(uint32 _thawingPeriod) returns()
func (_Staking *StakingTransactor) SetThawingPeriod(opts *bind.TransactOpts, _thawingPeriod uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setThawingPeriod", _thawingPeriod)
}

// SetThawingPeriod is a paid mutator transaction binding the contract method 0x32bc9108.
//
// Solidity: function setThawingPeriod(uint32 _thawingPeriod) returns()
func (_Staking *StakingSession) SetThawingPeriod(_thawingPeriod uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetThawingPeriod(&_Staking.TransactOpts, _thawingPeriod)
}

// SetThawingPeriod is a paid mutator transaction binding the contract method 0x32bc9108.
//
// Solidity: function setThawingPeriod(uint32 _thawingPeriod) returns()
func (_Staking *StakingTransactorSession) SetThawingPeriod(_thawingPeriod uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetThawingPeriod(&_Staking.TransactOpts, _thawingPeriod)
}

// Slash is a paid mutator transaction binding the contract method 0xe76fede6.
//
// Solidity: function slash(address _indexer, uint256 _tokens, uint256 _reward, address _beneficiary) returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts, _indexer common.Address, _tokens *big.Int, _reward *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash", _indexer, _tokens, _reward, _beneficiary)
}

// Slash is a paid mutator transaction binding the contract method 0xe76fede6.
//
// Solidity: function slash(address _indexer, uint256 _tokens, uint256 _reward, address _beneficiary) returns()
func (_Staking *StakingSession) Slash(_indexer common.Address, _tokens *big.Int, _reward *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _indexer, _tokens, _reward, _beneficiary)
}

// Slash is a paid mutator transaction binding the contract method 0xe76fede6.
//
// Solidity: function slash(address _indexer, uint256 _tokens, uint256 _reward, address _beneficiary) returns()
func (_Staking *StakingTransactorSession) Slash(_indexer common.Address, _tokens *big.Int, _reward *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _indexer, _tokens, _reward, _beneficiary)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokens) returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", _tokens)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokens) returns()
func (_Staking *StakingSession) Stake(_tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, _tokens)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokens) returns()
func (_Staking *StakingTransactorSession) Stake(_tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, _tokens)
}

// StakeTo is a paid mutator transaction binding the contract method 0xa2a31722.
//
// Solidity: function stakeTo(address _indexer, uint256 _tokens) returns()
func (_Staking *StakingTransactor) StakeTo(opts *bind.TransactOpts, _indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeTo", _indexer, _tokens)
}

// StakeTo is a paid mutator transaction binding the contract method 0xa2a31722.
//
// Solidity: function stakeTo(address _indexer, uint256 _tokens) returns()
func (_Staking *StakingSession) StakeTo(_indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeTo(&_Staking.TransactOpts, _indexer, _tokens)
}

// StakeTo is a paid mutator transaction binding the contract method 0xa2a31722.
//
// Solidity: function stakeTo(address _indexer, uint256 _tokens) returns()
func (_Staking *StakingTransactorSession) StakeTo(_indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeTo(&_Staking.TransactOpts, _indexer, _tokens)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _indexer, uint256 _shares) returns(uint256)
func (_Staking *StakingTransactor) Undelegate(opts *bind.TransactOpts, _indexer common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegate", _indexer, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _indexer, uint256 _shares) returns(uint256)
func (_Staking *StakingSession) Undelegate(_indexer common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, _indexer, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _indexer, uint256 _shares) returns(uint256)
func (_Staking *StakingTransactorSession) Undelegate(_indexer common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, _indexer, _shares)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake", _tokens)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_Staking *StakingSession) Unstake(_tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, _tokens)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_Staking *StakingTransactorSession) Unstake(_tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, _tokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingSession) Withdraw() (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts)
}

// WithdrawDelegated is a paid mutator transaction binding the contract method 0x51a60b02.
//
// Solidity: function withdrawDelegated(address _indexer, address _delegateToIndexer) returns(uint256)
func (_Staking *StakingTransactor) WithdrawDelegated(opts *bind.TransactOpts, _indexer common.Address, _delegateToIndexer common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawDelegated", _indexer, _delegateToIndexer)
}

// WithdrawDelegated is a paid mutator transaction binding the contract method 0x51a60b02.
//
// Solidity: function withdrawDelegated(address _indexer, address _delegateToIndexer) returns(uint256)
func (_Staking *StakingSession) WithdrawDelegated(_indexer common.Address, _delegateToIndexer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawDelegated(&_Staking.TransactOpts, _indexer, _delegateToIndexer)
}

// WithdrawDelegated is a paid mutator transaction binding the contract method 0x51a60b02.
//
// Solidity: function withdrawDelegated(address _indexer, address _delegateToIndexer) returns(uint256)
func (_Staking *StakingTransactorSession) WithdrawDelegated(_indexer common.Address, _delegateToIndexer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawDelegated(&_Staking.TransactOpts, _indexer, _delegateToIndexer)
}

// StakingAllocationClosedIterator is returned from FilterAllocationClosed and is used to iterate over the raw logs and unpacked data for AllocationClosed events raised by the Staking contract.
type StakingAllocationClosedIterator struct {
	Event *StakingAllocationClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingAllocationClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAllocationClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingAllocationClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingAllocationClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAllocationClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAllocationClosed represents a AllocationClosed event raised by the Staking contract.
type StakingAllocationClosed struct {
	Indexer              common.Address
	SubgraphDeploymentID [32]byte
	Epoch                *big.Int
	Tokens               *big.Int
	AllocationID         common.Address
	EffectiveAllocation  *big.Int
	Sender               common.Address
	Poi                  [32]byte
	IsDelegator          bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAllocationClosed is a free log retrieval operation binding the contract event 0x7203ffa6902c4c2a85ac2612321460fa20e29a972c272ecedfdf95f944616269.
//
// Solidity: event AllocationClosed(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, uint256 effectiveAllocation, address sender, bytes32 poi, bool isDelegator)
func (_Staking *StakingFilterer) FilterAllocationClosed(opts *bind.FilterOpts, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (*StakingAllocationClosedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AllocationClosed", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &StakingAllocationClosedIterator{contract: _Staking.contract, event: "AllocationClosed", logs: logs, sub: sub}, nil
}

// WatchAllocationClosed is a free log subscription operation binding the contract event 0x7203ffa6902c4c2a85ac2612321460fa20e29a972c272ecedfdf95f944616269.
//
// Solidity: event AllocationClosed(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, uint256 effectiveAllocation, address sender, bytes32 poi, bool isDelegator)
func (_Staking *StakingFilterer) WatchAllocationClosed(opts *bind.WatchOpts, sink chan<- *StakingAllocationClosed, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AllocationClosed", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAllocationClosed)
				if err := _Staking.contract.UnpackLog(event, "AllocationClosed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllocationClosed is a log parse operation binding the contract event 0x7203ffa6902c4c2a85ac2612321460fa20e29a972c272ecedfdf95f944616269.
//
// Solidity: event AllocationClosed(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, uint256 effectiveAllocation, address sender, bytes32 poi, bool isDelegator)
func (_Staking *StakingFilterer) ParseAllocationClosed(log types.Log) (*StakingAllocationClosed, error) {
	event := new(StakingAllocationClosed)
	if err := _Staking.contract.UnpackLog(event, "AllocationClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingAllocationCollectedIterator is returned from FilterAllocationCollected and is used to iterate over the raw logs and unpacked data for AllocationCollected events raised by the Staking contract.
type StakingAllocationCollectedIterator struct {
	Event *StakingAllocationCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingAllocationCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAllocationCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingAllocationCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingAllocationCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAllocationCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAllocationCollected represents a AllocationCollected event raised by the Staking contract.
type StakingAllocationCollected struct {
	Indexer              common.Address
	SubgraphDeploymentID [32]byte
	Epoch                *big.Int
	Tokens               *big.Int
	AllocationID         common.Address
	From                 common.Address
	CurationFees         *big.Int
	RebateFees           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAllocationCollected is a free log retrieval operation binding the contract event 0x18040f6f54270f646d21bc8e963105c53499cbcebe6f2a5b32c7018e18a3451e.
//
// Solidity: event AllocationCollected(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, address from, uint256 curationFees, uint256 rebateFees)
func (_Staking *StakingFilterer) FilterAllocationCollected(opts *bind.FilterOpts, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (*StakingAllocationCollectedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AllocationCollected", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &StakingAllocationCollectedIterator{contract: _Staking.contract, event: "AllocationCollected", logs: logs, sub: sub}, nil
}

// WatchAllocationCollected is a free log subscription operation binding the contract event 0x18040f6f54270f646d21bc8e963105c53499cbcebe6f2a5b32c7018e18a3451e.
//
// Solidity: event AllocationCollected(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, address from, uint256 curationFees, uint256 rebateFees)
func (_Staking *StakingFilterer) WatchAllocationCollected(opts *bind.WatchOpts, sink chan<- *StakingAllocationCollected, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AllocationCollected", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAllocationCollected)
				if err := _Staking.contract.UnpackLog(event, "AllocationCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllocationCollected is a log parse operation binding the contract event 0x18040f6f54270f646d21bc8e963105c53499cbcebe6f2a5b32c7018e18a3451e.
//
// Solidity: event AllocationCollected(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, address from, uint256 curationFees, uint256 rebateFees)
func (_Staking *StakingFilterer) ParseAllocationCollected(log types.Log) (*StakingAllocationCollected, error) {
	event := new(StakingAllocationCollected)
	if err := _Staking.contract.UnpackLog(event, "AllocationCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingAllocationCreatedIterator is returned from FilterAllocationCreated and is used to iterate over the raw logs and unpacked data for AllocationCreated events raised by the Staking contract.
type StakingAllocationCreatedIterator struct {
	Event *StakingAllocationCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingAllocationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAllocationCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingAllocationCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingAllocationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAllocationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAllocationCreated represents a AllocationCreated event raised by the Staking contract.
type StakingAllocationCreated struct {
	Indexer              common.Address
	SubgraphDeploymentID [32]byte
	Epoch                *big.Int
	Tokens               *big.Int
	AllocationID         common.Address
	Metadata             [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAllocationCreated is a free log retrieval operation binding the contract event 0x0f73ab5f706106366951b51f760e0a6f60c794f233d90958d81c82ad84fa6e87.
//
// Solidity: event AllocationCreated(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, bytes32 metadata)
func (_Staking *StakingFilterer) FilterAllocationCreated(opts *bind.FilterOpts, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (*StakingAllocationCreatedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AllocationCreated", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &StakingAllocationCreatedIterator{contract: _Staking.contract, event: "AllocationCreated", logs: logs, sub: sub}, nil
}

// WatchAllocationCreated is a free log subscription operation binding the contract event 0x0f73ab5f706106366951b51f760e0a6f60c794f233d90958d81c82ad84fa6e87.
//
// Solidity: event AllocationCreated(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, bytes32 metadata)
func (_Staking *StakingFilterer) WatchAllocationCreated(opts *bind.WatchOpts, sink chan<- *StakingAllocationCreated, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AllocationCreated", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAllocationCreated)
				if err := _Staking.contract.UnpackLog(event, "AllocationCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllocationCreated is a log parse operation binding the contract event 0x0f73ab5f706106366951b51f760e0a6f60c794f233d90958d81c82ad84fa6e87.
//
// Solidity: event AllocationCreated(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, bytes32 metadata)
func (_Staking *StakingFilterer) ParseAllocationCreated(log types.Log) (*StakingAllocationCreated, error) {
	event := new(StakingAllocationCreated)
	if err := _Staking.contract.UnpackLog(event, "AllocationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingAssetHolderUpdateIterator is returned from FilterAssetHolderUpdate and is used to iterate over the raw logs and unpacked data for AssetHolderUpdate events raised by the Staking contract.
type StakingAssetHolderUpdateIterator struct {
	Event *StakingAssetHolderUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingAssetHolderUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAssetHolderUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingAssetHolderUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingAssetHolderUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAssetHolderUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAssetHolderUpdate represents a AssetHolderUpdate event raised by the Staking contract.
type StakingAssetHolderUpdate struct {
	Caller      common.Address
	AssetHolder common.Address
	Allowed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetHolderUpdate is a free log retrieval operation binding the contract event 0x3a8d5e92bb089ebd158e2c22dc449009d62b0df02b6a6792bb0c5fc33f240fcb.
//
// Solidity: event AssetHolderUpdate(address indexed caller, address indexed assetHolder, bool allowed)
func (_Staking *StakingFilterer) FilterAssetHolderUpdate(opts *bind.FilterOpts, caller []common.Address, assetHolder []common.Address) (*StakingAssetHolderUpdateIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetHolderRule []interface{}
	for _, assetHolderItem := range assetHolder {
		assetHolderRule = append(assetHolderRule, assetHolderItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AssetHolderUpdate", callerRule, assetHolderRule)
	if err != nil {
		return nil, err
	}
	return &StakingAssetHolderUpdateIterator{contract: _Staking.contract, event: "AssetHolderUpdate", logs: logs, sub: sub}, nil
}

// WatchAssetHolderUpdate is a free log subscription operation binding the contract event 0x3a8d5e92bb089ebd158e2c22dc449009d62b0df02b6a6792bb0c5fc33f240fcb.
//
// Solidity: event AssetHolderUpdate(address indexed caller, address indexed assetHolder, bool allowed)
func (_Staking *StakingFilterer) WatchAssetHolderUpdate(opts *bind.WatchOpts, sink chan<- *StakingAssetHolderUpdate, caller []common.Address, assetHolder []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetHolderRule []interface{}
	for _, assetHolderItem := range assetHolder {
		assetHolderRule = append(assetHolderRule, assetHolderItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AssetHolderUpdate", callerRule, assetHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAssetHolderUpdate)
				if err := _Staking.contract.UnpackLog(event, "AssetHolderUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetHolderUpdate is a log parse operation binding the contract event 0x3a8d5e92bb089ebd158e2c22dc449009d62b0df02b6a6792bb0c5fc33f240fcb.
//
// Solidity: event AssetHolderUpdate(address indexed caller, address indexed assetHolder, bool allowed)
func (_Staking *StakingFilterer) ParseAssetHolderUpdate(log types.Log) (*StakingAssetHolderUpdate, error) {
	event := new(StakingAssetHolderUpdate)
	if err := _Staking.contract.UnpackLog(event, "AssetHolderUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegationParametersUpdatedIterator is returned from FilterDelegationParametersUpdated and is used to iterate over the raw logs and unpacked data for DelegationParametersUpdated events raised by the Staking contract.
type StakingDelegationParametersUpdatedIterator struct {
	Event *StakingDelegationParametersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingDelegationParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegationParametersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingDelegationParametersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingDelegationParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegationParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegationParametersUpdated represents a DelegationParametersUpdated event raised by the Staking contract.
type StakingDelegationParametersUpdated struct {
	Indexer           common.Address
	IndexingRewardCut uint32
	QueryFeeCut       uint32
	CooldownBlocks    uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelegationParametersUpdated is a free log retrieval operation binding the contract event 0xdd5c1add84431df7ff63c721510522fbccafda37dfc33f0f5094d90135a8f22a.
//
// Solidity: event DelegationParametersUpdated(address indexed indexer, uint32 indexingRewardCut, uint32 queryFeeCut, uint32 cooldownBlocks)
func (_Staking *StakingFilterer) FilterDelegationParametersUpdated(opts *bind.FilterOpts, indexer []common.Address) (*StakingDelegationParametersUpdatedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegationParametersUpdated", indexerRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegationParametersUpdatedIterator{contract: _Staking.contract, event: "DelegationParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationParametersUpdated is a free log subscription operation binding the contract event 0xdd5c1add84431df7ff63c721510522fbccafda37dfc33f0f5094d90135a8f22a.
//
// Solidity: event DelegationParametersUpdated(address indexed indexer, uint32 indexingRewardCut, uint32 queryFeeCut, uint32 cooldownBlocks)
func (_Staking *StakingFilterer) WatchDelegationParametersUpdated(opts *bind.WatchOpts, sink chan<- *StakingDelegationParametersUpdated, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegationParametersUpdated", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegationParametersUpdated)
				if err := _Staking.contract.UnpackLog(event, "DelegationParametersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegationParametersUpdated is a log parse operation binding the contract event 0xdd5c1add84431df7ff63c721510522fbccafda37dfc33f0f5094d90135a8f22a.
//
// Solidity: event DelegationParametersUpdated(address indexed indexer, uint32 indexingRewardCut, uint32 queryFeeCut, uint32 cooldownBlocks)
func (_Staking *StakingFilterer) ParseDelegationParametersUpdated(log types.Log) (*StakingDelegationParametersUpdated, error) {
	event := new(StakingDelegationParametersUpdated)
	if err := _Staking.contract.UnpackLog(event, "DelegationParametersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingParameterUpdatedIterator is returned from FilterParameterUpdated and is used to iterate over the raw logs and unpacked data for ParameterUpdated events raised by the Staking contract.
type StakingParameterUpdatedIterator struct {
	Event *StakingParameterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingParameterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingParameterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingParameterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingParameterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingParameterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingParameterUpdated represents a ParameterUpdated event raised by the Staking contract.
type StakingParameterUpdated struct {
	Param string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParameterUpdated is a free log retrieval operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_Staking *StakingFilterer) FilterParameterUpdated(opts *bind.FilterOpts) (*StakingParameterUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ParameterUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingParameterUpdatedIterator{contract: _Staking.contract, event: "ParameterUpdated", logs: logs, sub: sub}, nil
}

// WatchParameterUpdated is a free log subscription operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_Staking *StakingFilterer) WatchParameterUpdated(opts *bind.WatchOpts, sink chan<- *StakingParameterUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ParameterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingParameterUpdated)
				if err := _Staking.contract.UnpackLog(event, "ParameterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParameterUpdated is a log parse operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_Staking *StakingFilterer) ParseParameterUpdated(log types.Log) (*StakingParameterUpdated, error) {
	event := new(StakingParameterUpdated)
	if err := _Staking.contract.UnpackLog(event, "ParameterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRebateClaimedIterator is returned from FilterRebateClaimed and is used to iterate over the raw logs and unpacked data for RebateClaimed events raised by the Staking contract.
type StakingRebateClaimedIterator struct {
	Event *StakingRebateClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingRebateClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRebateClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingRebateClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingRebateClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRebateClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRebateClaimed represents a RebateClaimed event raised by the Staking contract.
type StakingRebateClaimed struct {
	Indexer                   common.Address
	SubgraphDeploymentID      [32]byte
	AllocationID              common.Address
	Epoch                     *big.Int
	ForEpoch                  *big.Int
	Tokens                    *big.Int
	UnclaimedAllocationsCount *big.Int
	DelegationFees            *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterRebateClaimed is a free log retrieval operation binding the contract event 0xb5f11a762db39abff5529064f3103b1abb9a5a3ba3d61972c1a7006d09db7d20.
//
// Solidity: event RebateClaimed(address indexed indexer, bytes32 indexed subgraphDeploymentID, address indexed allocationID, uint256 epoch, uint256 forEpoch, uint256 tokens, uint256 unclaimedAllocationsCount, uint256 delegationFees)
func (_Staking *StakingFilterer) FilterRebateClaimed(opts *bind.FilterOpts, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (*StakingRebateClaimedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}
	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RebateClaimed", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &StakingRebateClaimedIterator{contract: _Staking.contract, event: "RebateClaimed", logs: logs, sub: sub}, nil
}

// WatchRebateClaimed is a free log subscription operation binding the contract event 0xb5f11a762db39abff5529064f3103b1abb9a5a3ba3d61972c1a7006d09db7d20.
//
// Solidity: event RebateClaimed(address indexed indexer, bytes32 indexed subgraphDeploymentID, address indexed allocationID, uint256 epoch, uint256 forEpoch, uint256 tokens, uint256 unclaimedAllocationsCount, uint256 delegationFees)
func (_Staking *StakingFilterer) WatchRebateClaimed(opts *bind.WatchOpts, sink chan<- *StakingRebateClaimed, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}
	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RebateClaimed", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRebateClaimed)
				if err := _Staking.contract.UnpackLog(event, "RebateClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRebateClaimed is a log parse operation binding the contract event 0xb5f11a762db39abff5529064f3103b1abb9a5a3ba3d61972c1a7006d09db7d20.
//
// Solidity: event RebateClaimed(address indexed indexer, bytes32 indexed subgraphDeploymentID, address indexed allocationID, uint256 epoch, uint256 forEpoch, uint256 tokens, uint256 unclaimedAllocationsCount, uint256 delegationFees)
func (_Staking *StakingFilterer) ParseRebateClaimed(log types.Log) (*StakingRebateClaimed, error) {
	event := new(StakingRebateClaimed)
	if err := _Staking.contract.UnpackLog(event, "RebateClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSetControllerIterator is returned from FilterSetController and is used to iterate over the raw logs and unpacked data for SetController events raised by the Staking contract.
type StakingSetControllerIterator struct {
	Event *StakingSetController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingSetControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSetController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingSetController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingSetControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSetControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSetController represents a SetController event raised by the Staking contract.
type StakingSetController struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetController is a free log retrieval operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_Staking *StakingFilterer) FilterSetController(opts *bind.FilterOpts) (*StakingSetControllerIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SetController")
	if err != nil {
		return nil, err
	}
	return &StakingSetControllerIterator{contract: _Staking.contract, event: "SetController", logs: logs, sub: sub}, nil
}

// WatchSetController is a free log subscription operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_Staking *StakingFilterer) WatchSetController(opts *bind.WatchOpts, sink chan<- *StakingSetController) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SetController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSetController)
				if err := _Staking.contract.UnpackLog(event, "SetController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetController is a log parse operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_Staking *StakingFilterer) ParseSetController(log types.Log) (*StakingSetController, error) {
	event := new(StakingSetController)
	if err := _Staking.contract.UnpackLog(event, "SetController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSetOperatorIterator is returned from FilterSetOperator and is used to iterate over the raw logs and unpacked data for SetOperator events raised by the Staking contract.
type StakingSetOperatorIterator struct {
	Event *StakingSetOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingSetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSetOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingSetOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingSetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSetOperator represents a SetOperator event raised by the Staking contract.
type StakingSetOperator struct {
	Indexer  common.Address
	Operator common.Address
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOperator is a free log retrieval operation binding the contract event 0xa3581229e2c315eb01303f468621e07aa9b628a23b1608162ae063f143355135.
//
// Solidity: event SetOperator(address indexed indexer, address indexed operator, bool allowed)
func (_Staking *StakingFilterer) FilterSetOperator(opts *bind.FilterOpts, indexer []common.Address, operator []common.Address) (*StakingSetOperatorIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SetOperator", indexerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingSetOperatorIterator{contract: _Staking.contract, event: "SetOperator", logs: logs, sub: sub}, nil
}

// WatchSetOperator is a free log subscription operation binding the contract event 0xa3581229e2c315eb01303f468621e07aa9b628a23b1608162ae063f143355135.
//
// Solidity: event SetOperator(address indexed indexer, address indexed operator, bool allowed)
func (_Staking *StakingFilterer) WatchSetOperator(opts *bind.WatchOpts, sink chan<- *StakingSetOperator, indexer []common.Address, operator []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SetOperator", indexerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSetOperator)
				if err := _Staking.contract.UnpackLog(event, "SetOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetOperator is a log parse operation binding the contract event 0xa3581229e2c315eb01303f468621e07aa9b628a23b1608162ae063f143355135.
//
// Solidity: event SetOperator(address indexed indexer, address indexed operator, bool allowed)
func (_Staking *StakingFilterer) ParseSetOperator(log types.Log) (*StakingSetOperator, error) {
	event := new(StakingSetOperator)
	if err := _Staking.contract.UnpackLog(event, "SetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSetRewardsDestinationIterator is returned from FilterSetRewardsDestination and is used to iterate over the raw logs and unpacked data for SetRewardsDestination events raised by the Staking contract.
type StakingSetRewardsDestinationIterator struct {
	Event *StakingSetRewardsDestination // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingSetRewardsDestinationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSetRewardsDestination)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingSetRewardsDestination)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingSetRewardsDestinationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSetRewardsDestinationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSetRewardsDestination represents a SetRewardsDestination event raised by the Staking contract.
type StakingSetRewardsDestination struct {
	Indexer     common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetRewardsDestination is a free log retrieval operation binding the contract event 0x29c33cd533c17d8916c8e471a4e2c4d1e34caa9b8844527c0bb182b3c104c7d3.
//
// Solidity: event SetRewardsDestination(address indexed indexer, address indexed destination)
func (_Staking *StakingFilterer) FilterSetRewardsDestination(opts *bind.FilterOpts, indexer []common.Address, destination []common.Address) (*StakingSetRewardsDestinationIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SetRewardsDestination", indexerRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &StakingSetRewardsDestinationIterator{contract: _Staking.contract, event: "SetRewardsDestination", logs: logs, sub: sub}, nil
}

// WatchSetRewardsDestination is a free log subscription operation binding the contract event 0x29c33cd533c17d8916c8e471a4e2c4d1e34caa9b8844527c0bb182b3c104c7d3.
//
// Solidity: event SetRewardsDestination(address indexed indexer, address indexed destination)
func (_Staking *StakingFilterer) WatchSetRewardsDestination(opts *bind.WatchOpts, sink chan<- *StakingSetRewardsDestination, indexer []common.Address, destination []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SetRewardsDestination", indexerRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSetRewardsDestination)
				if err := _Staking.contract.UnpackLog(event, "SetRewardsDestination", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetRewardsDestination is a log parse operation binding the contract event 0x29c33cd533c17d8916c8e471a4e2c4d1e34caa9b8844527c0bb182b3c104c7d3.
//
// Solidity: event SetRewardsDestination(address indexed indexer, address indexed destination)
func (_Staking *StakingFilterer) ParseSetRewardsDestination(log types.Log) (*StakingSetRewardsDestination, error) {
	event := new(StakingSetRewardsDestination)
	if err := _Staking.contract.UnpackLog(event, "SetRewardsDestination", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlasherUpdateIterator is returned from FilterSlasherUpdate and is used to iterate over the raw logs and unpacked data for SlasherUpdate events raised by the Staking contract.
type StakingSlasherUpdateIterator struct {
	Event *StakingSlasherUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingSlasherUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlasherUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingSlasherUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingSlasherUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlasherUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlasherUpdate represents a SlasherUpdate event raised by the Staking contract.
type StakingSlasherUpdate struct {
	Caller  common.Address
	Slasher common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlasherUpdate is a free log retrieval operation binding the contract event 0x87ea6771e87d96ce16dbe8eda64da9473733e4c1c568baf8ae47256c5bd765e9.
//
// Solidity: event SlasherUpdate(address indexed caller, address indexed slasher, bool allowed)
func (_Staking *StakingFilterer) FilterSlasherUpdate(opts *bind.FilterOpts, caller []common.Address, slasher []common.Address) (*StakingSlasherUpdateIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SlasherUpdate", callerRule, slasherRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlasherUpdateIterator{contract: _Staking.contract, event: "SlasherUpdate", logs: logs, sub: sub}, nil
}

// WatchSlasherUpdate is a free log subscription operation binding the contract event 0x87ea6771e87d96ce16dbe8eda64da9473733e4c1c568baf8ae47256c5bd765e9.
//
// Solidity: event SlasherUpdate(address indexed caller, address indexed slasher, bool allowed)
func (_Staking *StakingFilterer) WatchSlasherUpdate(opts *bind.WatchOpts, sink chan<- *StakingSlasherUpdate, caller []common.Address, slasher []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SlasherUpdate", callerRule, slasherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlasherUpdate)
				if err := _Staking.contract.UnpackLog(event, "SlasherUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlasherUpdate is a log parse operation binding the contract event 0x87ea6771e87d96ce16dbe8eda64da9473733e4c1c568baf8ae47256c5bd765e9.
//
// Solidity: event SlasherUpdate(address indexed caller, address indexed slasher, bool allowed)
func (_Staking *StakingFilterer) ParseSlasherUpdate(log types.Log) (*StakingSlasherUpdate, error) {
	event := new(StakingSlasherUpdate)
	if err := _Staking.contract.UnpackLog(event, "SlasherUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeDelegatedIterator is returned from FilterStakeDelegated and is used to iterate over the raw logs and unpacked data for StakeDelegated events raised by the Staking contract.
type StakingStakeDelegatedIterator struct {
	Event *StakingStakeDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakeDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStakeDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakeDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeDelegated represents a StakeDelegated event raised by the Staking contract.
type StakingStakeDelegated struct {
	Indexer   common.Address
	Delegator common.Address
	Tokens    *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeDelegated is a free log retrieval operation binding the contract event 0xcd0366dce5247d874ffc60a762aa7abbb82c1695bbb171609c1b8861e279eb73.
//
// Solidity: event StakeDelegated(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares)
func (_Staking *StakingFilterer) FilterStakeDelegated(opts *bind.FilterOpts, indexer []common.Address, delegator []common.Address) (*StakingStakeDelegatedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeDelegated", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeDelegatedIterator{contract: _Staking.contract, event: "StakeDelegated", logs: logs, sub: sub}, nil
}

// WatchStakeDelegated is a free log subscription operation binding the contract event 0xcd0366dce5247d874ffc60a762aa7abbb82c1695bbb171609c1b8861e279eb73.
//
// Solidity: event StakeDelegated(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares)
func (_Staking *StakingFilterer) WatchStakeDelegated(opts *bind.WatchOpts, sink chan<- *StakingStakeDelegated, indexer []common.Address, delegator []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeDelegated", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeDelegated)
				if err := _Staking.contract.UnpackLog(event, "StakeDelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeDelegated is a log parse operation binding the contract event 0xcd0366dce5247d874ffc60a762aa7abbb82c1695bbb171609c1b8861e279eb73.
//
// Solidity: event StakeDelegated(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares)
func (_Staking *StakingFilterer) ParseStakeDelegated(log types.Log) (*StakingStakeDelegated, error) {
	event := new(StakingStakeDelegated)
	if err := _Staking.contract.UnpackLog(event, "StakeDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeDelegatedLockedIterator is returned from FilterStakeDelegatedLocked and is used to iterate over the raw logs and unpacked data for StakeDelegatedLocked events raised by the Staking contract.
type StakingStakeDelegatedLockedIterator struct {
	Event *StakingStakeDelegatedLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakeDelegatedLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeDelegatedLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStakeDelegatedLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakeDelegatedLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeDelegatedLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeDelegatedLocked represents a StakeDelegatedLocked event raised by the Staking contract.
type StakingStakeDelegatedLocked struct {
	Indexer   common.Address
	Delegator common.Address
	Tokens    *big.Int
	Shares    *big.Int
	Until     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeDelegatedLocked is a free log retrieval operation binding the contract event 0x0430183f84d9c4502386d499da806543dee1d9de83c08b01e39a6d2116c43b25.
//
// Solidity: event StakeDelegatedLocked(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares, uint256 until)
func (_Staking *StakingFilterer) FilterStakeDelegatedLocked(opts *bind.FilterOpts, indexer []common.Address, delegator []common.Address) (*StakingStakeDelegatedLockedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeDelegatedLocked", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeDelegatedLockedIterator{contract: _Staking.contract, event: "StakeDelegatedLocked", logs: logs, sub: sub}, nil
}

// WatchStakeDelegatedLocked is a free log subscription operation binding the contract event 0x0430183f84d9c4502386d499da806543dee1d9de83c08b01e39a6d2116c43b25.
//
// Solidity: event StakeDelegatedLocked(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares, uint256 until)
func (_Staking *StakingFilterer) WatchStakeDelegatedLocked(opts *bind.WatchOpts, sink chan<- *StakingStakeDelegatedLocked, indexer []common.Address, delegator []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeDelegatedLocked", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeDelegatedLocked)
				if err := _Staking.contract.UnpackLog(event, "StakeDelegatedLocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeDelegatedLocked is a log parse operation binding the contract event 0x0430183f84d9c4502386d499da806543dee1d9de83c08b01e39a6d2116c43b25.
//
// Solidity: event StakeDelegatedLocked(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares, uint256 until)
func (_Staking *StakingFilterer) ParseStakeDelegatedLocked(log types.Log) (*StakingStakeDelegatedLocked, error) {
	event := new(StakingStakeDelegatedLocked)
	if err := _Staking.contract.UnpackLog(event, "StakeDelegatedLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeDelegatedWithdrawnIterator is returned from FilterStakeDelegatedWithdrawn and is used to iterate over the raw logs and unpacked data for StakeDelegatedWithdrawn events raised by the Staking contract.
type StakingStakeDelegatedWithdrawnIterator struct {
	Event *StakingStakeDelegatedWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakeDelegatedWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeDelegatedWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStakeDelegatedWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakeDelegatedWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeDelegatedWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeDelegatedWithdrawn represents a StakeDelegatedWithdrawn event raised by the Staking contract.
type StakingStakeDelegatedWithdrawn struct {
	Indexer   common.Address
	Delegator common.Address
	Tokens    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeDelegatedWithdrawn is a free log retrieval operation binding the contract event 0x1b2e7737e043c5cf1b587ceb4daeb7ae00148b9bda8f79f1093eead08f141952.
//
// Solidity: event StakeDelegatedWithdrawn(address indexed indexer, address indexed delegator, uint256 tokens)
func (_Staking *StakingFilterer) FilterStakeDelegatedWithdrawn(opts *bind.FilterOpts, indexer []common.Address, delegator []common.Address) (*StakingStakeDelegatedWithdrawnIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeDelegatedWithdrawn", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeDelegatedWithdrawnIterator{contract: _Staking.contract, event: "StakeDelegatedWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeDelegatedWithdrawn is a free log subscription operation binding the contract event 0x1b2e7737e043c5cf1b587ceb4daeb7ae00148b9bda8f79f1093eead08f141952.
//
// Solidity: event StakeDelegatedWithdrawn(address indexed indexer, address indexed delegator, uint256 tokens)
func (_Staking *StakingFilterer) WatchStakeDelegatedWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingStakeDelegatedWithdrawn, indexer []common.Address, delegator []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeDelegatedWithdrawn", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeDelegatedWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "StakeDelegatedWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeDelegatedWithdrawn is a log parse operation binding the contract event 0x1b2e7737e043c5cf1b587ceb4daeb7ae00148b9bda8f79f1093eead08f141952.
//
// Solidity: event StakeDelegatedWithdrawn(address indexed indexer, address indexed delegator, uint256 tokens)
func (_Staking *StakingFilterer) ParseStakeDelegatedWithdrawn(log types.Log) (*StakingStakeDelegatedWithdrawn, error) {
	event := new(StakingStakeDelegatedWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "StakeDelegatedWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the Staking contract.
type StakingStakeDepositedIterator struct {
	Event *StakingStakeDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStakeDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeDeposited represents a StakeDeposited event raised by the Staking contract.
type StakingStakeDeposited struct {
	Indexer common.Address
	Tokens  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed indexer, uint256 tokens)
func (_Staking *StakingFilterer) FilterStakeDeposited(opts *bind.FilterOpts, indexer []common.Address) (*StakingStakeDepositedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeDeposited", indexerRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeDepositedIterator{contract: _Staking.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed indexer, uint256 tokens)
func (_Staking *StakingFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *StakingStakeDeposited, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeDeposited", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeDeposited)
				if err := _Staking.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeDeposited is a log parse operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed indexer, uint256 tokens)
func (_Staking *StakingFilterer) ParseStakeDeposited(log types.Log) (*StakingStakeDeposited, error) {
	event := new(StakingStakeDeposited)
	if err := _Staking.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the Staking contract.
type StakingStakeLockedIterator struct {
	Event *StakingStakeLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStakeLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeLocked represents a StakeLocked event raised by the Staking contract.
type StakingStakeLocked struct {
	Indexer common.Address
	Tokens  *big.Int
	Until   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed indexer, uint256 tokens, uint256 until)
func (_Staking *StakingFilterer) FilterStakeLocked(opts *bind.FilterOpts, indexer []common.Address) (*StakingStakeLockedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeLocked", indexerRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeLockedIterator{contract: _Staking.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed indexer, uint256 tokens, uint256 until)
func (_Staking *StakingFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *StakingStakeLocked, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeLocked", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeLocked)
				if err := _Staking.contract.UnpackLog(event, "StakeLocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed indexer, uint256 tokens, uint256 until)
func (_Staking *StakingFilterer) ParseStakeLocked(log types.Log) (*StakingStakeLocked, error) {
	event := new(StakingStakeLocked)
	if err := _Staking.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeSlashedIterator is returned from FilterStakeSlashed and is used to iterate over the raw logs and unpacked data for StakeSlashed events raised by the Staking contract.
type StakingStakeSlashedIterator struct {
	Event *StakingStakeSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStakeSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeSlashed represents a StakeSlashed event raised by the Staking contract.
type StakingStakeSlashed struct {
	Indexer     common.Address
	Tokens      *big.Int
	Reward      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeSlashed is a free log retrieval operation binding the contract event 0xf2717be2f27d9d2d7d265e42dc556e40d2d9aeaba02f49c5286030f30c0571f3.
//
// Solidity: event StakeSlashed(address indexed indexer, uint256 tokens, uint256 reward, address beneficiary)
func (_Staking *StakingFilterer) FilterStakeSlashed(opts *bind.FilterOpts, indexer []common.Address) (*StakingStakeSlashedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeSlashed", indexerRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeSlashedIterator{contract: _Staking.contract, event: "StakeSlashed", logs: logs, sub: sub}, nil
}

// WatchStakeSlashed is a free log subscription operation binding the contract event 0xf2717be2f27d9d2d7d265e42dc556e40d2d9aeaba02f49c5286030f30c0571f3.
//
// Solidity: event StakeSlashed(address indexed indexer, uint256 tokens, uint256 reward, address beneficiary)
func (_Staking *StakingFilterer) WatchStakeSlashed(opts *bind.WatchOpts, sink chan<- *StakingStakeSlashed, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeSlashed", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeSlashed)
				if err := _Staking.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeSlashed is a log parse operation binding the contract event 0xf2717be2f27d9d2d7d265e42dc556e40d2d9aeaba02f49c5286030f30c0571f3.
//
// Solidity: event StakeSlashed(address indexed indexer, uint256 tokens, uint256 reward, address beneficiary)
func (_Staking *StakingFilterer) ParseStakeSlashed(log types.Log) (*StakingStakeSlashed, error) {
	event := new(StakingStakeSlashed)
	if err := _Staking.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Staking contract.
type StakingStakeWithdrawnIterator struct {
	Event *StakingStakeWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStakeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeWithdrawn represents a StakeWithdrawn event raised by the Staking contract.
type StakingStakeWithdrawn struct {
	Indexer common.Address
	Tokens  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed indexer, uint256 tokens)
func (_Staking *StakingFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, indexer []common.Address) (*StakingStakeWithdrawnIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeWithdrawn", indexerRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeWithdrawnIterator{contract: _Staking.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed indexer, uint256 tokens)
func (_Staking *StakingFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingStakeWithdrawn, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeWithdrawn", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed indexer, uint256 tokens)
func (_Staking *StakingFilterer) ParseStakeWithdrawn(log types.Log) (*StakingStakeWithdrawn, error) {
	event := new(StakingStakeWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
