package myapp

func AddRpc(url string) (bool, string) {

	AppSetting.RpcList = append(AppSetting.RpcList, url)

	isok := WriteSetting()

	if !isok {
		return false, "System error"
	}
	return true, ""
}
