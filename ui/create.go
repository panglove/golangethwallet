package ui

import (
	"YourMoney/myapp"
	"YourMoney/util/wallet"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

var nameEdit *widget.Entry

func GetCreateLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(Width, Height)

	inputTip := widget.NewLabel("Input Wallet Name!")

	nameEdit = widget.NewEntry()
	nameEdit.SetPlaceHolder("Your Wallet Name")

	createBt := widget.NewButton("Create", comfirmCreateBtClick)

	createBt.Resize(createBt.MinSize().Add(fyne.NewSize(200, 0)))

	SetWidgetHCenter(createBt, viewSize)
	SetWidgetY(createBt, 600)

	backToBt := widget.NewButton("Back", func() {

		myapp.WindowInstall.SetContent(GetChooseLayout())
	})

	backToBt.Resize(backToBt.MinSize().Add(fyne.NewSize(217, 0)))

	SetWidgetHCenter(backToBt, viewSize)
	SetWidgetY(backToBt, 700)

	inputTip.Resize(inputTip.MinSize())

	nameEdit.Resize(nameEdit.MinSize().Add(fyne.NewSize(200, 0)))

	SetWidgetHCenter(inputTip, viewSize)
	SetWidgetHCenter(nameEdit, viewSize)

	SetWidgetY(nameEdit, 400)

	SetWidgetY(inputTip, 300)

	lay := fyne.NewContainerWithLayout(&AbLayout{Width, Height}, inputTip, nameEdit, backToBt, createBt)

	return lay

}
func comfirmCreateBtClick() {

	nameStr := nameEdit.Text

	if len(nameStr) <= 0 {

		Alert("Wallet name is at least one digits")
		return
	}
	priKey, pubKey := wallet.CreateWallet()

	if len(priKey) <= 0 {

		Alert("System Error")
		return
	}

	isok, err := myapp.ImportWallet(nameStr, priKey, pubKey)

	if !isok {
		Alert(err)
		return
	}
	Comfirm(pubKey+"\n Create Success!", func() {
		myapp.WindowInstall.SetContent(GetWalletLayout())
	})

}
