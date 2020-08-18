package main

import (
	"EthSea/config"
	"EthSea/file"
	"EthSea/myapp"
	"EthSea/ui"
	"EthSea/util/wallet"
	"fmt"
	"fyne.io/fyne/app"
)

func main(){

	//fmt.Println(mathutil.FloatToWei(0.000001))
	//TestReadFile()
	OnStart()
	wallet.Init("")
}
func OnStart(){
	appInstall := app.NewWithID(config.AppID)
	winInstall := appInstall.NewWindow(config.AppName)
	myapp.AppInstall = appInstall
	myapp.WindowInstall = winInstall
	winInstall.SetContent(ui.GetIndexLayout())

	winInstall.SetMaster()
	winInstall.SetFixedSize(true)

	winInstall.ShowAndRun()

}


func TestReadFile(){


	str,err := file.ReadFileString("./go.mod")

	if err!=nil{

	}
	fmt.Println(str)


}


