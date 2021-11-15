package eth

import "github.com/ethereum/go-ethereum/ethclient"

// EthService ethereum service
type Service struct {
	// ks     *keystore.KeyStore
	Client *ethclient.Client
}
