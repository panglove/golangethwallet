package config

const (
	AppName = "Your Money"

	AppID = "com.dusea.yourmoney"

	AppVersion = "1.0.0"

	AppAuthor = "Sea"


	AppSaveFileName = "Sea"


	RpcUrl = "https://etx.8kpay.com"
)

type AppSetting struct {
	RpcUrl string
	PassWord   string
	WalletList []Wallet
	RpcList []string
}
type Wallet struct {
	Name       string
	PrivateKey string
	Address    string
}
