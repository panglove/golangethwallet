package ui

import (
	"EthSea/myapp"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)


func GetWalletLayout() fyne.CanvasObject{

	viewSize := fyne.NewSize(Width,Height)

	welcomeLabel := widget.NewLabel("Welcome to Eth Sea Wallet!")


	currLabel := widget.NewLabel("Current Wallet")

	addressSelect :=widget.NewSelect([]string{"0xbd4b5fa5c36ff03c78b6a48d849be49f91ac3137","0xbe4bcd19678021c43492cd0c246836b6ad6fa2f6"},AddressSelectChange)



//	addressSelect.PlaceHolder ="Select One Wallet"

	addressSelect.SetSelected("ahdashdasahdashdasiodhpvzhokpvhzoiodhpvzhokpvhzo")

	addressSelect.Resize(addressSelect.MinSize().Add(fyne.NewSize(50,0)))

	currLabel.Resize(currLabel.MinSize())

	transferBt := widget.NewButton("Go To Transfer",goToTransferBtClick)


	transferBt.Resize(transferBt.MinSize().Add(fyne.NewSize(322,0)))
	welcomeLabel.Resize(welcomeLabel.MinSize())
	SetWidgetHCenter(addressSelect,viewSize)
	SetWidgetHCenter(transferBt,viewSize)
	SetWidgetHCenter(welcomeLabel,viewSize)
	SetWidgetPosition(currLabel,13,70)

	SetWidgetY(addressSelect,100)


	SetWidgetY(transferBt,150)

	SetWidgetY(welcomeLabel,10)


	lay := fyne.NewContainerWithLayout (&AbLayout{Width,Height}, welcomeLabel,transferBt,addressSelect,currLabel)


	return lay


}
func AddressSelectChange(address string){


}
func goToTransferBtClick(){

	dialog.ShowCustom("Transfer","Close",GetTransferDialogLayout(),myapp.WindowInstall)



}
