package eth

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stakemachine/graph-indexer-cli/eth/abi/staking"
)

// const zeroPoi = "0x0000000000000000000000000000000000000000000000000000000000000000"

func (es *Service) CloseAllocation(privateKeyHex, address, contract, alloID, poi string) (*types.Transaction, error) {
	allocationID := common.HexToAddress(alloID)
	contractAddress := common.HexToAddress(contract)
	addressFrom := common.HexToAddress(address)
	if !common.IsHexAddress(alloID) {
		return nil, errors.New("allocation ID is invalid")
	}

	if !strings.HasPrefix(poi, "0x") {
		return nil, errors.New("poi should begin with 0x prefix")
	}

	poiBytes32 := common.Hex2BytesFixed(poi, 32)

	var poi32 [32]byte
	copy(poi32[:], poiBytes32[:32])

	instance, err := staking.NewStaking(contractAddress, es.Client)
	if err != nil {
		return nil, errors.New("init allocation ID: " + allocationID.String() + err.Error())
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	nonce, err := es.Client.PendingNonceAt(context.Background(), addressFrom)
	if err != nil {
		return nil, err
	}

	gasPrice, err := es.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := es.Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	gasTipCap, err := es.Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	transactOpts.From = addressFrom
	transactOpts.NoSend = true
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasFeeCap = gasPrice
	transactOpts.GasTipCap = gasTipCap
	transactOpts.GasLimit = 300000

	tx, err := instance.CloseAllocation(transactOpts, allocationID, poi32)
	if err != nil {
		return nil, errors.New("instance allocation ID: " + allocationID.String() + " " + err.Error())
	}
	return tx, nil
}

func (es *Service) GetAllocation(contractAddress, alloID string) (staking.IStakingDataAllocation, error) {
	allocationID := common.HexToAddress(alloID)
	instance, err := staking.NewStaking(common.HexToAddress(contractAddress), es.Client)
	if err != nil {
		return staking.IStakingDataAllocation{}, errors.New("allocation ID: " + allocationID.String() + err.Error())
	}
	allocationData, err := instance.GetAllocation(&bind.CallOpts{}, allocationID)
	if err != nil {
		return staking.IStakingDataAllocation{}, errors.New("allocation ID: " + allocationID.String() + err.Error())
	}

	return allocationData, nil
}

func (es *Service) GetAllocationState(contractAddress, alloID string) (uint8, error) {
	allocationID := common.HexToAddress(alloID)
	instance, err := staking.NewStaking(common.HexToAddress(contractAddress), es.Client)
	if err != nil {
		return 0, errors.New("allocation ID: " + allocationID.String() + err.Error())
	}
	allocationState, err := instance.GetAllocationState(&bind.CallOpts{}, allocationID)
	if err != nil {
		return 0, errors.New("allocation ID: " + allocationID.String() + err.Error())
	}

	return allocationState, nil
}
