package ui

import (
	"YourMoney/myapp"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

var backBt *widget.Button

var isHide bool = true

func GetChooseLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(Width,Height)

	tip1 := widget.NewLabel("Use PrivateKey To Import Wallet!")

	tip2 := widget.NewLabel("Create A New Wallet!")

	importBt := widget.NewButton("From PrivateKey", importBtClick)

	createBt := widget.NewButton("Create Wallet", createBtClick)
	createBt.Resize(fyne.NewSize(400, createBt.MinSize().Height))
	SetWidgetHCenter(createBt,viewSize)
	SetWidgetY(createBt,400)


	tip0 := widget.NewLabel("Use Keystore File To Import Wallet!")
	tip0.Resize(tip0.MinSize())
	SetWidgetHCenter(tip0,viewSize)
	SetWidgetY(tip0,250)
	fromFileBt := widget.NewButton("From Keystore", fromKeyStoreBtClick)
	fromFileBt.Resize(fyne.NewSize(400, fromFileBt.MinSize().Height))
	SetWidgetHCenter(fromFileBt,viewSize)

	SetWidgetY(fromFileBt,300)



	backBt = widget.NewButton("Back", backBtClick)
	backBt.Resize(fyne.NewSize(400, backBt.MinSize().Height))
	SetWidgetHCenter(backBt,viewSize)
	SetWidgetY(backBt,450)



	if isHide {
		backBt.Hide()
	}




	tip1.Resize(tip1.MinSize())

	tip2.Resize(tip2.MinSize())


	importBt.Resize(fyne.NewSize(400, importBt.MinSize().Height))

	SetWidgetHCenter(importBt,viewSize)

	SetWidgetY(tip1,150)
	SetWidgetY(tip2,350)
	SetWidgetHCenter(tip1,viewSize)

	SetWidgetHCenter(tip2,viewSize)

	SetWidgetY(importBt,200)

	lay := fyne.NewContainerWithLayout(&AbLayout{Width, Height}, importBt, createBt,fromFileBt, backBt,tip1,tip2,tip0)

	return lay

}
func backBtClick(){

	myapp.WindowInstall.SetContent(GetWalletLayout())


}
func fromKeyStoreBtClick() {
	myapp.WindowInstall.SetContent(GetImportKeyStoreLayout())

}
func importBtClick() {
	myapp.WindowInstall.SetContent(GetImportLayout())

}

func createBtClick() {
	myapp.WindowInstall.SetContent(GetCreateLayout())

}
