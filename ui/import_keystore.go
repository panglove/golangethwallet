package ui

import (
	"YourMoney/config"
	"YourMoney/myapp"
	"YourMoney/util/wallet"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

var keyStorePassEdit, walletNameEdit *widget.Entry
var fileChooseBt *widget.Button
var currKeyStoreStr string

func GetImportKeyStoreLayout() fyne.CanvasObject {
	currKeyStoreStr = ""
	viewSize := fyne.NewSize(Width, Height)

	inputTip := widget.NewLabel("Choose Your KeyStore File!")

	keyStorePassEdit = widget.NewEntry()
	keyStorePassEdit.SetPlaceHolder("Input Keystore Pass")

	walletNameEdit = widget.NewEntry()
	walletNameEdit.SetPlaceHolder("Input Wallet Name")

	fileChooseBt = widget.NewButton("Choose Keystore File", fileChooseBtClick)

	fileChooseBt.Resize(fileChooseBt.MinSize().Add(fyne.NewSize(200, 0)))

	SetWidgetHCenter(fileChooseBt, viewSize)
	SetWidgetY(fileChooseBt, 250)

	createBt := widget.NewButton("Import", comfirmImportKeystoreBtClick)

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

	keyStorePassEdit.Resize(keyStorePassEdit.MinSize().Add(fyne.NewSize(200, 0)))

	SetWidgetHCenter(inputTip, viewSize)
	SetWidgetHCenter(keyStorePassEdit, viewSize)

	SetWidgetY(keyStorePassEdit, 400)

	walletNameEdit.Resize(keyStorePassEdit.MinSize().Add(fyne.NewSize(200, 0)))

	SetWidgetHCenter(walletNameEdit, viewSize)

	SetWidgetY(walletNameEdit, 460)

	SetWidgetY(inputTip, 150)

	lay := fyne.NewContainerWithLayout(&AbLayout{Width, Height}, inputTip, keyStorePassEdit, backToBt, createBt, fileChooseBt, walletNameEdit)

	return lay

}
func fileChooseBtClick() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader==nil{
			return
		}

		fileChooseBt.SetText(reader.Name())

		fileBytes := make([]byte, 1024*1000)

		n, err := reader.Read(fileBytes)
		if err != nil {
			Alert("keystore file error")
			return
		}

		fileBytes = fileBytes[:n]

		currKeyStoreStr = string(fileBytes)

	}, myapp.WindowInstall)

}
func comfirmImportKeystoreBtClick() {

	nameStr := walletNameEdit.Text

	if len(nameStr) <= 0 {

		Alert("Wallet name is at least one digits")
		return
	}

	passStr := keyStorePassEdit.Text

	if len(passStr) <= 0 {

		Alert("Keystore password is at least one digits")
		return
	}

	if len(currKeyStoreStr) > 0 {

		privateKey, isOk := wallet.KeyStoreToPrivateKey(currKeyStoreStr, passStr)

		if !isOk {
			Alert("Keystore file or password error")
			return
		}
		ImportPrivateKey(nameStr, privateKey)

	} else {

		Alert("Please choose your keystore file!")

		return
	}

}

func ImportPrivateKey(nameEdit string, priEdit string) {

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
		nameEdit, priEdit, publicHex,
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
