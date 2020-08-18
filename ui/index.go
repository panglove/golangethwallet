package ui

import (
	"EthSea/config"
	"EthSea/myapp"
	"EthSea/util/storage"
	"encoding/json"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

const (
	Width  = 480
	Height = 800
)

var passEdit *widget.Entry

func GetIndexLayout() fyne.CanvasObject {
	viewSize := fyne.NewSize(Width, Height)

	welcomeLabel := widget.NewLabel("Welcome to Eth Sea Wallet!")

	versionLabel := widget.NewLabel("Make by " + config.AppAuthor + " V " + config.AppVersion)

	passEdit = widget.NewEntry()

	passEdit.SetPlaceHolder("please input your password")

	passEdit.Resize(passEdit.MinSize().Add(fyne.NewSize(100, 0)))

	startBt := widget.NewButton("Get Start", startBtClick)

	startBt.Resize(startBt.MinSize().Add(fyne.NewSize(200, 0)))
	welcomeLabel.Resize(welcomeLabel.MinSize())
	versionLabel.Resize(versionLabel.MinSize())

	SetWidgetHCenter(startBt, viewSize)

	SetWidgetHCenter(welcomeLabel, viewSize)

	SetWidgetHCenter(versionLabel, viewSize)

	SetWidgetHCenter(passEdit, viewSize)

	SetWidgetY(passEdit, 400)

	SetWidgetY(startBt, 500)

	SetWidgetY(welcomeLabel, 260)

	SetWidgetY(versionLabel, 600)

	lay := fyne.NewContainerWithLayout(&AbLayout{Width, Height}, welcomeLabel, startBt, passEdit, versionLabel)

	return lay

}
func startBtClick() {

	password := passEdit.Text

	if len(password) < 6 {

		dialog.ShowInformation("Tips", "Password length is at least six digits", myapp.WindowInstall)

		return

	}

	settingStr := storage.GetItem(config.AppSaveFileName)

	var appSetting *config.AppSetting

	if len(settingStr) <= 0 {
		appSetting = new(config.AppSetting)

		appSetting.PassWord = "" + password

		setByte, _ := json.Marshal(appSetting)

		isOk := storage.SetItem(config.AppSaveFileName, string(setByte))

		if !isOk {
			dialog.ShowInformation("Tips", "System error", myapp.WindowInstall)
			return
		}

	} else {

		appSetting = new(config.AppSetting)
		err := json.Unmarshal([]byte(settingStr), appSetting)

		if err != nil {
			dialog.ShowInformation("Tips", "System error", myapp.WindowInstall)
			return
		}
		if appSetting.PassWord != password {
			dialog.ShowInformation("Tips", "The password is incorrect", myapp.WindowInstall)
			return
		}

	}

	myapp.AppSetting = appSetting

	myapp.WindowInstall.SetContent(GetChooseLayout())

}
