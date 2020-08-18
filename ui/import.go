package ui

import (
	"EthSea/myapp"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)


func GetImportLayout() fyne.CanvasObject{
	viewSize := fyne.NewSize(Width,Height)

	inputTip := widget.NewLabel("Input your private key!")

	inputEdit := widget.NewEntry()
	inputEdit.SetPlaceHolder("your private key")


	importBt := widget.NewButton("Import",comfirImportBtClick)

	importBt.Resize(importBt.MinSize().Add(fyne.NewSize(200,0)))

	inputTip.Resize(inputTip.MinSize())

	inputEdit.Resize(inputEdit.MinSize().Add(fyne.NewSize(200,100)))

	SetWidgetHCenter(importBt,viewSize)

	SetWidgetHCenter(inputTip,viewSize)
	SetWidgetHCenter(inputEdit,viewSize)

	SetWidgetY(inputEdit,400)

	SetWidgetY(inputTip,300)

	SetWidgetY(importBt,600)




	lay := fyne.NewContainerWithLayout (&AbLayout{Width,Height}, inputTip,inputEdit,importBt)

	return lay


}
func comfirImportBtClick(){



	fmt.Println("confirm import")

	myapp.WindowInstall.SetContent(GetWalletLayout())


}