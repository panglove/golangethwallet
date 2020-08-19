package ui

import (
	"YourMoney/config"
	"YourMoney/myapp"
	"YourMoney/util/wallet"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

var (
	importNameEdit *widget.Entry

	importPrivateEdit *widget.Entry
)

func GetImportLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(Width, Height)

	inputTip := widget.NewLabel("Input your wallet name!")
	inputTip.Resize(inputTip.MinSize())
	SetWidgetHCenter(inputTip, viewSize)
	SetWidgetY(inputTip, 50)

	inputTip2 := widget.NewLabel("Input your private key!")
	inputTip2.Resize(inputTip2.MinSize())
	SetWidgetHCenter(inputTip2, viewSize)
	SetWidgetY(inputTip2, 250)
	importNameEdit = widget.NewEntry()
	importNameEdit.SetPlaceHolder("your wallet name")
	importNameEdit.Resize(importNameEdit.MinSize().Add(fyne.NewSize(200, 0)))
	SetWidgetHCenter(importNameEdit, viewSize)
	SetWidgetY(importNameEdit, 100)

	importPrivateEdit = widget.NewMultiLineEntry()
	importPrivateEdit.SetPlaceHolder("your private key")
	importPrivateEdit.Resize(importPrivateEdit.MinSize().Add(fyne.NewSize(200, 50)))
	importPrivateEdit.Wrapping = fyne.TextWrapBreak
	SetWidgetHCenter(importPrivateEdit, viewSize)
	SetWidgetY(importPrivateEdit, 300)

	importBt := widget.NewButton("Import", comfirImportBtClick)

	importBt.Resize(importBt.MinSize().Add(fyne.NewSize(200, 0)))

	SetWidgetHCenter(importBt, viewSize)

	SetWidgetY(importBt, 500)

	backToBt := widget.NewButton("Back", func() {
		myapp.WindowInstall.SetContent(GetChooseLayout())
	})

	backToBt.Resize(backToBt.MinSize().Add(fyne.NewSize(217, 0)))

	SetWidgetHCenter(backToBt, viewSize)
	SetWidgetY(backToBt, 600)

	lay := fyne.NewContainerWithLayout(&AbLayout{Width, Height}, inputTip, inputTip2, importNameEdit, importPrivateEdit,backToBt, importBt)

	return lay

}
func comfirImportBtClick() {

	nameEdit := importNameEdit.Text

	if len(nameEdit) <= 0 {

		Alert("Wallet name is at least one digits")
		return
	}

	priEdit := importPrivateEdit.Text

	if len(priEdit) < 10 {

		Alert("PrivateKey is at least ten digits")
		return
	}

	publicHex, isPri := wallet.PrivateKeyToPublicHex(priEdit)

	if !isPri {

		Alert("PrivateKey is not incorrect")
		return
	}
	fmt.Println("import wallet address :", publicHex)

	if myapp.IsAddressExists(publicHex) {
		Alert("The address already exists!")
		return
	}

	myapp.AppSetting.WalletList = append(myapp.AppSetting.WalletList, config.Wallet{
		nameEdit, priEdit,publicHex,
	})
	isok := myapp.WriteSetting()

	if !isok {

		Alert("System error")
		return
	}
	Comfirm(publicHex+"\n Import Success!", func() {
		myapp.WindowInstall.SetContent(GetWalletLayout())
	})

}
