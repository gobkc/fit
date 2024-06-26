build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
	.bin/upx fit
install:
	go install github.com/swaggo/swag/cmd/swag@latest
	sudo wget https://github.com/upx/upx/releases/download/v4.2.4/upx-4.2.4-amd64_linux.tar.xz
	mkdir -p .bin
	tar -xf upx-4.2.4-amd64_linux.tar.xz -C .bin --strip-components=1
	rm -rf upx-4.2.4-amd64_linux.tar.xz
doc:
	go fmt
	swag fmt
	swag init
deb:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
	.bin/upx fit
	mv fit debscript/usr/bin/fit
	cp public/assets/favicon.svg debscript/usr/share/icons/fit.svg
	dpkg-deb --build debscript fit-1.0.0-amd64_linux.deb
