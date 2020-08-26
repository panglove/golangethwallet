package myapp

import (
	"YourMoney/config"
	"YourMoney/util/mathutil"
	"YourMoney/util/wallet"
	"strings"
)

func GetWalletListString() []string {

	var walletIStringList []string

	for _, walletI := range AppSetting.WalletList {

		walletIStringList = append(walletIStringList, walletI.Name+":"+walletI.Address+"("+mathutil.FloatToString(wallet.GetBalance(walletI.Address))+")")

	}
	return walletIStringList
}

func IsAddressExists(address string) bool {

	for _, walletI := range AppSetting.WalletList {

		if strings.ToLower(address) == strings.ToLower(walletI.Address) {

			return true
		}

	}
	return false
}

func GetWalletByAddress(address string) *config.Wallet {
	for _, walletI := range AppSetting.WalletList {

		if strings.ToLower(address) == strings.ToLower(walletI.Address) {

			return &walletI
		}

	}
	return nil
}
func GetWalletByWalletString(walletI string) *config.Wallet {

	walletIInfo := strings.Split(walletI, ":")

	conIndex := strings.Index(walletIInfo[1], "(")

	walletIInfoStr := walletIInfo[1][:conIndex]

	return GetWalletByAddress(walletIInfoStr)

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
func ExportWalletPriviteKey(address string) string {

	eWallet := GetWalletByAddress(address)

	if eWallet != nil {

		return eWallet.PrivateKey

	} else
	{
		return ""
	}
}
