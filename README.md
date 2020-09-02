# 跨平台以太坊钱包
一个支持全平台的开源以太坊协议钱包，支持windows，mac，linux。后期支持手机端。


所需要的go版本 >= go1.14.4 


**1.调试运行** 

git clone https://github.com/panglove/golangethwallet.git

cd golangethwallet


go run main.go

**2.跨平台编译**


首先得安装docker容器

go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross


windows:

64位编译:

fyne-cross --targets=windows/amd64 yourmoney

32位编译:

fyne-cross --targets=windows/386 yourmoney


linux:
fyne-cross --targets=linux/amd64 yourmoney


macos:

fyne-cross --targets=darwin/amd64 yourmoney


**3.打包**

go get fyne.io/fyne/cmd/fyne



fyne package -os darwin -icon logo.png

fyne package -os linux -icon logo.png
fyne package -os windows -icon logo.png




**4.预览**


![image](https://gitee.com/seelove792/GoEthWallet/raw/master/image/5.png)
![image](https://gitee.com/seelove792/GoEthWallet/raw/master/image/6.png)
