package eth

import (
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func GetWalletAccount(mnemonic, path string) (accounts.Account, string, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return accounts.Account{}, "", err
	}

	accountPath := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(accountPath, true)
	if err != nil {
		return accounts.Account{}, "", err
	}

	privateKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return accounts.Account{}, "", err
	}
	return account, privateKeyHex, nil
}
