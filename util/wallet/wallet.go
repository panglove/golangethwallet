package wallet

import (
	"YourMoney/util/mathutil"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var CurrEthClient *ethclient.Client

func GetBalance(address string) float64 {

	b, err := CurrEthClient.BalanceAt(context.Background(), common.HexToAddress(address), nil)

	if err != nil {
		return 0
	}
	return mathutil.WeiToFloat(b)

}

func Init(rpUrl string) bool {
	var err error

	CurrEthClient, err = ethclient.Dial(rpUrl)

	if err == nil {
		return true
	}
	fmt.Println("rpc error ", err)
	return false
}
func PrivateKeyToPublicHex(privateKey string) (string, bool) {

	priviteKeyOb, err := crypto.HexToECDSA(privateKey)

	if err != nil {
		fmt.Println("chanid err", err)
		return "", false
	}

	publicKey := priviteKeyOb.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", false
	}
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return publicAddress.String(), true

}

func Tranfer(privateKey string, toAddress string, count float64) (string, bool) {

	priviteKey, err := crypto.HexToECDSA(privateKey)

	publicAddress, ok := PrivateKeyToPublicHex(privateKey)

	if !ok {
		return "", false
	}

	if err != nil {
		fmt.Println("chanid err", err)
		return "", false
	}

	nonce, err := CurrEthClient.PendingNonceAt(context.Background(), common.HexToAddress(publicAddress))

	if err != nil {
		fmt.Println("err", err)
		return "", false
	}

	weiMount := mathutil.FloatToWei(count)

	gasLimit := uint64(8000000)

	gasPrice, err2 := CurrEthClient.SuggestGasPrice(context.Background())

	if err2 != nil {
		fmt.Println("err", err2)
		return "", false
	}

	data := []byte("0x")

	newTrans := types.NewTransaction(nonce, common.HexToAddress(toAddress), weiMount, gasLimit, gasPrice, data)

	chainID, err := CurrEthClient.NetworkID(context.Background())
	if err != nil {
		fmt.Println("chanid err", err)
		return "", false
	}

	signer := types.NewEIP155Signer(chainID)

	signTx, err := types.SignTx(newTrans, signer, priviteKey)

	if err != nil {
		fmt.Println("sign err", err)
		return "", false
	}

	errTrans := CurrEthClient.SendTransaction(context.Background(), signTx)

	if errTrans != nil {
		fmt.Println("err", errTrans)

		return "", false
	}
	return signTx.Hash().String(), true
}

func CreateWallet() (pri string, pub string) {

	prig, err := crypto.GenerateKey()

	if err != nil {
		return "", ""
	}

	pri = common.Bytes2Hex(crypto.FromECDSA(prig))

	pub, _ = PrivateKeyToPublicHex(pri)

	return pri, pub

}
