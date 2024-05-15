build:
install:
	go install github.com/swaggo/swag/cmd/swag@latest
	sudo apt install upx
doc:
	go fmt
	swag fmt
	swag init