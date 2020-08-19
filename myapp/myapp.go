package myapp

import (
	"YourMoney/config"
	"YourMoney/util/storage"
	"encoding/json"
	"fyne.io/fyne"
)

var (
	AppInstall    fyne.App
	WindowInstall fyne.Window
	AppSavePath   string
	AppSetting    *config.AppSetting
)
func WriteSetting() bool {

	setByte, _ := json.Marshal(AppSetting)

	return storage.SetItem(config.AppSaveFileName, string(setByte))
}

func ReadSetting() bool {
	settingStr := storage.GetItem(config.AppSaveFileName)

	AppSetting = new(config.AppSetting)

	err := json.Unmarshal([]byte(settingStr), AppSetting)

	return  err==nil

}