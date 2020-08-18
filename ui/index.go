package ui

import (
	"EthSea/config"
	"EthSea/file"
	"EthSea/myapp"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
    "fyne.io/fyne/dialog"
	"path/filepath"
)

const (
	Width = 480
	Height = 800
)
var fileSaveEdit *widget.Entry

func GetIndexLayout() fyne.CanvasObject{
	viewSize := fyne.NewSize(Width,Height)

	welcomeLabel := widget.NewLabel("Welcome to Eth Sea Wallet!")

	versionLabel := widget.NewLabel("Make by "+config.AppAuthor+" V "+config.AppVersion)

	fileSaveEdit =widget.NewEntry()

	fileSaveEdit.SetText("/Users/pza/Documents/ethsea")

	fileSaveEdit.SetPlaceHolder("please input the path of wallet's data")

	fileSaveEdit.Resize(fileSaveEdit.MinSize().Add(fyne.NewSize(100,0)))

	startBt := widget.NewButton("Get Start",startBtClick)

	startBt.Resize(startBt.MinSize().Add(fyne.NewSize(200,0)))
	welcomeLabel.Resize(welcomeLabel.MinSize())
	versionLabel.Resize(versionLabel.MinSize())

	SetWidgetHCenter(startBt,viewSize)

	SetWidgetHCenter(welcomeLabel,viewSize)

	SetWidgetHCenter(versionLabel,viewSize)

	SetWidgetHCenter(fileSaveEdit,viewSize)


	SetWidgetY(fileSaveEdit,400)

	SetWidgetY(startBt,500)

	SetWidgetY(welcomeLabel,260)

	SetWidgetY(versionLabel,600)

	lay := fyne.NewContainerWithLayout (&AbLayout{Width,Height}, welcomeLabel,startBt,fileSaveEdit,versionLabel)


	return lay


}
func startBtClick(){

	savePath := fileSaveEdit.Text

	fmt.Println(savePath)

	if len(savePath) <= 0 {

		dialog.ShowInformation("Tips","please input the folder path of wallet's data",myapp.WindowInstall)

		return

	}

	if file.PathExists(savePath) {

		myapp.AppSavePath = savePath

	} else {
		dialog.ShowInformation("Tips","The folder path does not exist",myapp.WindowInstall)

		return
	}

	settingPath := filepath.Join(savePath,config.AppSaveFileName)

	fmt.Println(settingPath)

	if !file.PathExists(settingPath) {
		isCre := file.CreateFile(settingPath,true)

		if !isCre {
			dialog.ShowInformation("Tips","System error",myapp.WindowInstall)

			return
		}
	}




	myapp.WindowInstall.SetContent(GetChooseLayout())


}