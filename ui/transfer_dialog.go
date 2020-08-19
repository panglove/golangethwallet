package ui

import (
	"YourMoney/config"
	"YourMoney/util/wallet"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"strconv"
)

const (
	DialogWidth  = 300
	DialogHeight = 300
)

var addressEdit, countEdit *widget.Entry

var currSelectWallet *config.Wallet

func SetSelectWallet(wallet2 *config.Wallet) {

	currSelectWallet = wallet2

}

func GetTransferDialogLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(DialogWidth, DialogHeight)

	inputTip := widget.NewLabel("Input Address And Transfer Count")

	addressEdit = widget.NewMultiLineEntry()
	addressEdit.SetPlaceHolder("Receipt address")
	addressEdit.Wrapping = fyne.TextWrapBreak
	addressEdit.Resize(fyne.NewSize(340, 60))
	SetWidgetHCenter(addressEdit, viewSize)

	SetWidgetY(addressEdit, 45)
	countEdit = widget.NewEntry()
	countEdit.SetPlaceHolder("Transfer Count")
	countEdit.Resize(countEdit.MinSize().Add(fyne.NewSize(207, 0)))
	SetWidgetHCenter(countEdit, viewSize)

	SetWidgetY(countEdit, 120)

	transferBt := widget.NewButton("Transfer", confirmTransferBtClick)

	transferBt.Resize(transferBt.MinSize().Add(fyne.NewSize(260, 0)))

	inputTip.Resize(inputTip.MinSize())

	SetWidgetHCenter(transferBt, viewSize)

	SetWidgetHCenter(inputTip, viewSize)

	SetWidgetY(inputTip, 10)

	SetWidgetY(transferBt, 180)

	lay := fyne.NewContainerWithLayout(&AbLayout{DialogWidth, DialogHeight}, inputTip, addressEdit, countEdit, transferBt)

	return lay

}
func confirmTransferBtClick() {

	address := addressEdit.Text

	if len(address) < 42 {

		Alert("Address is at least 42 digits")
		return
	}

	count, err := strconv.ParseFloat(countEdit.Text, 64)

	if err != nil || count <= 0 {
		Alert("Transfer quantity must be bigger than 0")
		return
	}

	hash, isOk := wallet.Tranfer(currSelectWallet.PrivateKey, address, count)

	if isOk && len(hash) > 0 {

		Alert("Transaction has be sent!\n Tx hash:" + hash[:30]+"\n"+hash[31:])
		return
	} else {

		Alert("Send Transaction Error")
		return
	}

}
