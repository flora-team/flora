version := $(shell /bin/date "+%Y-%m-%d %H:%M")
BUILD_DIR := dist
build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" flora.go
	command -v upx &> /dev/null && upx flora
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ${BUILD_DIR}/flora-darwin flora.go
	command -v upx &> /dev/null && upx ${BUILD_DIR}/flora-darwin
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ${BUILD_DIR}/flora.exe flora.go
	command -v upx &> /dev/null && upx ${BUILD_DIR}/flora.exe
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ${BUILD_DIR}/flora-linux flora.go
	command -v upx &> /dev/null && upx ${BUILD_DIR}/flora-linux
