package ui

import (
	"EthSea/myapp"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)


func GetCreateLayout() fyne.CanvasObject{
	viewSize := fyne.NewSize(Width,Height)

	inputTip := widget.NewLabel("Input your password!")

	passEdit := widget.NewEntry()
	passEdit.SetPlaceHolder("your your password")


	createBt := widget.NewButton("Create",comfirCreateBtClick)

	createBt.Resize(createBt.MinSize().Add(fyne.NewSize(200,0)))

	inputTip.Resize(inputTip.MinSize())

	passEdit.Resize(passEdit.MinSize().Add(fyne.NewSize(200,0)))

	SetWidgetHCenter(createBt,viewSize)

	SetWidgetHCenter(inputTip,viewSize)
	SetWidgetHCenter(passEdit,viewSize)

	SetWidgetY(passEdit,400)

	SetWidgetY(inputTip,300)

	SetWidgetY(createBt,600)




	lay := fyne.NewContainerWithLayout (&AbLayout{Width,Height}, inputTip,passEdit,createBt)

	return lay


}
func comfirCreateBtClick(){
	myapp.WindowInstall.SetContent(GetChooseLayout())


}