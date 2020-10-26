package ui

import (
	"YourMoney/config"
	"YourMoney/myapp"
	"YourMoney/util/mathutil"
	"YourMoney/util/wallet"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"time"
)

var selectWallet *config.Wallet
var banlanceLabel *widget.Label

var addressSelect *widget.Select

func GetWalletLayout() fyne.CanvasObject {

	viewSize := fyne.NewSize(Width, Height)

	welcomeLabel := widget.NewLabel("Welcome to Your Money Wallet!")



	banlanceLabel = widget.NewLabel("Current Balance:0")

	banlanceLabel.Resize(banlanceLabel.MinSize())

	SetWidgetPosition(banlanceLabel, 13, 60)

	walletList := myapp.GetWalletListString()


	currLabel := widget.NewLabel("Current Wallet:")

	currLabel.Resize(currLabel.MinSize())

	SetWidgetPosition(currLabel, 13, 90)
	addressSelect = widget.NewSelect(walletList, AddressSelectChange)

	//	addressSelect.PlaceHolder ="Select One Wallet"

	addressSelect.SetSelected(walletList[0])

	addressSelect.Resize((fyne.NewSize(550, addressSelect.MinSize().Height)))

	SetWidgetHCenter(addressSelect, viewSize)

	SetWidgetY(addressSelect, 120)





	selectWallet = &myapp.AppSetting.WalletList[0]

	loadBalanceLabel()



	transferBt := widget.NewButton("Go To Transfer", goToTransferBtClick)

	transferBt.Resize((fyne.NewSize(550, transferBt.MinSize().Height)))

	SetWidgetHCenter(transferBt, viewSize)
	SetWidgetY(transferBt, 170)

	createAccountBt := widget.NewButton("Create A Wallet", goToChooseBtClick)
	createAccountBt.Resize((fyne.NewSize(550, createAccountBt.MinSize().Height)))
	SetWidgetHCenter(createAccountBt, viewSize)
	SetWidgetY(createAccountBt, 220)

	copyBt := widget.NewButton("Copy Wallet", copyAddressBtClick)
	copyBt.Resize((fyne.NewSize(550, copyBt.MinSize().Height)))
	SetWidgetHCenter(copyBt, viewSize)
	SetWidgetY(copyBt, 270)

	exportBt := widget.NewButton("Export Wallet", exportBtClick)
	exportBt.Resize((fyne.NewSize(550, exportBt.MinSize().Height)))
	SetWidgetHCenter(exportBt, viewSize)
	SetWidgetY(exportBt, 320)
	removeBt := widget.NewButton("Remove Wallet", removeBtClick)
	removeBt.Resize((fyne.NewSize(550, removeBt.MinSize().Height)))
	SetWidgetHCenter(removeBt, viewSize)
	SetWidgetY(removeBt, 370)

	welcomeLabel.Resize(welcomeLabel.MinSize())

	SetWidgetHCenter(welcomeLabel, viewSize)


	SetWidgetY(welcomeLabel, 10)

	lay := fyne.NewContainerWithLayout(&AbLayout{viewSize.Width, viewSize.Height}, welcomeLabel, banlanceLabel, exportBt,transferBt,removeBt, createAccountBt,copyBt, addressSelect, currLabel)

	IntervalLoadBalance()

	return lay

}
func exportBtClick(){

	if selectWallet!=nil {

		myapp.WindowInstall.Clipboard().SetContent(""+selectWallet.PrivateKey)
		Alert("PrivateKey Has Copyed")

	}
}
func removeBtClick(){

	if selectWallet!=nil {
		myapp.RemoveWallet(selectWallet.Address)
		myapp.WindowInstall.SetContent(GetWalletLayout())
		Alert("Remove Success")

	}
}
func IntervalLoadBalance() {
	go func() {
		for {
			loadBalanceLabel()

			addressSelect.Options = myapp.GetWalletListString()

			time.Sleep(5*time.Second)

		}
	}()
}

func loadBalanceLabel() {
	if selectWallet!=nil {
		banlanceLabel.SetText("Current Balance:" + mathutil.FloatToString(wallet.GetBalance(selectWallet.Address)))

	}

}

func AddressSelectChange(walletString string) {

	selectWallet = myapp.GetWalletByWalletString(walletString)

	if selectWallet!=nil {

		SetSelectWallet(selectWallet)

		loadBalanceLabel()
	}


}
func goToTransferBtClick() {

	dialog.ShowCustom("Transfer", "Close", GetTransferDialogLayout(), myapp.WindowInstall)

}
func goToChooseBtClick() {


	myapp.WindowInstall.SetContent(GetChooseLayout())

	backBt.Hidden = false
	isHide = false
}
func copyAddressBtClick(){

	myapp.WindowInstall.Clipboard().SetContent(""+selectWallet.Address)

	Alert("Copy Success")


}