package ui

import (
	"EthSea/myapp"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func GetChooseLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(Width,Height)

	tip1 := widget.NewLabel("Use PrivateKey To Import Wallet!")

	tip2 := widget.NewLabel("Create A New Wallet!")

	importBt := widget.NewButton("Import Wallet", importBtClick)

	createBt := widget.NewButton("Create Wallet", createBtClick)


	tip1.Resize(tip1.MinSize())

	tip2.Resize(tip2.MinSize())


	importBt.Resize(fyne.NewSize(400, importBt.MinSize().Height))

	createBt.Resize(fyne.NewSize(400, createBt.MinSize().Height))

	SetWidgetHCenter(importBt,viewSize)

	SetWidgetHCenter(createBt,viewSize)

	SetWidgetY(tip1,150)
	SetWidgetY(tip2,350)
	SetWidgetHCenter(tip1,viewSize)

	SetWidgetHCenter(tip2,viewSize)

	SetWidgetY(importBt,200)
	SetWidgetY(createBt,400)

	lay := fyne.NewContainerWithLayout(&AbLayout{Width, Height}, importBt, createBt,tip1,tip2)

	return lay

}
func importBtClick() {
	myapp.WindowInstall.SetContent(GetImportLayout())

}

func createBtClick() {
	myapp.WindowInstall.SetContent(GetCreateLayout())

}
