package main

import (
	"EthSea/config"
	"EthSea/myapp"
	"EthSea/ui"
	"fyne.io/fyne/app"
)

func main() {

	OnStart()
}
func OnStart() {
	appInstall := app.NewWithID(config.AppID)
	winInstall := appInstall.NewWindow(config.AppName)
	myapp.AppInstall = appInstall
	myapp.WindowInstall = winInstall

	LoadIndex()

}
func LoadIndex() {

	myapp.WindowInstall.SetContent(ui.GetIndexLayout())

	myapp.WindowInstall.SetMaster()
	myapp.WindowInstall.SetFixedSize(true)

	myapp.WindowInstall.ShowAndRun()

}
