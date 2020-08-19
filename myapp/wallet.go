package myapp

import (
	"YourMoney/config"
	"strings"
)

func GetWalletListString() []string {

	var walletStringList []string

	for _, wallet := range AppSetting.WalletList {

		walletStringList = append(walletStringList, wallet.Name+":"+wallet.Address)

	}
	return walletStringList
}

func IsAddressExists(address string) bool {

	for _, wallet := range AppSetting.WalletList {

		if strings.ToLower(address) == strings.ToLower(wallet.Address) {

			return true
		}

	}
	return false
}

func GetWalletByAddress(address string) *config.Wallet {
	for _, wallet := range AppSetting.WalletList {

		if strings.ToLower(address) == strings.ToLower(wallet.Address) {

			return &wallet
		}

	}
	return nil
}
func GetWalletByWalletString(wallet string) *config.Wallet {

	walletInfo := strings.Split(wallet, ":")

	return GetWalletByAddress(walletInfo[1])

}

func ImportWallet(name string, pri string, pub string) (bool, string) {

	if IsAddressExists(pub) {
		return false, "The address already exists!"
	}

	AppSetting.WalletList = append(AppSetting.WalletList, config.Wallet{
		name, pri, pub,
	})

	isok := WriteSetting()

	if !isok {
		return false, "System error"
	}
	return true, ""
}
