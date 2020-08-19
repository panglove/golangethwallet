# golangethwallet
a open source 's eth gui wallet with golang,support windows ,macos ,linux ,android and so on!


go version >= go1.14.4 


**1.clone and run** 

git clone https://github.com/panglove/golangethwallet.git

cd golangethwallet


go run main.go

**2.cross build**


you must install docker

next 

go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross


windows:

64 bits:

fyne-cross --targets=windows/amd64 yourmoney

32 bits:

fyne-cross --targets=windows/386 yourmoney


linux:
fyne-cross --targets=linux/amd64 yourmoney


macos:

fyne-cross --targets=darwin/amd64 yourmoney


**3.package**

go get fyne.io/fyne/cmd/fyne



fyne package -os darwin -icon logo.png

fyne package -os linux -icon logo.png
fyne package -os windows -icon logo.png
