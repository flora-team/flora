version := $(shell /bin/date "+%Y-%m-%d %H:%M")

build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" flora.go
	command -v upx &> /dev/null && upx flora
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o flora-darwin flora.go
	command -v upx &> /dev/null && upx flora-darwin
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o flora.exe flora.go
	command -v upx &> /dev/null && upx flora.exe
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o flora-linux flora.go
	command -v upx &> /dev/null && upx flora-linux