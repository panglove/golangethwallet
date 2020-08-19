package ui

import (
	"YourMoney/config"
	"YourMoney/myapp"
	"YourMoney/util/storage"
	"YourMoney/util/wallet"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

const (
	Width  = 480
	Height = 800
)

var passEdit *widget.Entry

var addRpcDialog dialog.Dialog

var rpcSelect *widget.Select

func GetIndexLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(Width, Height)

	welcomeLabel := widget.NewLabel("Welcome to Eth Sea Wallet!")

	versionLabel := widget.NewLabel("Make by " + config.AppAuthor + " V " + config.AppVersion)

	passEdit = widget.NewPasswordEntry()

	passEdit.SetPlaceHolder("please input your password")

	passEdit.Resize(fyne.NewSize(450, passEdit.MinSize().Height))

	startBt := widget.NewButton("Get Start", startBtClick)

	startBt.Resize(startBt.MinSize().Add(fyne.NewSize(200, 0)))
	SetWidgetHCenter(startBt, viewSize)
	SetWidgetY(startBt, 500)

	welcomeLabel.Resize(welcomeLabel.MinSize())
	versionLabel.Resize(versionLabel.MinSize())

	SetWidgetHCenter(welcomeLabel, viewSize)

	SetWidgetHCenter(versionLabel, viewSize)

	SetWidgetHCenter(passEdit, viewSize)

	SetWidgetY(passEdit, 400)

	SetWidgetY(welcomeLabel, 60)

	SetWidgetY(versionLabel, 600)

	//RPC List

	currLabel := widget.NewLabel("Current Rpc Url:")

	currLabel.Resize(currLabel.MinSize())

	SetWidgetPosition(currLabel, 13, 250)

	rpcSelect = widget.NewSelect(myapp.AppSetting.RpcList, rpcSelectChange)

	//	rpcSelect.PlaceHolder ="Select One Wallet"
	if len(myapp.AppSetting.RpcList) > 0 {
		rpcSelect.SetSelected(myapp.AppSetting.RpcUrl)
	}

	rpcSelect.Resize((fyne.NewSize(400, rpcSelect.MinSize().Height)))

	SetWidgetHCenter(rpcSelect, viewSize)
	SetWidgetPosition(rpcSelect, 16, 280)
	addBt := widget.NewButton("Add", addBtClick)
	addBt.Resize(addBt.MinSize().Add(fyne.NewSize(0, 0)))
	SetWidgetHCenter(addBt, viewSize)
	SetWidgetPosition(addBt, 425, 280)

	lay := fyne.NewContainerWithLayout(&AbLayout{Width, Height}, welcomeLabel, startBt, addBt, passEdit, versionLabel, rpcSelect, currLabel)

	return lay

}
func addBtClick() {
	if addRpcDialog == nil {
		addRpcDialog = dialog.NewCustom("Add Rpc", "Close", GetAddRpcDialogLayout(), myapp.WindowInstall)
	}
	addRpcDialog.Show()

}
func finishAddRpc() {
	myapp.ReadSetting()
	rpcSelect.Options = myapp.AppSetting.RpcList
	rpcSelect.Refresh()
	if len(myapp.AppSetting.RpcList) > 0 && len(rpcSelect.Selected) == 0 {
		rpcSelect.SetSelected(myapp.AppSetting.RpcList[0])
	}
	addRpcDialog.Hide()
}
func rpcSelectChange(s string) {
	myapp.AppSetting.RpcUrl = s
	myapp.WriteSetting()
}

func startBtClick() {

	password := passEdit.Text

	if len(password) < 6 {

		dialog.ShowInformation("Tips", "Password length is at least six digits", myapp.WindowInstall)

		return

	}

	settingStr := storage.GetItem(config.AppSaveFileName)

	if len(settingStr) <= 0 || len(myapp.AppSetting.PassWord) <= 0 {

		myapp.AppSetting.PassWord = "" + password

		isOk := myapp.WriteSetting()

		if !isOk {
			dialog.ShowInformation("Tips", "System error", myapp.WindowInstall)
			return
		}

	} else {

		isok := myapp.ReadSetting()

		if !isok {
			dialog.ShowInformation("Tips", "System error", myapp.WindowInstall)
			return
		}
		if len(myapp.AppSetting.PassWord) > 0 && myapp.AppSetting.PassWord != password {
			dialog.ShowInformation("Tips", "The password is incorrect", myapp.WindowInstall)
			return
		}

	}

	isConnect := wallet.Init(myapp.AppSetting.RpcUrl)

	if !isConnect {
		Comfirm("Rpc Network Error", func() {

			myapp.AppInstall.Quit()
		})

	}

	if len(myapp.AppSetting.WalletList) == 0 {

		myapp.WindowInstall.SetContent(GetChooseLayout())

	} else {

		myapp.WindowInstall.SetContent(GetWalletLayout())
	}

}
