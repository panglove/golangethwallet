package main

import (
	"YourMoney/config"
	"YourMoney/myapp"
	"YourMoney/ui"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
)

func main() {

	OnStart()
}
func OnStart() {


	appInstall := app.NewWithID(config.AppID)
	appInstall.Settings().SetTheme(theme.LightTheme())
	winInstall := appInstall.NewWindow(config.AppName)
	myapp.AppInstall = appInstall
	myapp.WindowInstall = winInstall

	LoadIndex()

}
func LoadIndex() {
	myapp.ReadSetting()

	if len(myapp.AppSetting.RpcList) ==0 {
		myapp.AppSetting.RpcList =[]string{
			config.RpcUrl,
		}
		myapp.AppSetting.RpcUrl = config.RpcUrl
		myapp.WriteSetting()
	}

	myapp.WindowInstall.SetContent(ui.GetIndexLayout())

	myapp.WindowInstall.SetMaster()
	myapp.WindowInstall.SetFixedSize(true)
	myapp.WindowInstall.ShowAndRun()



}
