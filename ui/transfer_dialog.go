package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

const (
	DialogWidth  = 300
	DialogHeight = 300
)

func GetTransferDialogLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(DialogWidth, DialogHeight)

	inputTip := widget.NewLabel("Input Address And Transfer Count")

	addressEdit := widget.NewEntry()
	addressEdit.SetPlaceHolder("Receipt address")
	addressEdit.Resize(addressEdit.MinSize().Add(fyne.NewSize(200, 0)))
	SetWidgetHCenter(addressEdit, viewSize)

	SetWidgetY(addressEdit, 50)
	countEdit := widget.NewEntry()
	countEdit.SetPlaceHolder("Transfer Count")
	countEdit.Resize(countEdit.MinSize().Add(fyne.NewSize(207, 0)))
	SetWidgetHCenter(countEdit, viewSize)

	SetWidgetY(countEdit, 100)


	transferBt := widget.NewButton("Transfer", confirmTransferBtClick)

	transferBt.Resize(transferBt.MinSize().Add(fyne.NewSize(260, 0)))

	inputTip.Resize(inputTip.MinSize())

	SetWidgetHCenter(transferBt, viewSize)

	SetWidgetHCenter(inputTip, viewSize)


	SetWidgetY(inputTip, 10)

	SetWidgetY(transferBt, 180)

	lay := fyne.NewContainerWithLayout(&AbLayout{DialogWidth, DialogHeight}, inputTip, addressEdit,countEdit, transferBt)

	return lay

}
func confirmTransferBtClick() {

}
