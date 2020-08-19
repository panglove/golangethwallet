package ui

import (
	"YourMoney/myapp"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"strings"
)
var rpcEdit *widget.Entry
func GetAddRpcDialogLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(DialogWidth, DialogHeight)

	inputTip := widget.NewLabel("Input Rpc Url")

	rpcEdit = widget.NewEntry()
	rpcEdit.SetPlaceHolder("Rpc Url")
	rpcEdit.Resize(rpcEdit.MinSize().Add(fyne.NewSize(230, 0)))
	SetWidgetHCenter(rpcEdit, viewSize)

	SetWidgetY(rpcEdit, 30)

	addRpcBt := widget.NewButton("Add", confirmAddRpcBtClick)

	addRpcBt.Resize(addRpcBt.MinSize().Add(fyne.NewSize(260, 0)))

	inputTip.Resize(inputTip.MinSize())

	SetWidgetHCenter(addRpcBt, viewSize)

	SetWidgetHCenter(inputTip, viewSize)

	SetWidgetY(inputTip, 0)

	SetWidgetY(addRpcBt, 100)

	lay := fyne.NewContainerWithLayout(&AbLayout{DialogWidth, DialogHeight}, inputTip, rpcEdit, addRpcBt)

	return lay

}
func confirmAddRpcBtClick() {

	url := rpcEdit.Text

	if len(url) < 1 || strings.Index(url,"http")!=0 {

		Alert("Rpc url is not incorrect")
		return
	}
	isok, err := myapp.AddRpc(url)

	if !isok {
		Alert(err)
		return
	}
	finishAddRpc()

}
